package mongodb

import (
	"bytes"
	"context"
	"time"

	"github.com/quintans/eventstore"
	"github.com/quintans/eventstore/common"
	"github.com/quintans/eventstore/sink"
	"github.com/quintans/eventstore/store"
	"github.com/quintans/faults"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Feed struct {
	connString       string
	dbName           string
	eventsCollection string
	partitions       uint32
	partitionsLow    uint32
	partitionsHi     uint32
}

type FeedOption func(*Feed)

func WithPartitions(partitions, partitionsLow, partitionsHi uint32) FeedOption {
	return func(p *Feed) {
		if partitions <= 1 {
			return
		}
		p.partitions = partitions
		p.partitionsLow = partitionsLow
		p.partitionsHi = partitionsHi
	}
}

func WithFeedEventsCollection(eventsCollection string) FeedOption {
	return func(p *Feed) {
		p.eventsCollection = eventsCollection
	}
}

func NewFeed(connString, database string, opts ...FeedOption) (Feed, error) {
	m := Feed{
		dbName:           database,
		connString:       connString,
		eventsCollection: "events",
	}

	for _, o := range opts {
		o(&m)
	}
	return m, nil
}

type ChangeEvent struct {
	FullDocument Event `bson:"fullDocument,omitempty"`
}

func (m Feed) Feed(ctx context.Context, sinker sink.Sinker) error {
	var lastResumeToken []byte
	err := store.LastEventIDInSink(ctx, sinker, m.partitionsLow, m.partitionsHi, func(resumeToken []byte) error {
		if bytes.Compare(resumeToken, lastResumeToken) > 0 {
			lastResumeToken = resumeToken
		}
		return nil
	})
	if err != nil {
		return err
	}

	ctx2, cancel := context.WithTimeout(ctx, 10*time.Second)
	client, err := mongo.Connect(ctx2, options.Client().ApplyURI(m.connString))
	cancel()
	if err != nil {
		return faults.Errorf("Unable to connect to '%s': %w", m.connString, err)
	}
	defer func() {
		client.Disconnect(context.Background())
	}()

	match := bson.D{
		{"operationType", "insert"},
	}
	if m.partitions > 1 {
		match = append(match, partitionFilter("fullDocument.aggregate_id_hash", m.partitions, m.partitionsLow, m.partitionsHi))
	}

	matchPipeline := bson.D{{Key: "$match", Value: match}}
	pipeline := mongo.Pipeline{matchPipeline}

	eventsCollection := client.Database(m.dbName).Collection(m.eventsCollection)
	var eventsStream *mongo.ChangeStream
	if len(lastResumeToken) != 0 {
		log.Infof("Starting feeding (partitions: [%d-%d]) from '%X'", m.partitionsLow, m.partitionsHi, lastResumeToken)
		eventsStream, err = eventsCollection.Watch(ctx, pipeline, options.ChangeStream().SetResumeAfter(bson.Raw(lastResumeToken)))
		if err != nil {
			return faults.Wrap(err)
		}
	} else {
		log.Infof("Starting feeding (partitions: [%d-%d]) from the beginning", m.partitionsLow, m.partitionsHi)
		eventsStream, err = eventsCollection.Watch(ctx, pipeline, options.ChangeStream().SetStartAtOperationTime(&primitive.Timestamp{}))
		if err != nil {
			return faults.Wrap(err)
		}
	}
	defer eventsStream.Close(ctx)

	for eventsStream.Next(ctx) {
		var data ChangeEvent
		if err := eventsStream.Decode(&data); err != nil {
			return faults.Wrap(err)
		}
		eventDoc := data.FullDocument

		for k, d := range eventDoc.Details {
			if k == len(eventDoc.Details)-1 {
				// we update the resume token on the last event of the transaction
				lastResumeToken = []byte(eventsStream.ResumeToken())
			}
			event := eventstore.Event{
				ID: common.NewMessageID(eventDoc.ID, uint8(k)),
				// the resume token should be from the last fully completed sinked doc, because it may fail midway.
				// We should use the last eventID to filter out the ones that were successfully sent.
				ResumeToken:      lastResumeToken,
				AggregateID:      eventDoc.AggregateID,
				AggregateIDHash:  eventDoc.AggregateIDHash,
				AggregateVersion: eventDoc.AggregateVersion,
				AggregateType:    eventDoc.AggregateType,
				Kind:             d.Kind,
				Body:             d.Body,
				IdempotencyKey:   eventDoc.IdempotencyKey,
				Labels:           eventDoc.Labels,
				CreatedAt:        eventDoc.CreatedAt,
			}
			err = sinker.Sink(ctx, event)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
