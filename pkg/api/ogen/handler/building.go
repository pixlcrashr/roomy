package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"github.com/pixlcrashr/roomy/pkg/api/ogen/handler/converter"
	dbgen "github.com/pixlcrashr/roomy/pkg/db/gen"
	"github.com/pixlcrashr/roomy/pkg/db/model"
	"gorm.io/gorm"
)

// BuildingHandler handles building-related operations.
type BuildingHandler struct {
	db *gorm.DB
}

// NewBuildingHandler creates a new BuildingHandler.
func NewBuildingHandler(db *gorm.DB) *BuildingHandler {
	return &BuildingHandler{db: db}
}

// CreateBuilding creates a new building.
// POST /buildings
func (h *BuildingHandler) CreateBuilding(ctx context.Context, req *gen.CreateBuildingRequest) (gen.CreateBuildingRes, error) {
	id := uuid.New()
	var description, location *string
	if req.Description.IsSet() {
		desc := req.Description.Value
		description = &desc
	}
	if req.Location.IsSet() {
		loc := req.Location.Value
		location = &loc
	}

	if err := dbgen.BuildingQuery[model.Building](h.db).Insert(ctx, id, req.Name, description, location, nil); err != nil {
		return nil, err
	}

	building, err := dbgen.BuildingQuery[model.Building](h.db).GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return converter.BuildingToAPI(building), nil
}

// GetBuilding gets building details.
// GET /buildings/{buildingId}
func (h *BuildingHandler) GetBuilding(ctx context.Context, params gen.GetBuildingParams) (gen.GetBuildingRes, error) {
	building, err := dbgen.BuildingQuery[model.Building](h.db).GetByID(ctx, params.BuildingId)
	if err != nil {
		return nil, err
	}
	if building == nil {
		return &gen.GetBuildingNotFound{}, nil
	}
	return converter.BuildingToAPI(building), nil
}

// UpdateBuilding updates a building.
// PUT /buildings/{buildingId}
func (h *BuildingHandler) UpdateBuilding(ctx context.Context, req *gen.UpdateBuildingRequest, params gen.UpdateBuildingParams) (gen.UpdateBuildingRes, error) {
	existing, err := dbgen.BuildingQuery[model.Building](h.db).GetByID(ctx, params.BuildingId)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return &gen.UpdateBuildingNotFound{}, nil
	}

	var name, description, location *string
	if req.Name.IsSet() {
		n := req.Name.Value
		name = &n
	}
	if req.Description.IsSet() {
		desc := req.Description.Value
		description = &desc
	}
	if req.Location.IsSet() {
		loc := req.Location.Value
		location = &loc
	}

	if err := dbgen.BuildingQuery[model.Building](h.db).Save(ctx, params.BuildingId, name, description, location, nil); err != nil {
		return nil, err
	}

	updated, err := dbgen.BuildingQuery[model.Building](h.db).GetByID(ctx, params.BuildingId)
	if err != nil {
		return nil, err
	}

	return converter.BuildingToAPI(updated), nil
}

// DeleteBuilding deletes a building.
// DELETE /buildings/{buildingId}
func (h *BuildingHandler) DeleteBuilding(ctx context.Context, params gen.DeleteBuildingParams) (gen.DeleteBuildingRes, error) {
	existing, err := dbgen.BuildingQuery[model.Building](h.db).GetByID(ctx, params.BuildingId)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return &gen.DeleteBuildingNotFound{}, nil
	}

	if err := dbgen.BuildingQuery[model.Building](h.db).Remove(ctx, params.BuildingId); err != nil {
		return nil, err
	}
	return &gen.DeleteBuildingNoContent{}, nil
}

// ListBuildings lists all buildings.
// GET /buildings
func (h *BuildingHandler) ListBuildings(ctx context.Context, params gen.ListBuildingsParams) (*gen.PaginatedBuildingList, error) {
	limit := 20
	offset := 0
	if params.Limit.IsSet() {
		limit = params.Limit.Value
	}
	if params.Offset.IsSet() {
		offset = params.Offset.Value
	}

	var search *string
	if params.Search.IsSet() {
		s := "%" + params.Search.Value + "%"
		search = &s
	}

	buildings, err := dbgen.BuildingQuery[model.Building](h.db).List(ctx, limit, offset, search)
	if err != nil {
		return nil, err
	}

	total, err := dbgen.BuildingQuery[model.Building](h.db).CountAll(ctx, search)
	if err != nil {
		return nil, err
	}

	return &gen.PaginatedBuildingList{
		Items: converter.BuildingPointersToAPI(buildings),
		Total: int(total),
	}, nil
}

// GetBuildingAvailability returns time slots that are NOT blocked for the building.
// GET /buildings/{buildingId}/availability
func (h *BuildingHandler) GetBuildingAvailability(ctx context.Context, params gen.GetBuildingAvailabilityParams) (gen.GetBuildingAvailabilityRes, error) {
	building, err := dbgen.BuildingQuery[model.Building](h.db).GetByID(ctx, params.BuildingId)
	if err != nil {
		return nil, err
	}
	if building == nil {
		return &gen.GetBuildingAvailabilityNotFound{}, nil
	}

	// TODO: Calculate availability based on blocking periods
	return &gen.AvailabilityResponse{
		TimeSlots: []gen.TimeSlot{},
	}, nil
}

// GetBuildingBlocking returns all blocking periods for this building.
// GET /buildings/{buildingId}/blocking
func (h *BuildingHandler) GetBuildingBlocking(ctx context.Context, params gen.GetBuildingBlockingParams) (gen.GetBuildingBlockingRes, error) {
	building, err := dbgen.BuildingQuery[model.Building](h.db).GetByID(ctx, params.BuildingId)
	if err != nil {
		return nil, err
	}
	if building == nil {
		return &gen.GetBuildingBlockingNotFound{}, nil
	}

	blockings, err := dbgen.BlockingQuery[model.Blocking](h.db).ListByEntity(ctx, "building", params.BuildingId)
	if err != nil {
		return nil, err
	}

	return &gen.BlockingResponse{
		Entries: converter.BlockingsToAPI(blockings),
	}, nil
}

// GetBuildingCalendar returns .ics file with blocking periods and reservations.
// GET /buildings/{buildingId}/calendar.ics
func (h *BuildingHandler) GetBuildingCalendar(ctx context.Context, params gen.GetBuildingCalendarParams) (gen.GetBuildingCalendarRes, error) {
	// TODO: Generate ICS calendar
	return &gen.GetBuildingCalendarOKHeaders{}, nil
}

// AddBuildingBlockingEntries adds blocking periods to building.
// POST /buildings/{buildingId}/blocking/entries
func (h *BuildingHandler) AddBuildingBlockingEntries(ctx context.Context, req *gen.AddBuildingBlockingEntriesReq, params gen.AddBuildingBlockingEntriesParams) (gen.AddBuildingBlockingEntriesRes, error) {
	building, err := dbgen.BuildingQuery[model.Building](h.db).GetByID(ctx, params.BuildingId)
	if err != nil {
		return nil, err
	}
	if building == nil {
		return &gen.AddBuildingBlockingEntriesNotFound{}, nil
	}

	for _, entry := range req.Entries {
		id := uuid.New()
		var reason, recurrencePattern *string
		if entry.Reason.IsSet() {
			r := entry.Reason.Value
			reason = &r
		}
		if entry.RecurrenceRule.IsSet() {
			rp := entry.RecurrenceRule.Value
			recurrencePattern = &rp
		}

		if err := dbgen.BlockingQuery[model.Blocking](h.db).Insert(
			ctx, id, "building", params.BuildingId, string(entry.BlockingType),
			entry.StartTime, entry.EndTime, entry.IsRecurring.Or(false),
			recurrencePattern, nil, reason, nil,
		); err != nil {
			return nil, err
		}
	}

	blockings, err := dbgen.BlockingQuery[model.Blocking](h.db).ListByEntity(ctx, "building", params.BuildingId)
	if err != nil {
		return nil, err
	}

	return &gen.BlockingResponse{
		Entries: converter.BlockingsToAPI(blockings),
	}, nil
}

// RemoveBuildingBlockingEntries removes blocking periods by IDs.
// DELETE /buildings/{buildingId}/blocking/entries
func (h *BuildingHandler) RemoveBuildingBlockingEntries(ctx context.Context, req *gen.RemoveBuildingBlockingEntriesReq, params gen.RemoveBuildingBlockingEntriesParams) (gen.RemoveBuildingBlockingEntriesRes, error) {
	if err := dbgen.BlockingQuery[model.Blocking](h.db).DeleteByIDs(ctx, "building", params.BuildingId, req.Ids); err != nil {
		return nil, err
	}
	return &gen.RemoveBuildingBlockingEntriesNoContent{}, nil
}

// ReplaceBuildingBlocking replaces all building blocking periods.
// PUT /buildings/{buildingId}/blocking
func (h *BuildingHandler) ReplaceBuildingBlocking(ctx context.Context, req *gen.ReplaceBuildingBlockingReq, params gen.ReplaceBuildingBlockingParams) (gen.ReplaceBuildingBlockingRes, error) {
	building, err := dbgen.BuildingQuery[model.Building](h.db).GetByID(ctx, params.BuildingId)
	if err != nil {
		return nil, err
	}
	if building == nil {
		return &gen.ReplaceBuildingBlockingNotFound{}, nil
	}

	if err := dbgen.BlockingQuery[model.Blocking](h.db).DeleteByEntity(ctx, "building", params.BuildingId); err != nil {
		return nil, err
	}

	for _, entry := range req.Entries {
		id := uuid.New()
		var reason, recurrencePattern *string
		if entry.Reason.IsSet() {
			r := entry.Reason.Value
			reason = &r
		}
		if entry.RecurrenceRule.IsSet() {
			rp := entry.RecurrenceRule.Value
			recurrencePattern = &rp
		}

		if err := dbgen.BlockingQuery[model.Blocking](h.db).Insert(
			ctx, id, "building", params.BuildingId, string(entry.BlockingType),
			entry.StartTime, entry.EndTime, entry.IsRecurring.Or(false),
			recurrencePattern, nil, reason, nil,
		); err != nil {
			return nil, err
		}
	}

	blockings, err := dbgen.BlockingQuery[model.Blocking](h.db).ListByEntity(ctx, "building", params.BuildingId)
	if err != nil {
		return nil, err
	}

	return &gen.BlockingResponse{
		Entries: converter.BlockingsToAPI(blockings),
	}, nil
}

// ListBuildingAreas lists areas in building.
// GET /buildings/{buildingId}/areas
func (h *BuildingHandler) ListBuildingAreas(ctx context.Context, params gen.ListBuildingAreasParams) (gen.ListBuildingAreasRes, error) {
	building, err := dbgen.BuildingQuery[model.Building](h.db).GetByID(ctx, params.BuildingId)
	if err != nil {
		return nil, err
	}
	if building == nil {
		return &gen.ListBuildingAreasNotFound{}, nil
	}

	areas, err := dbgen.AreaQuery[model.Area](h.db).ListByBuilding(ctx, params.BuildingId)
	if err != nil {
		return nil, err
	}

	return &gen.ListBuildingAreasOKApplicationJSON(converter.AreaPointersToAPI(areas)), nil
}
