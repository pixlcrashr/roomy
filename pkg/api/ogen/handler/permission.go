package handler

import (
	"context"

	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"gorm.io/gorm"
)

// PermissionHandler handles permission-related operations.
type PermissionHandler struct {
	db *gorm.DB
}

// NewPermissionHandler creates a new PermissionHandler.
func NewPermissionHandler(db *gorm.DB) *PermissionHandler {
	return &PermissionHandler{db: db}
}

// ListPermissions lists all available permissions.
// GET /permissions
func (h *PermissionHandler) ListPermissions(ctx context.Context) (gen.ListPermissionsRes, error) {
	// TODO: Implement list permissions
	result := gen.ListPermissionsOKApplicationJSON([]gen.Permission{})
	return &result, nil
}
