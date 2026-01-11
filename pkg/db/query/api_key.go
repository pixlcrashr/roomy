package query

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

type APIKeyQuery interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(ctx context.Context, id uuid.UUID) (*model.APIKey, error)

	// SELECT * FROM @@table WHERE value = @value
	GetByValue(ctx context.Context, value string) (*model.APIKey, error)

	// SELECT * FROM @@table WHERE user_id = @userID ORDER BY created_at DESC
	ListByUser(ctx context.Context, userID uuid.UUID) ([]*model.APIKey, error)

	// INSERT INTO @@table (id, user_id, name, value, expires_at, created_at)
	// VALUES (@id, @userID, @name, @value, @expiresAt, NOW())
	Insert(ctx context.Context, id uuid.UUID, userID uuid.UUID, name string, value string, expiresAt *time.Time) error

	// UPDATE @@table SET last_used_at = NOW() WHERE id = @id
	UpdateLastUsed(ctx context.Context, id uuid.UUID) error

	// DELETE FROM @@table WHERE id = @id AND user_id = @userID
	Remove(ctx context.Context, id uuid.UUID, userID uuid.UUID) error
}
