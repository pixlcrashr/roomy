package converter

import (
	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

func GroupToAPI(m *model.Group) *gen.Group {
	if m == nil {
		return nil
	}
	g := &gen.Group{
		ID:        m.ID,
		Name:      m.Name,
		IsSystem:  m.IsSystem,
		IsDefault: m.IsDefault,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
	if m.Description != nil {
		g.Description.SetTo(*m.Description)
	}
	return g
}

func GroupsToAPI(models []*model.Group) []gen.Group {
	result := make([]gen.Group, len(models))
	for i, m := range models {
		result[i] = *GroupToAPI(m)
	}
	return result
}

func CreateGroupRequestToModel(req *gen.CreateGroupRequest) *model.Group {
	if req == nil {
		return nil
	}
	m := &model.Group{
		Name: req.Name,
	}
	if req.Description.IsSet() {
		desc := req.Description.Value
		m.Description = &desc
	}
	return m
}

func UpdateGroupRequestToModel(req *gen.UpdateGroupRequest, existing *model.Group) {
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
}
