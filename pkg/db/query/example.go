package query

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

type ExampleQuery interface {
	// SELECT ei.* FROM @@table ei
	// LEFT JOIN event_schemas es ON es.id = ei.event_schema_id
	// LEFT JOIN events e ON e.id = es.event_id
	// {{where}}
	//   {{if eventSchemaID != nil}} ei.event_schema_id = @eventSchemaID {{end}}
	//   {{if processedAt != nil}} AND ei.processed_at = @processedAt {{end}}
	//   {{if eventID != nil}} AND es.event_id = @eventID {{end}}
	//   {{if eventCustomID != nil}} AND e.custom_id = @eventCustomID {{end}}
	//   {{if startDate != nil}} AND ei.created_at >= @startDate {{end}}
	//   {{if endDate != nil}} AND ei.created_at <= @endDate {{end}}
	// {{end}}
	// ORDER BY ei.created_at DESC
	// LIMIT @limit OFFSET @offset
	List(
		ctx context.Context,
		limit int,
		offset int,
		eventSchemaID *uuid.UUID,
		processedAt *time.Time,
		eventID *uuid.UUID,
		eventCustomID *string,
		startDate *time.Time,
		endDate *time.Time,
	) ([]*model.Example, error)
}
