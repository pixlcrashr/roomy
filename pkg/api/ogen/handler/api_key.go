package handler

import (
	"context"

	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"gorm.io/gorm"
)

// APIKeyHandler handles API key operations.
type APIKeyHandler struct {
	db *gorm.DB
}

// NewAPIKeyHandler creates a new APIKeyHandler.
func NewAPIKeyHandler(db *gorm.DB) *APIKeyHandler {
	return &APIKeyHandler{db: db}
}

// CreateApiKey creates a new API key for the current user.
// POST /api-keys
func (h *APIKeyHandler) CreateApiKey(ctx context.Context, req *gen.CreateAPIKeyRequest) (gen.CreateApiKeyRes, error) {
	// TODO: Implement create API key
	return &gen.APIKeyWithSecret{}, nil
}

// ListApiKeys lists all API keys for the current user.
// GET /api-keys
func (h *APIKeyHandler) ListApiKeys(ctx context.Context) (gen.ListApiKeysRes, error) {
	// TODO: Implement list API keys
	result := gen.ListApiKeysOKApplicationJSON([]gen.APIKey{})
	return &result, nil
}

// RevokeApiKey revokes an API key.
// DELETE /api-keys/{keyId}
func (h *APIKeyHandler) RevokeApiKey(ctx context.Context, params gen.RevokeApiKeyParams) (gen.RevokeApiKeyRes, error) {
	// TODO: Implement revoke API key
	return &gen.RevokeApiKeyNoContent{}, nil
}
