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

// AreaHandler handles area-related operations.
type AreaHandler struct {
	db *gorm.DB
}

// NewAreaHandler creates a new AreaHandler.
func NewAreaHandler(db *gorm.DB) *AreaHandler {
	return &AreaHandler{db: db}
}

// CreateArea creates a new area.
// POST /areas
func (h *AreaHandler) CreateArea(ctx context.Context, req *gen.CreateAreaRequest) (gen.CreateAreaRes, error) {
	id := uuid.New()
	var description *string
	if req.Description.IsSet() {
		desc := req.Description.Value
		description = &desc
	}

	if err := dbgen.AreaQuery[model.Area](h.db).Insert(ctx, id, req.BuildingId, req.Name, description, nil); err != nil {
		return nil, err
	}

	area, err := dbgen.AreaQuery[model.Area](h.db).GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return converter.AreaToAPI(area), nil
}

// GetArea gets area details.
// GET /areas/{areaId}
func (h *AreaHandler) GetArea(ctx context.Context, params gen.GetAreaParams) (gen.GetAreaRes, error) {
	area, err := dbgen.AreaQuery[model.Area](h.db).GetByID(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}
	if area == nil {
		return &gen.GetAreaNotFound{}, nil
	}
	return converter.AreaToAPI(area), nil
}

// UpdateArea updates an area.
// PUT /areas/{areaId}
func (h *AreaHandler) UpdateArea(ctx context.Context, req *gen.UpdateAreaRequest, params gen.UpdateAreaParams) (gen.UpdateAreaRes, error) {
	existing, err := dbgen.AreaQuery[model.Area](h.db).GetByID(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return &gen.UpdateAreaNotFound{}, nil
	}

	var name, description *string
	if req.Name.IsSet() {
		n := req.Name.Value
		name = &n
	}
	if req.Description.IsSet() {
		desc := req.Description.Value
		description = &desc
	}

	if err := dbgen.AreaQuery[model.Area](h.db).Save(ctx, params.AreaId, name, description, nil); err != nil {
		return nil, err
	}

	updated, err := dbgen.AreaQuery[model.Area](h.db).GetByID(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}

	return converter.AreaToAPI(updated), nil
}

// DeleteArea deletes an area.
// DELETE /areas/{areaId}
func (h *AreaHandler) DeleteArea(ctx context.Context, params gen.DeleteAreaParams) (gen.DeleteAreaRes, error) {
	existing, err := dbgen.AreaQuery[model.Area](h.db).GetByID(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return &gen.DeleteAreaNotFound{}, nil
	}

	if err := dbgen.AreaQuery[model.Area](h.db).Remove(ctx, params.AreaId); err != nil {
		return nil, err
	}
	return &gen.DeleteAreaNoContent{}, nil
}

// ListAreas lists all areas.
// GET /areas
func (h *AreaHandler) ListAreas(ctx context.Context, params gen.ListAreasParams) (*gen.PaginatedAreaList, error) {
	limit := 20
	offset := 0
	if params.Limit.IsSet() {
		limit = params.Limit.Value
	}
	if params.Offset.IsSet() {
		offset = params.Offset.Value
	}

	var buildingID *uuid.UUID
	if params.BuildingId.IsSet() {
		bid := params.BuildingId.Value
		buildingID = &bid
	}

	var search *string
	if params.Search.IsSet() {
		s := "%" + params.Search.Value + "%"
		search = &s
	}

	areas, err := dbgen.AreaQuery[model.Area](h.db).List(ctx, limit, offset, buildingID, search)
	if err != nil {
		return nil, err
	}

	total, err := dbgen.AreaQuery[model.Area](h.db).CountAll(ctx, buildingID, search)
	if err != nil {
		return nil, err
	}

	return &gen.PaginatedAreaList{
		Items: converter.AreaPointersToAPI(areas),
		Total: int(total),
	}, nil
}

// GetAreaAvailability returns availability status for each place in the area.
// GET /areas/{areaId}/availability
func (h *AreaHandler) GetAreaAvailability(ctx context.Context, params gen.GetAreaAvailabilityParams) (gen.GetAreaAvailabilityRes, error) {
	area, err := dbgen.AreaQuery[model.Area](h.db).GetByID(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}
	if area == nil {
		return &gen.GetAreaAvailabilityNotFound{}, nil
	}

	// TODO: Calculate availability per place
	return &gen.AreaAvailabilityResponse{
		Places: []gen.PlaceAvailability{},
	}, nil
}

// GetAreaBlocking returns blocking periods for the area.
// GET /areas/{areaId}/blocking
func (h *AreaHandler) GetAreaBlocking(ctx context.Context, params gen.GetAreaBlockingParams) (gen.GetAreaBlockingRes, error) {
	area, err := dbgen.AreaQuery[model.Area](h.db).GetByID(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}
	if area == nil {
		return &gen.GetAreaBlockingNotFound{}, nil
	}

	blockings, err := dbgen.BlockingQuery[model.Blocking](h.db).ListByEntity(ctx, "area", params.AreaId)
	if err != nil {
		return nil, err
	}

	// TODO: Include inherited blocking from building if requested
	return &gen.BlockingResponse{
		Entries: converter.BlockingsToAPI(blockings),
	}, nil
}

// GetAreaCalendar returns .ics file with blocking periods and reservations.
// GET /areas/{areaId}/calendar.ics
func (h *AreaHandler) GetAreaCalendar(ctx context.Context, params gen.GetAreaCalendarParams) (gen.GetAreaCalendarRes, error) {
	// TODO: Generate ICS calendar
	return &gen.GetAreaCalendarOKHeaders{}, nil
}

// AddAreaBlockingEntries adds blocking periods to area.
// POST /areas/{areaId}/blocking/entries
func (h *AreaHandler) AddAreaBlockingEntries(ctx context.Context, req *gen.AddAreaBlockingEntriesReq, params gen.AddAreaBlockingEntriesParams) (gen.AddAreaBlockingEntriesRes, error) {
	area, err := dbgen.AreaQuery[model.Area](h.db).GetByID(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}
	if area == nil {
		return &gen.AddAreaBlockingEntriesNotFound{}, nil
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
			ctx, id, "area", params.AreaId, string(entry.BlockingType),
			entry.StartTime, entry.EndTime, entry.IsRecurring.Or(false),
			recurrencePattern, nil, reason, nil,
		); err != nil {
			return nil, err
		}
	}

	blockings, err := dbgen.BlockingQuery[model.Blocking](h.db).ListByEntity(ctx, "area", params.AreaId)
	if err != nil {
		return nil, err
	}

	return &gen.BlockingResponse{
		Entries: converter.BlockingsToAPI(blockings),
	}, nil
}

// RemoveAreaBlockingEntries removes blocking periods by IDs.
// DELETE /areas/{areaId}/blocking/entries
func (h *AreaHandler) RemoveAreaBlockingEntries(ctx context.Context, req *gen.RemoveAreaBlockingEntriesReq, params gen.RemoveAreaBlockingEntriesParams) (gen.RemoveAreaBlockingEntriesRes, error) {
	if err := dbgen.BlockingQuery[model.Blocking](h.db).DeleteByIDs(ctx, "area", params.AreaId, req.Ids); err != nil {
		return nil, err
	}
	return &gen.RemoveAreaBlockingEntriesNoContent{}, nil
}

// ReplaceAreaBlocking replaces all area blocking periods.
// PUT /areas/{areaId}/blocking
func (h *AreaHandler) ReplaceAreaBlocking(ctx context.Context, req *gen.ReplaceAreaBlockingReq, params gen.ReplaceAreaBlockingParams) (gen.ReplaceAreaBlockingRes, error) {
	area, err := dbgen.AreaQuery[model.Area](h.db).GetByID(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}
	if area == nil {
		return &gen.ReplaceAreaBlockingNotFound{}, nil
	}

	if err := dbgen.BlockingQuery[model.Blocking](h.db).DeleteByEntity(ctx, "area", params.AreaId); err != nil {
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
			ctx, id, "area", params.AreaId, string(entry.BlockingType),
			entry.StartTime, entry.EndTime, entry.IsRecurring.Or(false),
			recurrencePattern, nil, reason, nil,
		); err != nil {
			return nil, err
		}
	}

	blockings, err := dbgen.BlockingQuery[model.Blocking](h.db).ListByEntity(ctx, "area", params.AreaId)
	if err != nil {
		return nil, err
	}

	return &gen.BlockingResponse{
		Entries: converter.BlockingsToAPI(blockings),
	}, nil
}

// ListAreaPlaces lists places in area.
// GET /areas/{areaId}/places
func (h *AreaHandler) ListAreaPlaces(ctx context.Context, params gen.ListAreaPlacesParams) (gen.ListAreaPlacesRes, error) {
	area, err := dbgen.AreaQuery[model.Area](h.db).GetByID(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}
	if area == nil {
		return &gen.ListAreaPlacesNotFound{}, nil
	}

	places, err := dbgen.PlaceQuery[model.Place](h.db).ListByArea(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}

	return &gen.ListAreaPlacesOKApplicationJSON(converter.PlacePointersToAPI(places)), nil
}

// GetAreaRoomPlan gets room plan data.
// GET /areas/{areaId}/roomPlan
func (h *AreaHandler) GetAreaRoomPlan(ctx context.Context, params gen.GetAreaRoomPlanParams) (gen.GetAreaRoomPlanRes, error) {
	area, err := dbgen.AreaQuery[model.Area](h.db).GetByID(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}
	if area == nil {
		return &gen.GetAreaRoomPlanNotFound{}, nil
	}

	markers, err := dbgen.PlaceMarkerQuery[model.PlaceMarker](h.db).ListByArea(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}
	area.PlaceMarkers = markers

	return converter.RoomPlanToAPI(area), nil
}

// UpdateAreaRoomPlan updates room plan.
// PUT /areas/{areaId}/roomPlan
func (h *AreaHandler) UpdateAreaRoomPlan(ctx context.Context, req *gen.UpdateAreaRoomPlanReq, params gen.UpdateAreaRoomPlanParams) (gen.UpdateAreaRoomPlanRes, error) {
	area, err := dbgen.AreaQuery[model.Area](h.db).GetByID(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}
	if area == nil {
		return &gen.UpdateAreaRoomPlanNotFound{}, nil
	}

	var roomPlanImage *string
	if req.ImageUrl.IsSet() {
		imageURL := req.ImageUrl.Value.String()
		roomPlanImage = &imageURL
	}

	if err := dbgen.AreaQuery[model.Area](h.db).Save(ctx, params.AreaId, nil, nil, roomPlanImage); err != nil {
		return nil, err
	}

	updated, err := dbgen.AreaQuery[model.Area](h.db).GetByID(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}

	markers, err := dbgen.PlaceMarkerQuery[model.PlaceMarker](h.db).ListByArea(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}
	updated.PlaceMarkers = markers

	return converter.RoomPlanToAPI(updated), nil
}

// DeleteAreaRoomPlan deletes room plan.
// DELETE /areas/{areaId}/roomPlan
func (h *AreaHandler) DeleteAreaRoomPlan(ctx context.Context, params gen.DeleteAreaRoomPlanParams) (gen.DeleteAreaRoomPlanRes, error) {
	area, err := dbgen.AreaQuery[model.Area](h.db).GetByID(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}
	if area == nil {
		return &gen.DeleteAreaRoomPlanNotFound{}, nil
	}

	if err := dbgen.PlaceMarkerQuery[model.PlaceMarker](h.db).DeleteByArea(ctx, params.AreaId); err != nil {
		return nil, err
	}

	if err := dbgen.AreaQuery[model.Area](h.db).ClearRoomPlan(ctx, params.AreaId); err != nil {
		return nil, err
	}

	return &gen.DeleteAreaRoomPlanNoContent{}, nil
}

// AddAreaRoomPlanMarkers adds place markers to room plan.
// POST /areas/{areaId}/roomPlan/markers
func (h *AreaHandler) AddAreaRoomPlanMarkers(ctx context.Context, req *gen.AddAreaRoomPlanMarkersReq, params gen.AddAreaRoomPlanMarkersParams) (gen.AddAreaRoomPlanMarkersRes, error) {
	area, err := dbgen.AreaQuery[model.Area](h.db).GetByID(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}
	if area == nil {
		return &gen.AddAreaRoomPlanMarkersNotFound{}, nil
	}

	for _, m := range req.Markers {
		id := uuid.New()
		if err := dbgen.PlaceMarkerQuery[model.PlaceMarker](h.db).Insert(
			ctx, id, params.AreaId, m.PlaceId,
			float64(m.X), float64(m.Y), float64(m.Width), float64(m.Height),
		); err != nil {
			return nil, err
		}
	}

	markers, err := dbgen.PlaceMarkerQuery[model.PlaceMarker](h.db).ListByArea(ctx, params.AreaId)
	if err != nil {
		return nil, err
	}
	area.PlaceMarkers = markers

	return converter.RoomPlanToAPI(area), nil
}

// RemoveAreaRoomPlanMarkers removes place markers by IDs.
// DELETE /areas/{areaId}/roomPlan/markers
func (h *AreaHandler) RemoveAreaRoomPlanMarkers(ctx context.Context, req *gen.RemoveAreaRoomPlanMarkersReq, params gen.RemoveAreaRoomPlanMarkersParams) (gen.RemoveAreaRoomPlanMarkersRes, error) {
	if err := dbgen.PlaceMarkerQuery[model.PlaceMarker](h.db).DeleteByIDs(ctx, params.AreaId, req.Ids); err != nil {
		return nil, err
	}
	return &gen.RemoveAreaRoomPlanMarkersNoContent{}, nil
}
