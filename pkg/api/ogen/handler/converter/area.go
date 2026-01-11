package converter

import (
	"net/url"

	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

func AreaToAPI(m *model.Area) *gen.Area {
	if m == nil {
		return nil
	}
	a := &gen.Area{
		ID:         m.ID,
		BuildingId: m.BuildingID,
		Name:       m.Name,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
	if m.Description != nil {
		a.Description.SetTo(*m.Description)
	}
	if m.Location != nil {
		a.Location.SetTo(*m.Location)
	}
	if m.RoomPlanURL != nil {
		parsedURL, err := url.Parse(*m.RoomPlanURL)
		if err == nil {
			a.RoomPlan.SetTo(gen.RoomPlan{
				ImageUrl: *parsedURL,
				Markers:  PlaceMarkersToAPI(m.PlaceMarkers),
			})
		}
	}
	return a
}

func AreasToAPI(models []*model.Area) []gen.Area {
	result := make([]gen.Area, len(models))
	for i, m := range models {
		result[i] = *AreaToAPI(m)
	}
	return result
}

func CreateAreaRequestToModel(req *gen.CreateAreaRequest) *model.Area {
	if req == nil {
		return nil
	}
	m := &model.Area{
		BuildingID: req.BuildingId,
		Name:       req.Name,
	}
	if req.Description.IsSet() {
		desc := req.Description.Value
		m.Description = &desc
	}
	if req.Location.IsSet() {
		loc := req.Location.Value
		m.Location = &loc
	}
	return m
}

func UpdateAreaRequestToModel(req *gen.UpdateAreaRequest, existing *model.Area) {
	if req == nil || existing == nil {
		return
	}
	if req.Name.IsSet() {
		existing.Name = req.Name.Value
	}
	if req.Description.IsSet() {
		desc := req.Description.Value
		existing.Description = &desc
	}
	if req.Location.IsSet() {
		loc := req.Location.Value
		existing.Location = &loc
	}
}
