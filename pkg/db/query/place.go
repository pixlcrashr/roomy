package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

type PlaceQuery interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(ctx context.Context, id uuid.UUID) (*model.Place, error)

	// SELECT p.* FROM @@table p
	// LEFT JOIN areas a ON a.id = p.area_id
	// {{where}}
	//   {{if areaID != nil}} p.area_id = @areaID {{end}}
	//   {{if buildingID != nil}} AND a.building_id = @buildingID {{end}}
	//   {{if search != nil}} AND (p.name ILIKE @search OR p.description ILIKE @search) {{end}}
	//   {{if minCapacity != nil}} AND p.seat_capacity >= @minCapacity {{end}}
	//   {{if isBookable != nil}} AND p.is_bookable = @isBookable {{end}}
	// {{end}}
	// ORDER BY p.created_at DESC
	// LIMIT @limit OFFSET @offset
	List(ctx context.Context, limit int, offset int, areaID *uuid.UUID, buildingID *uuid.UUID, search *string, minCapacity *int, isBookable *bool) ([]*model.Place, error)

	// SELECT COUNT(p.id) FROM @@table p
	// LEFT JOIN areas a ON a.id = p.area_id
	// {{where}}
	//   {{if areaID != nil}} p.area_id = @areaID {{end}}
	//   {{if buildingID != nil}} AND a.building_id = @buildingID {{end}}
	//   {{if search != nil}} AND (p.name ILIKE @search OR p.description ILIKE @search) {{end}}
	//   {{if minCapacity != nil}} AND p.seat_capacity >= @minCapacity {{end}}
	//   {{if isBookable != nil}} AND p.is_bookable = @isBookable {{end}}
	// {{end}}
	CountAll(ctx context.Context, areaID *uuid.UUID, buildingID *uuid.UUID, search *string, minCapacity *int, isBookable *bool) (int64, error)

	// SELECT * FROM @@table WHERE area_id = @areaID ORDER BY created_at DESC
	ListByArea(ctx context.Context, areaID uuid.UUID) ([]*model.Place, error)

	// INSERT INTO @@table (
	//   id, area_id, name, description, seat_capacity, is_bookable, is_disabled,
	//   max_reservation_duration, max_reservations_per_day, max_reservations_per_week,
	//   max_reservations_per_month, max_reservations_per_year, requires_check_in,
	//   check_in_timeout_minutes, has_whitelist, booking_window_days, min_booking_duration,
	//   time_slot_interval_minutes, day_start_hour, day_end_hour, created_at, updated_at
	// ) VALUES (
	//   @id, @areaID, @name, @description, @seatCapacity, @isBookable, @isDisabled,
	//   @maxReservationDuration, @maxReservationsPerDay, @maxReservationsPerWeek,
	//   @maxReservationsPerMonth, @maxReservationsPerYear, @requiresCheckIn,
	//   @checkInTimeoutMinutes, @hasWhitelist, @bookingWindowDays, @minBookingDuration,
	//   @timeSlotIntervalMinutes, @dayStartHour, @dayEndHour, NOW(), NOW()
	// )
	Insert(
		ctx context.Context,
		id uuid.UUID,
		areaID uuid.UUID,
		name string,
		description *string,
		seatCapacity int,
		isBookable bool,
		isDisabled bool,
		maxReservationDuration *int,
		maxReservationsPerDay *int,
		maxReservationsPerWeek *int,
		maxReservationsPerMonth *int,
		maxReservationsPerYear *int,
		requiresCheckIn bool,
		checkInTimeoutMinutes *int,
		hasWhitelist bool,
		bookingWindowDays *int,
		minBookingDuration *int,
		timeSlotIntervalMinutes int,
		dayStartHour int,
		dayEndHour int,
	) error

	// UPDATE @@table
	// {{set}}
	//   {{if name != nil}} name = @name, {{end}}
	//   {{if description != nil}} description = @description, {{end}}
	//   {{if seatCapacity != nil}} seat_capacity = @seatCapacity, {{end}}
	//   {{if isBookable != nil}} is_bookable = @isBookable, {{end}}
	//   {{if isDisabled != nil}} is_disabled = @isDisabled, {{end}}
	//   updated_at = NOW()
	// {{end}}
	// WHERE id = @id
	Save(ctx context.Context, id uuid.UUID, name *string, description *string, seatCapacity *int, isBookable *bool, isDisabled *bool) error

	// UPDATE @@table
	// {{set}}
	//   {{if maxReservationDuration != nil}} max_reservation_duration = @maxReservationDuration, {{end}}
	//   {{if maxReservationsPerDay != nil}} max_reservations_per_day = @maxReservationsPerDay, {{end}}
	//   {{if maxReservationsPerWeek != nil}} max_reservations_per_week = @maxReservationsPerWeek, {{end}}
	//   {{if maxReservationsPerMonth != nil}} max_reservations_per_month = @maxReservationsPerMonth, {{end}}
	//   {{if maxReservationsPerYear != nil}} max_reservations_per_year = @maxReservationsPerYear, {{end}}
	//   {{if requiresCheckIn != nil}} requires_check_in = @requiresCheckIn, {{end}}
	//   {{if checkInTimeoutMinutes != nil}} check_in_timeout_minutes = @checkInTimeoutMinutes, {{end}}
	//   {{if hasWhitelist != nil}} has_whitelist = @hasWhitelist, {{end}}
	//   {{if bookingWindowDays != nil}} booking_window_days = @bookingWindowDays, {{end}}
	//   {{if minBookingDuration != nil}} min_booking_duration = @minBookingDuration, {{end}}
	//   updated_at = NOW()
	// {{end}}
	// WHERE id = @id
	UpdateConstraints(
		ctx context.Context,
		id uuid.UUID,
		maxReservationDuration *int,
		maxReservationsPerDay *int,
		maxReservationsPerWeek *int,
		maxReservationsPerMonth *int,
		maxReservationsPerYear *int,
		requiresCheckIn *bool,
		checkInTimeoutMinutes *int,
		hasWhitelist *bool,
		bookingWindowDays *int,
		minBookingDuration *int,
	) error

	// UPDATE @@table
	// {{set}}
	//   {{if timeSlotIntervalMinutes != nil}} time_slot_interval_minutes = @timeSlotIntervalMinutes, {{end}}
	//   {{if dayStartHour != nil}} day_start_hour = @dayStartHour, {{end}}
	//   {{if dayEndHour != nil}} day_end_hour = @dayEndHour, {{end}}
	//   updated_at = NOW()
	// {{end}}
	// WHERE id = @id
	UpdateTimeSlots(ctx context.Context, id uuid.UUID, timeSlotIntervalMinutes *int, dayStartHour *int, dayEndHour *int) error

	// DELETE FROM @@table WHERE id = @id
	Remove(ctx context.Context, id uuid.UUID) error
}
