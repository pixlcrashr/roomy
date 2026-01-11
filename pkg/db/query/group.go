package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

type GroupQuery interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(ctx context.Context, id uuid.UUID) (*model.Group, error)

	// SELECT * FROM @@table ORDER BY created_at DESC
	List(ctx context.Context) ([]*model.Group, error)

	// SELECT * FROM @@table WHERE is_default = true
	ListDefaultGroups(ctx context.Context) ([]*model.Group, error)

	// SELECT * FROM @@table WHERE is_system = true
	GetSystemGroup(ctx context.Context) (*model.Group, error)

	// INSERT INTO @@table (id, name, description, is_system, is_default, created_at, updated_at)
	// VALUES (@id, @name, @description, false, @isDefault, NOW(), NOW())
	Insert(ctx context.Context, id uuid.UUID, name string, description *string, isDefault bool) error

	// UPDATE @@table
	// {{set}}
	//   {{if name != nil}} name = @name, {{end}}
	//   {{if description != nil}} description = @description, {{end}}
	//   updated_at = NOW()
	// {{end}}
	// WHERE id = @id AND is_system = false
	Save(ctx context.Context, id uuid.UUID, name *string, description *string) error

	// UPDATE @@table SET is_default = @isDefault, updated_at = NOW() WHERE id = @id
	SetDefault(ctx context.Context, id uuid.UUID, isDefault bool) error

	// DELETE FROM @@table WHERE id = @id AND is_system = false
	Remove(ctx context.Context, id uuid.UUID) error
}
