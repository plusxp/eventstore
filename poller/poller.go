package poller

import (
	"context"
	"time"

	"github.com/quintans/eventstore"
	log "github.com/sirupsen/logrus"
)

const (
	maxWait = time.Minute
)

type Repository interface {
	GetLastEventID(ctx context.Context, filter eventstore.Filter) (string, error)
	GetEvents(ctx context.Context, afterEventID string, limit int, filter eventstore.Filter) ([]eventstore.Event, error)
}

type Start int

const (
	END Start = iota
	BEGINNING
	SEQUENCE
)

type EventHandler func(ctx context.Context, e eventstore.Event) error

type Cancel func()

type Option func(*Poller)

func WithPollInterval(pi time.Duration) Option {
	return func(p *Poller) {
		p.pollInterval = pi
	}
}

func WithFilter(filter eventstore.Filter) Option {
	return func(p *Poller) {
		p.filter = filter
	}
}

func WithAggregateTypes(at ...string) Option {
	return func(p *Poller) {
		p.filter.AggregateTypes = at
	}
}

func WithLabels(labels ...eventstore.Label) Option {
	return func(p *Poller) {
		p.filter.Labels = labels
	}
}

func WithLimit(limit int) Option {
	return func(p *Poller) {
		if limit > 0 {
			p.limit = limit
		}
	}
}

func New(repo Repository, options ...Option) *Poller {
	p := &Poller{
		repo:         repo,
		pollInterval: 500 * time.Millisecond,
		limit:        20,
		filter:       eventstore.Filter{},
	}
	for _, o := range options {
		o(p)
	}
	return p
}

type Poller struct {
	repo         Repository
	pollInterval time.Duration
	filter       eventstore.Filter
	limit        int
}

type StartOption struct {
	startFrom    Start
	afterEventID string
}

func StartEnd() StartOption {
	return StartOption{
		startFrom: END,
	}
}

func StartBeginning() StartOption {
	return StartOption{
		startFrom: BEGINNING,
	}
}

func StartAt(afterEventID string) StartOption {
	return StartOption{
		startFrom:    SEQUENCE,
		afterEventID: afterEventID,
	}
}

func (p *Poller) Handle(ctx context.Context, startOption StartOption, handler EventHandler) error {
	var afterEventID string
	var err error
	switch startOption.startFrom {
	case END:
		afterEventID, err = p.repo.GetLastEventID(ctx, eventstore.Filter{})
		if err != nil {
			return err
		}
	case BEGINNING:
	case SEQUENCE:
		afterEventID = startOption.afterEventID
	}
	return p.handle(ctx, afterEventID, handler)
}

func (p *Poller) handle(ctx context.Context, afterEventID string, handler EventHandler) error {
	wait := p.pollInterval
	for {
		eid, err := p.retrieve(ctx, handler, afterEventID, "")
		if err != nil {
			wait += 2 * wait
			if wait > maxWait {
				wait = maxWait
			}
			log.WithField("backoff", wait).
				WithError(err).
				Error("Failure retrieving events. Backing off.")
		} else {
			afterEventID = eid
			wait = p.pollInterval
		}

		select {
		case <-ctx.Done():
			return nil
		case _ = <-time.After(p.pollInterval):
		}
	}
}

type Sink interface {
	LastEventID(ctx context.Context) (string, error)
	Send(ctx context.Context, e eventstore.Event) error
}

// Forward forwars the handling to a sink.
// eg: a message queue
func (p *Poller) Forward(ctx context.Context, sink Sink) error {
	id, err := sink.LastEventID(ctx)
	if err != nil {
		return err
	}
	return p.handle(ctx, id, sink.Send)
}

func (p *Poller) ReplayUntil(ctx context.Context, handler EventHandler, untilEventID string) (string, error) {
	return p.retrieve(ctx, handler, "", untilEventID)
}

func (p *Poller) ReplayFromUntil(ctx context.Context, handler EventHandler, afterEventID, untilEventID string) (string, error) {
	return p.retrieve(ctx, handler, afterEventID, untilEventID)
}

func (p *Poller) retrieve(ctx context.Context, handler EventHandler, afterEventID, untilEventID string) (string, error) {
	loop := true
	for loop {
		events, err := p.repo.GetEvents(ctx, afterEventID, p.limit, p.filter)
		if err != nil {
			return "", err
		}
		for _, evt := range events {
			err := handler(ctx, evt)
			if err != nil {
				return "", err
			}
			afterEventID = evt.ID
			if evt.ID == untilEventID {
				return untilEventID, nil
			}
		}
		loop = len(events) == p.limit
	}
	return afterEventID, nil
}
