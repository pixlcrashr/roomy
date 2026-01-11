package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

type BuildingQuery interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(ctx context.Context, id uuid.UUID) (*model.Building, error)

	// SELECT * FROM @@table
	// {{where}}
	//   {{if search != nil}} name ILIKE @search OR description ILIKE @search {{end}}
	// {{end}}
	// ORDER BY created_at DESC
	// LIMIT @limit OFFSET @offset
	List(ctx context.Context, limit int, offset int, search *string) ([]*model.Building, error)

	// SELECT COUNT(*) FROM @@table
	// {{where}}
	//   {{if search != nil}} name ILIKE @search OR description ILIKE @search {{end}}
	// {{end}}
	CountAll(ctx context.Context, search *string) (int64, error)

	// INSERT INTO @@table (id, name, description, location, opening_hours, created_at, updated_at)
	// VALUES (@id, @name, @description, @location, @openingHours, NOW(), NOW())
	Insert(ctx context.Context, id uuid.UUID, name string, description *string, location *string, openingHours *string) error

	// UPDATE @@table
	// {{set}}
	//   {{if name != nil}} name = @name, {{end}}
	//   {{if description != nil}} description = @description, {{end}}
	//   {{if location != nil}} location = @location, {{end}}
	//   {{if openingHours != nil}} opening_hours = @openingHours, {{end}}
	//   updated_at = NOW()
	// {{end}}
	// WHERE id = @id
	Save(ctx context.Context, id uuid.UUID, name *string, description *string, location *string, openingHours *string) error

	// DELETE FROM @@table WHERE id = @id
	Remove(ctx context.Context, id uuid.UUID) error
}
