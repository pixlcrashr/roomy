package handler

import (
	"context"

	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"gorm.io/gorm"
)

// ReservationHandler handles reservation-related operations.
type ReservationHandler struct {
	db *gorm.DB
}

// NewReservationHandler creates a new ReservationHandler.
func NewReservationHandler(db *gorm.DB) *ReservationHandler {
	return &ReservationHandler{db: db}
}

// CreateReservation creates a single reservation or multiple recurring reservations.
// POST /reservations
func (h *ReservationHandler) CreateReservation(ctx context.Context, req *gen.CreateReservationRequest) (gen.CreateReservationRes, error) {
	// TODO: Implement reservation creation
	return &gen.CreateReservationCreated{
		Type:        gen.ReservationCreateReservationCreated,
		Reservation: gen.Reservation{},
	}, nil
}

// GetReservation gets reservation details.
// GET /reservations/{reservationId}
func (h *ReservationHandler) GetReservation(ctx context.Context, params gen.GetReservationParams) (gen.GetReservationRes, error) {
	// TODO: Implement get reservation
	return &gen.GetReservationNotFound{}, nil
}

// UpdateReservation modifies the reservation's start and/or end time.
// PUT /reservations/{reservationId}
func (h *ReservationHandler) UpdateReservation(ctx context.Context, req *gen.UpdateReservationRequest, params gen.UpdateReservationParams) (gen.UpdateReservationRes, error) {
	// TODO: Implement update reservation
	return &gen.UpdateReservationNotFound{}, nil
}

// CancelReservation cancels a reservation.
// DELETE /reservations/{reservationId}
func (h *ReservationHandler) CancelReservation(ctx context.Context, params gen.CancelReservationParams) (gen.CancelReservationRes, error) {
	// TODO: Implement cancel reservation
	return &gen.CancelReservationNoContent{}, nil
}

// ListReservations lists reservations.
// GET /reservations
func (h *ReservationHandler) ListReservations(ctx context.Context, params gen.ListReservationsParams) (gen.ListReservationsRes, error) {
	// TODO: Implement list reservations
	return &gen.PaginatedReservationList{
		Data: []gen.Reservation{},
		Meta: gen.PaginationMeta{
			Page:       1,
			Limit:      20,
			Total:      0,
			TotalPages: 0,
		},
	}, nil
}

// CheckInReservation checks in to a reservation.
// POST /reservations/{reservationId}/checkIn
func (h *ReservationHandler) CheckInReservation(ctx context.Context, req gen.OptCheckInReservationReq, params gen.CheckInReservationParams) (gen.CheckInReservationRes, error) {
	// TODO: Implement check-in
	return &gen.CheckInReservationNotFound{}, nil
}

// GetReservationShareLink returns a shareable URL and metadata.
// GET /reservations/{reservationId}/share
func (h *ReservationHandler) GetReservationShareLink(ctx context.Context, params gen.GetReservationShareLinkParams) (gen.GetReservationShareLinkRes, error) {
	// TODO: Implement share link generation
	return &gen.GetReservationShareLinkNotFound{}, nil
}

// ExportReservations exports reservations as CSV.
// GET /reservations/export
func (h *ReservationHandler) ExportReservations(ctx context.Context, params gen.ExportReservationsParams) (gen.ExportReservationsRes, error) {
	// TODO: Implement CSV export
	return &gen.ExportReservationsOK{
		Data: nil,
	}, nil
}
