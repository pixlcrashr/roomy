package handler

import (
	"context"

	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"gorm.io/gorm"
)

// AuthHandler handles authentication-related operations.
type AuthHandler struct {
	db *gorm.DB
}

// NewAuthHandler creates a new AuthHandler.
func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

// InitiateOAuthLogin redirects to GitLab authorization URL.
// GET /auth/login
func (h *AuthHandler) InitiateOAuthLogin(ctx context.Context) error {
	// TODO: Implement OAuth login redirect
	return nil
}

// HandleOAuthCallback handles the OAuth callback from GitLab.
// GET /auth/callback
func (h *AuthHandler) HandleOAuthCallback(ctx context.Context, params gen.HandleOAuthCallbackParams) (gen.HandleOAuthCallbackRes, error) {
	// TODO: Implement OAuth callback handling
	return &gen.AuthTokens{}, nil
}

// GetCurrentUser gets the current user profile.
// GET /auth/me
func (h *AuthHandler) GetCurrentUser(ctx context.Context) (gen.GetCurrentUserRes, error) {
	// TODO: Get user from context (set by auth middleware)
	return &gen.User{}, nil
}

// Logout invalidates the current session.
// POST /auth/logout
func (h *AuthHandler) Logout(ctx context.Context) error {
	// TODO: Implement logout
	return nil
}

// RefreshToken refreshes the access token.
// POST /auth/refresh
func (h *AuthHandler) RefreshToken(ctx context.Context, req *gen.RefreshTokenReq) (gen.RefreshTokenRes, error) {
	// TODO: Implement token refresh
	return &gen.AuthTokens{}, nil
}
