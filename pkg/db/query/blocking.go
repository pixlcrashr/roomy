package query

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

type BlockingQuery interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(ctx context.Context, id uuid.UUID) (*model.Blocking, error)

	// SELECT * FROM @@table WHERE entity_type = @entityType AND entity_id = @entityID
	// ORDER BY start_time
	ListByEntity(ctx context.Context, entityType string, entityID uuid.UUID) ([]*model.Blocking, error)

	// SELECT * FROM @@table
	// {{where}}
	//   entity_type = @entityType AND entity_id = @entityID
	//   {{if startAfter != nil}} AND start_time >= @startAfter {{end}}
	//   {{if endBefore != nil}} AND end_time <= @endBefore {{end}}
	// {{end}}
	// ORDER BY start_time
	ListByEntityAndTimeRange(
		ctx context.Context,
		entityType string,
		entityID uuid.UUID,
		startAfter *time.Time,
		endBefore *time.Time,
	) ([]*model.Blocking, error)

	// SELECT * FROM @@table WHERE
	// (entity_type = 'place' AND entity_id = @placeID)
	// OR (entity_type = 'area' AND entity_id = @areaID)
	// OR (entity_type = 'building' AND entity_id = @buildingID)
	// ORDER BY start_time
	ListInheritedForPlace(ctx context.Context, placeID uuid.UUID, areaID uuid.UUID, buildingID uuid.UUID) ([]*model.Blocking, error)

	// INSERT INTO @@table (
	//   id, entity_type, entity_id, blocking_type, start_time, end_time,
	//   is_recurring, recurrence_pattern, recurrence_end, reason, created_by, created_at
	// ) VALUES (
	//   @id, @entityType, @entityID, @blockingType, @startTime, @endTime,
	//   @isRecurring, @recurrencePattern, @recurrenceEnd, @reason, @createdBy, NOW()
	// )
	Insert(
		ctx context.Context,
		id uuid.UUID,
		entityType string,
		entityID uuid.UUID,
		blockingType string,
		startTime time.Time,
		endTime time.Time,
		isRecurring bool,
		recurrencePattern *string,
		recurrenceEnd *time.Time,
		reason *string,
		createdBy *uuid.UUID,
	) error

	// DELETE FROM @@table WHERE id = @id
	Remove(ctx context.Context, id uuid.UUID) error

	// DELETE FROM @@table WHERE entity_type = @entityType AND entity_id = @entityID AND id IN (@ids)
	DeleteByIDs(ctx context.Context, entityType string, entityID uuid.UUID, ids []uuid.UUID) error

	// DELETE FROM @@table WHERE entity_type = @entityType AND entity_id = @entityID
	DeleteByEntity(ctx context.Context, entityType string, entityID uuid.UUID) error
}
