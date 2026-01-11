package query

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

type ReservationQuery interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(ctx context.Context, id uuid.UUID) (*model.Reservation, error)

	// SELECT * FROM @@table
	// {{where}}
	//   {{if placeID != nil}} place_id = @placeID {{end}}
	//   {{if userID != nil}} AND user_id = @userID {{end}}
	//   {{if status != nil}} AND status = @status {{end}}
	//   {{if startAfter != nil}} AND start_time >= @startAfter {{end}}
	//   {{if endBefore != nil}} AND end_time <= @endBefore {{end}}
	// {{end}}
	// ORDER BY start_time DESC
	// LIMIT @limit OFFSET @offset
	List(
		ctx context.Context,
		limit int,
		offset int,
		placeID *uuid.UUID,
		userID *uuid.UUID,
		status *string,
		startAfter *time.Time,
		endBefore *time.Time,
	) ([]*model.Reservation, error)

	// SELECT COUNT(*) FROM @@table
	// {{where}}
	//   {{if placeID != nil}} place_id = @placeID {{end}}
	//   {{if userID != nil}} AND user_id = @userID {{end}}
	//   {{if status != nil}} AND status = @status {{end}}
	//   {{if startAfter != nil}} AND start_time >= @startAfter {{end}}
	//   {{if endBefore != nil}} AND end_time <= @endBefore {{end}}
	// {{end}}
	CountAll(
		ctx context.Context,
		placeID *uuid.UUID,
		userID *uuid.UUID,
		status *string,
		startAfter *time.Time,
		endBefore *time.Time,
	) (int64, error)

	// SELECT * FROM @@table WHERE user_id = @userID ORDER BY start_time DESC
	ListByUser(ctx context.Context, userID uuid.UUID) ([]*model.Reservation, error)

	// SELECT * FROM @@table WHERE place_id = @placeID AND start_time >= @startTime AND end_time <= @endTime ORDER BY start_time
	ListByPlaceAndTimeRange(ctx context.Context, placeID uuid.UUID, startTime time.Time, endTime time.Time) ([]*model.Reservation, error)

	// INSERT INTO @@table (
	//   id, place_id, user_id, start_time, end_time, status, is_recurring,
	//   recurring_group_id, created_at, updated_at
	// ) VALUES (
	//   @id, @placeID, @userID, @startTime, @endTime, @status, @isRecurring,
	//   @recurringGroupID, NOW(), NOW()
	// )
	Insert(
		ctx context.Context,
		id uuid.UUID,
		placeID uuid.UUID,
		userID uuid.UUID,
		startTime time.Time,
		endTime time.Time,
		status string,
		isRecurring bool,
		recurringGroupID *uuid.UUID,
	) error

	// UPDATE @@table
	// {{set}}
	//   {{if startTime != nil}} start_time = @startTime, {{end}}
	//   {{if endTime != nil}} end_time = @endTime, {{end}}
	//   updated_at = NOW()
	// {{end}}
	// WHERE id = @id
	Save(ctx context.Context, id uuid.UUID, startTime *time.Time, endTime *time.Time) error

	// UPDATE @@table SET status = @status, updated_at = NOW() WHERE id = @id
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) error

	// UPDATE @@table SET status = 'cancelled', cancel_reason = @reason, updated_at = NOW() WHERE id = @id
	Cancel(ctx context.Context, id uuid.UUID, reason *string) error

	// UPDATE @@table SET status = 'checkedIn', check_in_time = NOW(), updated_at = NOW() WHERE id = @id
	CheckIn(ctx context.Context, id uuid.UUID) error

	// DELETE FROM @@table WHERE id = @id
	Remove(ctx context.Context, id uuid.UUID) error
}
