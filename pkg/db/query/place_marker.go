package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

type PlaceMarkerQuery interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(ctx context.Context, id uuid.UUID) (*model.PlaceMarker, error)

	// SELECT * FROM @@table WHERE area_id = @areaID ORDER BY created_at
	ListByArea(ctx context.Context, areaID uuid.UUID) ([]*model.PlaceMarker, error)

	// INSERT INTO @@table (id, area_id, place_id, x, y, width, height, created_at)
	// VALUES (@id, @areaID, @placeID, @x, @y, @width, @height, NOW())
	Insert(ctx context.Context, id uuid.UUID, areaID uuid.UUID, placeID uuid.UUID, x float64, y float64, width float64, height float64) error

	// UPDATE @@table SET x = @x, y = @y, width = @width, height = @height WHERE id = @id
	Save(ctx context.Context, id uuid.UUID, x float64, y float64, width float64, height float64) error

	// DELETE FROM @@table WHERE id = @id
	Remove(ctx context.Context, id uuid.UUID) error

	// DELETE FROM @@table WHERE area_id = @areaID AND id IN (@ids)
	DeleteByIDs(ctx context.Context, areaID uuid.UUID, ids []uuid.UUID) error

	// DELETE FROM @@table WHERE area_id = @areaID
	DeleteByArea(ctx context.Context, areaID uuid.UUID) error
}
