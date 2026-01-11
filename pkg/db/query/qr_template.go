package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

type QRTemplateQuery interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(ctx context.Context, id uuid.UUID) (*model.QRTemplate, error)

	// SELECT * FROM @@table ORDER BY created_at DESC
	List(ctx context.Context) ([]*model.QRTemplate, error)

	// INSERT INTO @@table (id, name, html_template, created_at, updated_at)
	// VALUES (@id, @name, @htmlTemplate, NOW(), NOW())
	Insert(ctx context.Context, id uuid.UUID, name string, htmlTemplate string) error

	// UPDATE @@table
	// {{set}}
	//   {{if name != nil}} name = @name, {{end}}
	//   {{if htmlTemplate != nil}} html_template = @htmlTemplate, {{end}}
	//   updated_at = NOW()
	// {{end}}
	// WHERE id = @id
	Save(ctx context.Context, id uuid.UUID, name *string, htmlTemplate *string) error

	// DELETE FROM @@table WHERE id = @id
	Remove(ctx context.Context, id uuid.UUID) error
}
