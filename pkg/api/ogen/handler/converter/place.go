package converter

import (
	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

func PlaceToAPI(m *model.Place) *gen.Place {
	if m == nil {
		return nil
	}
	p := &gen.Place{
		ID:              m.ID,
		AreaId:          m.AreaID,
		Name:            m.Name,
		Capacity:        m.Capacity,
		IsBookable:      m.IsBookable,
		BookingMethod:   gen.PlaceBookingMethod(m.BookingMethod),
		IsDisabled:      m.IsDisabled,
		RequiresCheckIn: m.RequiresCheckIn,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
	if m.Description != nil {
		p.Description.SetTo(*m.Description)
	}
	if m.Location != nil {
		p.Location.SetTo(*m.Location)
	}
	if len(m.Equipment) > 0 {
		p.Equipment = EquipmentSliceToAPI(m.Equipment)
	}
	return p
}

func PlacesToAPI(models []*model.Place) []gen.Place {
	result := make([]gen.Place, len(models))
	for i, m := range models {
		result[i] = *PlaceToAPI(m)
	}
	return result
}

func CreatePlaceRequestToModel(req *gen.CreatePlaceRequest) *model.Place {
	if req == nil {
		return nil
	}
	m := &model.Place{
		AreaID:     req.AreaId,
		Name:       req.Name,
		Capacity:   1,
		IsBookable: true,
	}
	if req.Description.IsSet() {
		desc := req.Description.Value
		m.Description = &desc
	}
	if req.Location.IsSet() {
		loc := req.Location.Value
		m.Location = &loc
	}
	if req.Capacity.IsSet() {
		m.Capacity = req.Capacity.Value
	}
	if req.IsBookable.IsSet() {
		m.IsBookable = req.IsBookable.Value
	}
	if req.BookingMethod.IsSet() {
		m.BookingMethod = model.BookingMethod(req.BookingMethod.Value)
	}
	if req.RequiresCheckIn.IsSet() {
		m.RequiresCheckIn = req.RequiresCheckIn.Value
	}
	return m
}

func UpdatePlaceRequestToModel(req *gen.UpdatePlaceRequest, existing *model.Place) {
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
	if req.Capacity.IsSet() {
		existing.Capacity = req.Capacity.Value
	}
	if req.IsBookable.IsSet() {
		existing.IsBookable = req.IsBookable.Value
	}
	if req.BookingMethod.IsSet() {
		existing.BookingMethod = model.BookingMethod(req.BookingMethod.Value)
	}
	if req.IsDisabled.IsSet() {
		existing.IsDisabled = req.IsDisabled.Value
	}
	if req.RequiresCheckIn.IsSet() {
		existing.RequiresCheckIn = req.RequiresCheckIn.Value
	}
}
