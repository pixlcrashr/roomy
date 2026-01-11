package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

type AreaQuery interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(ctx context.Context, id uuid.UUID) (*model.Area, error)

	// SELECT * FROM @@table
	// {{where}}
	//   {{if buildingID != nil}} building_id = @buildingID {{end}}
	//   {{if search != nil}} AND (name ILIKE @search OR description ILIKE @search) {{end}}
	// {{end}}
	// ORDER BY created_at DESC
	// LIMIT @limit OFFSET @offset
	List(ctx context.Context, limit int, offset int, buildingID *uuid.UUID, search *string) ([]*model.Area, error)

	// SELECT COUNT(*) FROM @@table
	// {{where}}
	//   {{if buildingID != nil}} building_id = @buildingID {{end}}
	//   {{if search != nil}} AND (name ILIKE @search OR description ILIKE @search) {{end}}
	// {{end}}
	CountAll(ctx context.Context, buildingID *uuid.UUID, search *string) (int64, error)

	// SELECT * FROM @@table WHERE building_id = @buildingID ORDER BY created_at DESC
	ListByBuilding(ctx context.Context, buildingID uuid.UUID) ([]*model.Area, error)

	// INSERT INTO @@table (id, building_id, name, description, room_plan_image, created_at, updated_at)
	// VALUES (@id, @buildingID, @name, @description, @roomPlanImage, NOW(), NOW())
	Insert(ctx context.Context, id uuid.UUID, buildingID uuid.UUID, name string, description *string, roomPlanImage *string) error

	// UPDATE @@table
	// {{set}}
	//   {{if name != nil}} name = @name, {{end}}
	//   {{if description != nil}} description = @description, {{end}}
	//   {{if roomPlanImage != nil}} room_plan_image = @roomPlanImage, {{end}}
	//   updated_at = NOW()
	// {{end}}
	// WHERE id = @id
	Save(ctx context.Context, id uuid.UUID, name *string, description *string, roomPlanImage *string) error

	// UPDATE @@table SET room_plan_image = NULL, updated_at = NOW() WHERE id = @id
	ClearRoomPlan(ctx context.Context, id uuid.UUID) error

	// DELETE FROM @@table WHERE id = @id
	Remove(ctx context.Context, id uuid.UUID) error
}
