package handler

import (
	"context"

	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"gorm.io/gorm"
)

// EquipmentHandler handles equipment-related operations.
type EquipmentHandler struct {
	db *gorm.DB
}

// NewEquipmentHandler creates a new EquipmentHandler.
func NewEquipmentHandler(db *gorm.DB) *EquipmentHandler {
	return &EquipmentHandler{db: db}
}

// CreateEquipment creates a new equipment type.
// POST /equipment
func (h *EquipmentHandler) CreateEquipment(ctx context.Context, req *gen.CreateEquipmentRequest) (gen.CreateEquipmentRes, error) {
	// TODO: Implement create equipment
	return &gen.Equipment{}, nil
}

// GetEquipment gets equipment details.
// GET /equipment/{equipmentId}
func (h *EquipmentHandler) GetEquipment(ctx context.Context, params gen.GetEquipmentParams) (gen.GetEquipmentRes, error) {
	// TODO: Implement get equipment
	return &gen.Equipment{}, nil
}

// UpdateEquipment updates an equipment type.
// PUT /equipment/{equipmentId}
func (h *EquipmentHandler) UpdateEquipment(ctx context.Context, req *gen.UpdateEquipmentRequest, params gen.UpdateEquipmentParams) (gen.UpdateEquipmentRes, error) {
	// TODO: Implement update equipment
	return &gen.UpdateEquipmentNotFound{}, nil
}

// DeleteEquipment deletes an equipment type.
// DELETE /equipment/{equipmentId}
func (h *EquipmentHandler) DeleteEquipment(ctx context.Context, params gen.DeleteEquipmentParams) (gen.DeleteEquipmentRes, error) {
	// TODO: Implement delete equipment
	return &gen.DeleteEquipmentNoContent{}, nil
}

// ListEquipment lists all equipment types.
// GET /equipment
func (h *EquipmentHandler) ListEquipment(ctx context.Context) ([]gen.Equipment, error) {
	// TODO: Implement list equipment
	return []gen.Equipment{}, nil
}
