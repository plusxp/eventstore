package eventstore

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

const (
	uniqueViolation = "23505"
	lag             = -200 * time.Millisecond
)

var _ EventStore = (*ESPostgreSQL)(nil)
var _ Tracker = (*ESPostgreSQL)(nil)

// NewESPostgreSQL creates a new instance of ESPostgreSQL
func NewESPostgreSQL(dburl string, snapshotThreshold int) (*ESPostgreSQL, error) {
	db, err := sqlx.Open("postgres", dburl)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &ESPostgreSQL{
		db:                db,
		snapshotThreshold: snapshotThreshold,
	}, nil
}

// ESPostgreSQL is the implementation of an Event Store in PostgreSQL
type ESPostgreSQL struct {
	db                *sqlx.DB
	snapshotThreshold int
}

func (es *ESPostgreSQL) GetByID(ctx context.Context, aggregateID string, aggregate Aggregater) error {
	snap, err := es.getSnapshot(ctx, aggregateID)
	if err != nil {
		return err
	}

	var events []PgEvent
	if snap != nil {
		err = json.Unmarshal(snap.Body, aggregate)
		if err != nil {
			return err
		}
		events, err = es.getEvents(ctx, aggregateID, snap.AggregateVersion)
	} else {
		events, err = es.getEvents(ctx, aggregateID, -1)
	}
	if err != nil {
		return err
	}

	for _, v := range events {
		aggregate.ApplyChangeFromHistory(Event{
			AggregateID:      v.AggregateID,
			AggregateVersion: v.AggregateVersion,
			AggregateType:    v.AggregateType,
			Kind:             v.Kind,
			Body:             v.Body,
			CreatedAt:        v.CreatedAt,
		})
	}

	return nil
}

func (es *ESPostgreSQL) withTx(ctx context.Context, fn func(context.Context, *sql.Tx) error) (err error) {
	tx, err := es.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
		if err != nil {
			tx.Rollback()
		}
	}()
	err = fn(ctx, tx)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (es *ESPostgreSQL) Save(ctx context.Context, aggregate Aggregater, options Options) (err error) {
	events := aggregate.GetEvents()
	if len(events) == 0 {
		return nil
	}

	tName := nameFor(aggregate)
	version := aggregate.GetVersion()
	oldVersion := version

	var eventID string
	var takeSnapshot bool
	defer func() {
		if err != nil {
			aggregate.SetVersion(oldVersion)
			return
		}
		if takeSnapshot {
			snap, err := buildSnapshot(aggregate, eventID)
			if err != nil {
				go es.saveSnapshot(ctx, snap)
			}
		}
	}()

	err = es.withTx(ctx, func(c context.Context, tx *sql.Tx) error {
		for _, e := range events {
			version++
			aggregate.SetVersion(version)
			body, err := json.Marshal(e)
			if err != nil {
				return err
			}
			labels, err := json.Marshal(options.Labels)
			if err != nil {
				return err
			}

			err = es.saveEvent(c, tx, aggregate, tName, nameFor(e), body, options.IdempotencyKey, labels)
			if err != nil {
				return err
			}

			if version > es.snapshotThreshold-1 &&
				version%es.snapshotThreshold == 0 {
				takeSnapshot = true
			}

		}
		aggregate.ClearEvents()
		return nil
	})

	return err
}

func nameFor(x interface{}) string {
	t := reflect.TypeOf(x)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}

func (es *ESPostgreSQL) HasIdempotencyKey(ctx context.Context, aggregateID, idempotencyKey string) (bool, error) {
	var exists int
	err := es.db.GetContext(ctx, &exists, `SELECT EXISTS(SELECT 1 FROM events WHERE idempotency_key=$1 AND aggregate_type=$2) AS "EXISTS"`, idempotencyKey, aggregateID)
	if err != nil {
		return false, fmt.Errorf("Unable to verify the existence of the idempotency key: %w", err)
	}
	return exists != 0, nil
}

func (es *ESPostgreSQL) GetLastEventID(ctx context.Context) (string, error) {
	var eventID string
	safetyMargin := time.Now().Add(lag)
	if err := es.db.GetContext(ctx, &eventID, `
	SELECT * FROM events
	WHERE created_at <= $1'
	ORDER BY id DESC LIMIT 1
	`, safetyMargin); err != nil {
		if err != sql.ErrNoRows {
			return "", fmt.Errorf("Unable to get the last event ID: %w", err)
		}
	}
	return eventID, nil
}

func (es *ESPostgreSQL) GetEventsForAggregate(ctx context.Context, afterEventID string, aggregateID string, batchSize int) ([]Event, error) {
	events := []Event{}
	safetyMargin := time.Now().Add(lag)
	if err := es.db.SelectContext(ctx, &events, `
	SELECT * FROM events
	WHERE id > $1
	AND aggregate_id = $2
	AND created_at <= $3
	ORDER BY id ASC LIMIT $4
	`, afterEventID, aggregateID, safetyMargin, batchSize); err != nil {
		if err != sql.ErrNoRows {
			return nil, fmt.Errorf("Unable to get events after '%s': %w", afterEventID, err)
		}
	}
	return events, nil

}

func (es *ESPostgreSQL) GetEvents(ctx context.Context, afterEventID string, batchSize int, filter Filter) ([]Event, error) {
	safetyMargin := time.Now().Add(lag)
	args := []interface{}{afterEventID, safetyMargin}
	var query bytes.Buffer
	query.WriteString("SELECT * FROM events WHERE id > $1 AND created_at <= $2")
	if len(filter.AggregateTypes) > 0 {
		query.WriteString(" AND aggregate_type IN ($3)")
		args = append(args, filter.AggregateTypes)
	}
	if len(filter.Labels) > 0 {
		for k, v := range filter.Labels {
			k = escape(k)

			query.WriteString(" AND (")
			first := true
			for _, x := range v {
				if !first {
					query.WriteString(" OR ")
				}
				first = false
				x = escape(x)
				query.WriteString(fmt.Sprintf(`labels  @> '{"%s": "%s"}'`, k, x))
			}
			query.WriteString(")")
		}
	}
	query.WriteString(" ORDER BY id ASC LIMIT ")
	query.WriteString(strconv.Itoa(batchSize))

	rows, err := es.queryEvents(ctx, query.String(), args)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, fmt.Errorf("Unable to get events after '%s' for filter %+v: %w", afterEventID, filter, err)
		}
	}
	return rows, nil
}

func (es *ESPostgreSQL) queryEvents(ctx context.Context, query string, args []interface{}) ([]Event, error) {
	rows, err := es.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	events := []Event{}
	for rows.Next() {
		pg := PgEvent{}
		err := rows.StructScan(&pg)
		if err != nil {
			return nil, fmt.Errorf("Unable to scan to struct: %w", err)
		}
		events = append(events, Event{
			ID:               pg.ID,
			AggregateID:      pg.AggregateID,
			AggregateVersion: pg.AggregateVersion,
			AggregateType:    pg.AggregateType,
			Kind:             pg.Kind,
			Body:             pg.Body,
			Labels:           pg.Labels,
			CreatedAt:        pg.CreatedAt,
		})
	}
	return events, nil
}

type ForgetRequest struct {
	AggregateID     string
	AggregateFields []string
	Events          []EventKind
}

type EventKind struct {
	Kind   string
	Fields []string
}

func (es *ESPostgreSQL) Forget(ctx context.Context, request ForgetRequest) error {
	for _, evt := range request.Events {
		sql := joinAndEscape(evt.Fields)
		sql = fmt.Sprintf("UPDATE events SET body =  body - '{%s}'::text[] WHERE aggregate_id = $1 AND kind = $2", sql)
		_, err := es.db.ExecContext(ctx, sql, request.AggregateID, evt.Kind)
		if err != nil {
			return fmt.Errorf("Unable to forget events: %w", err)
		}
	}

	sql := joinAndEscape(request.AggregateFields)
	sql = fmt.Sprintf("UPDATE snapshots SET body =  body - '{%s}'::text[] WHERE aggregate_id = $1", sql)
	_, err := es.db.ExecContext(ctx, sql, request.AggregateID)
	if err != nil {
		return fmt.Errorf("Unable to forget snapshots: %w", err)
	}

	return nil
}

func joinAndEscape(s []string) string {
	fields := strings.Join(s, ", ")
	return escape(fields)
}

func escape(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}

func (es *ESPostgreSQL) saveEvent(ctx context.Context, tx *sql.Tx, aggregate Aggregater, tName, eName string, body []byte, idempotencyKey string, labels []byte) error {
	eventID := xid.New().String()
	now := time.Now().UTC()

	_, err := tx.ExecContext(ctx,
		`INSERT INTO events (id, aggregate_id, aggregate_version, aggregate_type, kind, body, idempotency_key, labels, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		eventID, aggregate.GetID(), aggregate.GetVersion(), tName, eName, body, idempotencyKey, labels, now)
	if err != nil {
		if pgerr, ok := err.(*pq.Error); ok {
			if pgerr.Code == uniqueViolation {
				return ErrConcurrentModification
			}
		}
		return fmt.Errorf("Unable to insert event: %w", err)
	}
	return nil
}

func (es *ESPostgreSQL) getSnapshot(ctx context.Context, aggregateID string) (*PgSnapshot, error) {
	snap := &PgSnapshot{}
	if err := es.db.GetContext(ctx, snap, "SELECT * FROM snapshots WHERE aggregate_id = $1 ORDER BY id DESC LIMIT 1", aggregateID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Unable to get snapshot for aggregate '%s': %w", aggregateID, err)
	}
	return snap, nil
}

func (es *ESPostgreSQL) saveSnapshot(ctx context.Context, snap interface{}) {
	_, err := es.db.NamedExecContext(ctx,
		`INSERT INTO snapshots (id, aggregate_id, aggregate_version, body, created_at)
	     VALUES (:id, :aggregate_id, :aggregate_version, :body, :created_at)`, snap)

	if err != nil {
		log.WithField("snapshot", snap).
			WithError(err).
			Error("Failed to save snapshot")
	}
}

func buildSnapshot(agg Aggregater, eventID string) (*PgSnapshot, error) {
	body, err := json.Marshal(agg)
	if err != nil {
		log.WithField("aggregate", agg).
			WithError(err).
			Error("Failed to serialize aggregate")
		return nil, err
	}

	return &PgSnapshot{
		ID:               eventID,
		AggregateID:      agg.GetID(),
		AggregateVersion: agg.GetVersion(),
		Body:             body,
		CreatedAt:        time.Now().UTC(),
	}, nil
}

func (es *ESPostgreSQL) getEvents(ctx context.Context, aggregateID string, snapVersion int) ([]PgEvent, error) {
	query := "SELECT * FROM events e WHERE e.aggregate_id = $1"
	args := []interface{}{aggregateID}
	if snapVersion > -1 {
		query += " AND e.aggregate_version > $2"
		args = append(args, snapVersion)
	}
	events := []PgEvent{}
	if err := es.db.SelectContext(ctx, &events, query, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Aggregate '%s' was not found: %w", aggregateID, err)
		}
		return nil, fmt.Errorf("Unable to get events for Aggregate '%s': %w", aggregateID, err)
	}
	return events, nil
}

type PgEvent struct {
	ID               string    `db:"id"`
	AggregateID      string    `db:"aggregate_id"`
	AggregateVersion int       `db:"aggregate_version"`
	AggregateType    string    `db:"aggregate_type"`
	Kind             string    `db:"kind"`
	Body             Json      `db:"body"`
	IdempotencyKey   string    `db:"idempotency_key"`
	Labels           Json      `db:"labels"`
	CreatedAt        time.Time `db:"created_at"`
}

type PgSnapshot struct {
	ID               string    `db:"id"`
	AggregateID      string    `db:"aggregate_id"`
	AggregateVersion int       `db:"aggregate_version"`
	Body             Json      `db:"body"`
	CreatedAt        time.Time `db:"created_at"`
}
