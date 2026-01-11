package converter

import (
	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

func BuildingToAPI(m *model.Building) *gen.Building {
	if m == nil {
		return nil
	}
	b := &gen.Building{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
	if m.Description != nil {
		b.Description.SetTo(*m.Description)
	}
	if m.Location != nil {
		b.Location.SetTo(*m.Location)
	}
	return b
}

func BuildingsToAPI(models []model.Building) []gen.Building {
	result := make([]gen.Building, len(models))
	for i, m := range models {
		result[i] = *BuildingToAPI(&m)
	}
	return result
}

func BuildingPointersToAPI(models []*model.Building) []gen.Building {
	result := make([]gen.Building, len(models))
	for i, m := range models {
		result[i] = *BuildingToAPI(m)
	}
	return result
}

func CreateBuildingRequestToModel(req *gen.CreateBuildingRequest) *model.Building {
	if req == nil {
		return nil
	}
	m := &model.Building{
		Name: req.Name,
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

func UpdateBuildingRequestToModel(req *gen.UpdateBuildingRequest, existing *model.Building) {
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
