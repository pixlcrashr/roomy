package converter

import (
	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

func PlaceMarkerToAPI(m *model.PlaceMarker) *gen.PlaceMarker {
	if m == nil {
		return nil
	}
	return &gen.PlaceMarker{
		ID:      m.ID,
		PlaceId: m.PlaceID,
		X:       float32(m.X),
		Y:       float32(m.Y),
		Width:   float32(m.Width),
		Height:  float32(m.Height),
		Shape:   gen.PlaceMarkerShape(m.Shape),
	}
}

func PlaceMarkersToAPI(models []*model.PlaceMarker) []gen.PlaceMarker {
	result := make([]gen.PlaceMarker, len(models))
	for i, m := range models {
		result[i] = *PlaceMarkerToAPI(m)
	}
	return result
}

func CreatePlaceMarkerRequestToModel(req *gen.CreatePlaceMarkerRequest, areaID string) *model.PlaceMarker {
	if req == nil {
		return nil
	}
	return &model.PlaceMarker{
		PlaceID: req.PlaceId,
		X:       float64(req.X),
		Y:       float64(req.Y),
		Width:   float64(req.Width),
		Height:  float64(req.Height),
		Shape:   model.PlaceMarkerShape(req.Shape),
	}
}
