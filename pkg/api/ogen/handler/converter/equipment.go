package converter

import (
	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

func EquipmentToAPI(m *model.Equipment) *gen.Equipment {
	if m == nil {
		return nil
	}
	e := &gen.Equipment{
		ID:   m.ID,
		Name: m.Name,
	}
	if m.Icon != nil {
		e.Icon.SetTo(*m.Icon)
	}
	if m.Description != nil {
		e.Description.SetTo(*m.Description)
	}
	return e
}

func EquipmentSliceToAPI(models []*model.Equipment) []gen.Equipment {
	result := make([]gen.Equipment, len(models))
	for i, m := range models {
		result[i] = *EquipmentToAPI(m)
	}
	return result
}

func CreateEquipmentRequestToModel(req *gen.CreateEquipmentRequest) *model.Equipment {
	if req == nil {
		return nil
	}
	m := &model.Equipment{
		Name: req.Name,
	}
	if req.Icon.IsSet() {
		icon := req.Icon.Value
		m.Icon = &icon
	}
	if req.Description.IsSet() {
		desc := req.Description.Value
		m.Description = &desc
	}
	return m
}

func UpdateEquipmentRequestToModel(req *gen.UpdateEquipmentRequest, existing *model.Equipment) {
	if req == nil || existing == nil {
		return
	}
	if req.Name.IsSet() {
		existing.Name = req.Name.Value
	}
	if req.Icon.IsSet() {
		icon := req.Icon.Value
		existing.Icon = &icon
	}
	if req.Description.IsSet() {
		desc := req.Description.Value
		existing.Description = &desc
	}
}
