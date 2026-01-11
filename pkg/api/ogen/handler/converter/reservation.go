package converter

import (
	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

func ReservationToAPI(m *model.Reservation) *gen.Reservation {
	if m == nil {
		return nil
	}
	r := &gen.Reservation{
		ID:          m.ID,
		PlaceId:     m.PlaceID,
		UserId:      m.UserID,
		StartTime:   m.StartTime,
		EndTime:     m.EndTime,
		Status:      gen.ReservationStatus(m.Status),
		IsRecurring: m.IsRecurring,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
	if m.CheckInTime != nil {
		r.CheckInTime.SetTo(*m.CheckInTime)
	}
	if m.CancelReason != nil {
		r.CancelReason.SetTo(*m.CancelReason)
	}
	if m.RecurringGroupID != nil {
		r.RecurringGroupId.SetTo(*m.RecurringGroupID)
	}
	return r
}

func ReservationsToAPI(models []*model.Reservation) []gen.Reservation {
	result := make([]gen.Reservation, len(models))
	for i, m := range models {
		result[i] = *ReservationToAPI(m)
	}
	return result
}

func CreateReservationRequestToModel(req *gen.CreateReservationRequest, userID uuid.UUID) *model.Reservation {
	if req == nil {
		return nil
	}
	m := &model.Reservation{
		PlaceID:   req.PlaceId,
		UserID:    userID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Status:    model.ReservationStatusPending,
	}
	if req.Recurrence.IsSet() {
		m.IsRecurring = true
	}
	return m
}
