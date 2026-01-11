package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

type EquipmentQuery interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(ctx context.Context, id uuid.UUID) (*model.Equipment, error)

	// SELECT * FROM @@table ORDER BY name
	List(ctx context.Context) ([]*model.Equipment, error)

	// SELECT e.* FROM @@table e
	// INNER JOIN place_equipment pe ON pe.equipment_id = e.id
	// WHERE pe.place_id = @placeID
	// ORDER BY e.name
	ListByPlace(ctx context.Context, placeID uuid.UUID) ([]*model.Equipment, error)

	// INSERT INTO @@table (id, name, icon, created_at, updated_at)
	// VALUES (@id, @name, @icon, NOW(), NOW())
	Insert(ctx context.Context, id uuid.UUID, name string, icon *string) error

	// UPDATE @@table
	// {{set}}
	//   {{if name != nil}} name = @name, {{end}}
	//   {{if icon != nil}} icon = @icon, {{end}}
	//   updated_at = NOW()
	// {{end}}
	// WHERE id = @id
	Save(ctx context.Context, id uuid.UUID, name *string, icon *string) error

	// DELETE FROM @@table WHERE id = @id
	Remove(ctx context.Context, id uuid.UUID) error
}
