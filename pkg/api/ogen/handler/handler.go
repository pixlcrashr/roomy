package handler

import (
	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"gorm.io/gorm"
)

// Handler implements the gen.Handler interface by composing domain-specific handlers.
type Handler struct {
	db *gorm.DB

	// Embed domain handlers
	*AuthHandler
	*BuildingHandler
	*AreaHandler
	*PlaceHandler
	*ReservationHandler
	*UserHandler
	*GroupHandler
	*EquipmentHandler
	*QRTemplateHandler
	*AuditLogHandler
	*APIKeyHandler
	*StatisticsHandler
	*PermissionHandler
}

// Verify that Handler implements gen.Handler at compile time.
var _ gen.Handler = (*Handler)(nil)

// NewHandler creates a new Handler with all domain handlers initialized.
func NewHandler(db *gorm.DB) *Handler {
	h := &Handler{
		db: db,
	}

	h.AuthHandler = NewAuthHandler(db)
	h.BuildingHandler = NewBuildingHandler(db)
	h.AreaHandler = NewAreaHandler(db)
	h.PlaceHandler = NewPlaceHandler(db)
	h.ReservationHandler = NewReservationHandler(db)
	h.UserHandler = NewUserHandler(db)
	h.GroupHandler = NewGroupHandler(db)
	h.EquipmentHandler = NewEquipmentHandler(db)
	h.QRTemplateHandler = NewQRTemplateHandler(db)
	h.AuditLogHandler = NewAuditLogHandler(db)
	h.APIKeyHandler = NewAPIKeyHandler(db)
	h.StatisticsHandler = NewStatisticsHandler(db)
	h.PermissionHandler = NewPermissionHandler(db)

	return h
}
