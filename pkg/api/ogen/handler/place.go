package handler

import (
	"context"

	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"github.com/pixlcrashr/roomy/pkg/api/ogen/handler/converter"
	"github.com/pixlcrashr/roomy/pkg/db/model"
	"gorm.io/gorm"
)

// PlaceHandler handles place-related operations.
type PlaceHandler struct {
	db *gorm.DB
}

// NewPlaceHandler creates a new PlaceHandler.
func NewPlaceHandler(db *gorm.DB) *PlaceHandler {
	return &PlaceHandler{db: db}
}

// CreatePlace creates a new place.
// POST /places
func (h *PlaceHandler) CreatePlace(ctx context.Context, req *gen.CreatePlaceRequest) (gen.CreatePlaceRes, error) {
	place := converter.CreatePlaceRequestToModel(req)
	if err := h.db.WithContext(ctx).Create(place).Error; err != nil {
		return nil, err
	}
	return converter.PlaceToAPI(place), nil
}

// GetPlace gets place details.
// GET /places/{placeId}
func (h *PlaceHandler) GetPlace(ctx context.Context, params gen.GetPlaceParams) (gen.GetPlaceRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).Preload("Equipment").First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.GetPlaceNotFound{}, nil
		}
		return nil, err
	}
	return converter.PlaceToAPI(&place), nil
}

// UpdatePlace updates a place.
// PUT /places/{placeId}
func (h *PlaceHandler) UpdatePlace(ctx context.Context, req *gen.UpdatePlaceRequest, params gen.UpdatePlaceParams) (gen.UpdatePlaceRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.UpdatePlaceNotFound{}, nil
		}
		return nil, err
	}
	converter.UpdatePlaceRequestToModel(req, &place)
	if err := h.db.WithContext(ctx).Save(&place).Error; err != nil {
		return nil, err
	}
	return converter.PlaceToAPI(&place), nil
}

// DeletePlace deletes a place.
// DELETE /places/{placeId}
func (h *PlaceHandler) DeletePlace(ctx context.Context, params gen.DeletePlaceParams) (gen.DeletePlaceRes, error) {
	result := h.db.WithContext(ctx).Delete(&model.Place{}, "id = ?", params.PlaceId)
	if result.Error != nil {
		return &gen.Error{Message: result.Error.Error()}, nil
	}
	if result.RowsAffected == 0 {
		return &gen.DeletePlaceNotFound{}, nil
	}
	return &gen.DeletePlaceNoContent{}, nil
}

// ListPlaces lists/searches places.
// GET /places
func (h *PlaceHandler) ListPlaces(ctx context.Context, params gen.ListPlacesParams) (*gen.PaginatedPlaceList, error) {
	var places []model.Place
	var total int64

	query := h.db.WithContext(ctx).Model(&model.Place{})

	// Apply filters
	if params.AreaId.IsSet() {
		query = query.Where("area_id = ?", params.AreaId.Value)
	}
	if params.BuildingId.IsSet() {
		query = query.Joins("JOIN areas ON areas.id = places.area_id").Where("areas.building_id = ?", params.BuildingId.Value)
	}
	if params.Search.IsSet() {
		search := "%" + params.Search.Value + "%"
		query = query.Where("name ILIKE ? OR description ILIKE ?", search, search)
	}
	if params.MinCapacity.IsSet() {
		query = query.Where("seat_capacity >= ?", params.MinCapacity.Value)
	}
	if params.IsBookable.IsSet() {
		query = query.Where("is_bookable = ?", params.IsBookable.Value)
	}

	// Count total
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Apply pagination
	limit := 20
	offset := 0
	if params.Limit.IsSet() {
		limit = params.Limit.Value
	}
	if params.Offset.IsSet() {
		offset = params.Offset.Value
	}

	if err := query.Limit(limit).Offset(offset).Preload("Equipment").Find(&places).Error; err != nil {
		return nil, err
	}

	return &gen.PaginatedPlaceList{
		Items: converter.PlacesToAPI(places),
		Total: int(total),
	}, nil
}

// GetPlaceAvailability returns available time slots for this place.
// GET /places/{placeId}/availability
func (h *PlaceHandler) GetPlaceAvailability(ctx context.Context, params gen.GetPlaceAvailabilityParams) (gen.GetPlaceAvailabilityRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.GetPlaceAvailabilityNotFound{}, nil
		}
		return nil, err
	}

	// TODO: Calculate availability based on blocking and reservations
	return &gen.AvailabilityResponse{
		TimeSlots: []gen.TimeSlot{},
	}, nil
}

// GetPlaceBlocking returns blocking periods for this place.
// GET /places/{placeId}/blocking
func (h *PlaceHandler) GetPlaceBlocking(ctx context.Context, params gen.GetPlaceBlockingParams) (gen.GetPlaceBlockingRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).Preload("Area").First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.GetPlaceBlockingNotFound{}, nil
		}
		return nil, err
	}

	var blockings []model.Blocking
	query := h.db.WithContext(ctx).Where("entity_type = ? AND entity_id = ?", "place", params.PlaceId)

	// Include inherited blocking from area and building
	if params.IncludeInherited.IsSet() && params.IncludeInherited.Value {
		query = h.db.WithContext(ctx).Where(
			"(entity_type = ? AND entity_id = ?) OR (entity_type = ? AND entity_id = ?) OR (entity_type = ? AND entity_id = ?)",
			"place", params.PlaceId,
			"area", place.AreaID,
			"building", place.Area.BuildingID,
		)
	}

	if err := query.Find(&blockings).Error; err != nil {
		return nil, err
	}

	return &gen.BlockingResponse{
		Entries: converter.BlockingsToAPI(blockings),
	}, nil
}

// GetPlaceCalendar returns .ics file with blocking periods and reservations.
// GET /places/{placeId}/calendar.ics
func (h *PlaceHandler) GetPlaceCalendar(ctx context.Context, params gen.GetPlaceCalendarParams) (gen.GetPlaceCalendarRes, error) {
	// TODO: Generate ICS calendar
	return &gen.GetPlaceCalendarOKHeaders{}, nil
}

// AddPlaceBlockingEntries adds blocking periods to place.
// POST /places/{placeId}/blocking/entries
func (h *PlaceHandler) AddPlaceBlockingEntries(ctx context.Context, req *gen.AddPlaceBlockingEntriesReq, params gen.AddPlaceBlockingEntriesParams) (gen.AddPlaceBlockingEntriesRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.AddPlaceBlockingEntriesNotFound{}, nil
		}
		return nil, err
	}

	var blockings []model.Blocking
	for _, entry := range req.Entries {
		blocking := converter.BlockingEntryToModel(&entry, "place", params.PlaceId)
		blockings = append(blockings, *blocking)
	}

	if err := h.db.WithContext(ctx).Create(&blockings).Error; err != nil {
		return nil, err
	}

	return &gen.BlockingResponse{
		Entries: converter.BlockingsToAPI(blockings),
	}, nil
}

// RemovePlaceBlockingEntries removes blocking periods by IDs.
// DELETE /places/{placeId}/blocking/entries
func (h *PlaceHandler) RemovePlaceBlockingEntries(ctx context.Context, req *gen.RemovePlaceBlockingEntriesReq, params gen.RemovePlaceBlockingEntriesParams) (gen.RemovePlaceBlockingEntriesRes, error) {
	result := h.db.WithContext(ctx).Where("entity_type = ? AND entity_id = ? AND id IN ?", "place", params.PlaceId, req.Ids).Delete(&model.Blocking{})
	if result.Error != nil {
		return &gen.Error{Message: result.Error.Error()}, nil
	}
	return &gen.RemovePlaceBlockingEntriesNoContent{}, nil
}

// ReplacePlaceBlocking replaces all place blocking periods.
// PUT /places/{placeId}/blocking
func (h *PlaceHandler) ReplacePlaceBlocking(ctx context.Context, req *gen.ReplacePlaceBlockingReq, params gen.ReplacePlaceBlockingParams) (gen.ReplacePlaceBlockingRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.ReplacePlaceBlockingNotFound{}, nil
		}
		return nil, err
	}

	// Delete existing
	if err := h.db.WithContext(ctx).Where("entity_type = ? AND entity_id = ?", "place", params.PlaceId).Delete(&model.Blocking{}).Error; err != nil {
		return nil, err
	}

	// Create new
	var blockings []model.Blocking
	for _, entry := range req.Entries {
		blocking := converter.BlockingEntryToModel(&entry, "place", params.PlaceId)
		blockings = append(blockings, *blocking)
	}

	if len(blockings) > 0 {
		if err := h.db.WithContext(ctx).Create(&blockings).Error; err != nil {
			return nil, err
		}
	}

	return &gen.BlockingResponse{
		Entries: converter.BlockingsToAPI(blockings),
	}, nil
}

// GetPlaceConstraints gets place constraints.
// GET /places/{placeId}/constraints
func (h *PlaceHandler) GetPlaceConstraints(ctx context.Context, params gen.GetPlaceConstraintsParams) (gen.GetPlaceConstraintsRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.GetPlaceConstraintsNotFound{}, nil
		}
		return nil, err
	}
	return converter.PlaceConstraintsToAPI(&place), nil
}

// UpdatePlaceConstraints updates place constraints.
// PUT /places/{placeId}/constraints
func (h *PlaceHandler) UpdatePlaceConstraints(ctx context.Context, req *gen.PlaceConstraints, params gen.UpdatePlaceConstraintsParams) (gen.UpdatePlaceConstraintsRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.UpdatePlaceConstraintsNotFound{}, nil
		}
		return nil, err
	}

	converter.UpdatePlaceConstraintsToModel(req, &place)

	if err := h.db.WithContext(ctx).Save(&place).Error; err != nil {
		return nil, err
	}

	return converter.PlaceConstraintsToAPI(&place), nil
}

// GetPlaceTimeSlots returns the booking grid configuration for this place.
// GET /places/{placeId}/timeSlots
func (h *PlaceHandler) GetPlaceTimeSlots(ctx context.Context, params gen.GetPlaceTimeSlotsParams) (gen.GetPlaceTimeSlotsRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.GetPlaceTimeSlotsNotFound{}, nil
		}
		return nil, err
	}
	return converter.TimeSlotConfigToAPI(&place), nil
}

// UpdatePlaceTimeSlots updates time slot configuration.
// PUT /places/{placeId}/timeSlots
func (h *PlaceHandler) UpdatePlaceTimeSlots(ctx context.Context, req *gen.TimeSlotConfig, params gen.UpdatePlaceTimeSlotsParams) (gen.UpdatePlaceTimeSlotsRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.UpdatePlaceTimeSlotsNotFound{}, nil
		}
		return nil, err
	}

	converter.UpdateTimeSlotConfigToModel(req, &place)

	if err := h.db.WithContext(ctx).Save(&place).Error; err != nil {
		return nil, err
	}

	return converter.TimeSlotConfigToAPI(&place), nil
}

// GetPlaceEquipment gets place equipment/amenities.
// GET /places/{placeId}/equipment
func (h *PlaceHandler) GetPlaceEquipment(ctx context.Context, params gen.GetPlaceEquipmentParams) (gen.GetPlaceEquipmentRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).Preload("Equipment").First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.GetPlaceEquipmentNotFound{}, nil
		}
		return nil, err
	}

	return &gen.GetPlaceEquipmentOKApplicationJSON(converter.EquipmentsToAPI(place.Equipment)), nil
}

// AddPlaceEquipment adds equipment to place.
// POST /places/{placeId}/equipment
func (h *PlaceHandler) AddPlaceEquipment(ctx context.Context, req *gen.AddPlaceEquipmentReq, params gen.AddPlaceEquipmentParams) (gen.AddPlaceEquipmentRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.AddPlaceEquipmentNotFound{}, nil
		}
		return nil, err
	}

	// Load equipment by IDs
	var equipment []model.Equipment
	if err := h.db.WithContext(ctx).Where("id IN ?", req.EquipmentIds).Find(&equipment).Error; err != nil {
		return nil, err
	}

	// Add equipment association
	if err := h.db.WithContext(ctx).Model(&place).Association("Equipment").Append(&equipment); err != nil {
		return nil, err
	}

	// Reload with equipment
	if err := h.db.WithContext(ctx).Preload("Equipment").First(&place, "id = ?", params.PlaceId).Error; err != nil {
		return nil, err
	}

	return &gen.AddPlaceEquipmentOKApplicationJSON(converter.EquipmentsToAPI(place.Equipment)), nil
}

// RemovePlaceEquipment removes equipment from place.
// DELETE /places/{placeId}/equipment
func (h *PlaceHandler) RemovePlaceEquipment(ctx context.Context, req *gen.RemovePlaceEquipmentReq, params gen.RemovePlaceEquipmentParams) (gen.RemovePlaceEquipmentRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.RemovePlaceEquipmentNotFound{}, nil
		}
		return nil, err
	}

	// Load equipment by IDs
	var equipment []model.Equipment
	if err := h.db.WithContext(ctx).Where("id IN ?", req.EquipmentIds).Find(&equipment).Error; err != nil {
		return nil, err
	}

	// Remove equipment association
	if err := h.db.WithContext(ctx).Model(&place).Association("Equipment").Delete(&equipment); err != nil {
		return nil, err
	}

	return &gen.RemovePlaceEquipmentNoContent{}, nil
}

// GetPlaceWhitelist gets user whitelist.
// GET /places/{placeId}/whitelist
func (h *PlaceHandler) GetPlaceWhitelist(ctx context.Context, params gen.GetPlaceWhitelistParams) (gen.GetPlaceWhitelistRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).Preload("WhitelistUsers").First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.GetPlaceWhitelistNotFound{}, nil
		}
		return nil, err
	}

	return &gen.GetPlaceWhitelistOKApplicationJSON(converter.UserReferencesToAPI(place.WhitelistUsers)), nil
}

// AddPlaceWhitelistUsers adds users to whitelist.
// POST /places/{placeId}/whitelist
func (h *PlaceHandler) AddPlaceWhitelistUsers(ctx context.Context, req *gen.AddPlaceWhitelistUsersReq, params gen.AddPlaceWhitelistUsersParams) (gen.AddPlaceWhitelistUsersRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.AddPlaceWhitelistUsersNotFound{}, nil
		}
		return nil, err
	}

	// Load users by IDs
	var users []model.User
	if err := h.db.WithContext(ctx).Where("id IN ?", req.UserIds).Find(&users).Error; err != nil {
		return nil, err
	}

	// Add user association
	if err := h.db.WithContext(ctx).Model(&place).Association("WhitelistUsers").Append(&users); err != nil {
		return nil, err
	}

	// Reload with whitelist
	if err := h.db.WithContext(ctx).Preload("WhitelistUsers").First(&place, "id = ?", params.PlaceId).Error; err != nil {
		return nil, err
	}

	return &gen.AddPlaceWhitelistUsersOKApplicationJSON(converter.UserReferencesToAPI(place.WhitelistUsers)), nil
}

// RemovePlaceWhitelistUsers removes users from whitelist.
// DELETE /places/{placeId}/whitelist
func (h *PlaceHandler) RemovePlaceWhitelistUsers(ctx context.Context, req *gen.RemovePlaceWhitelistUsersReq, params gen.RemovePlaceWhitelistUsersParams) (gen.RemovePlaceWhitelistUsersRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.RemovePlaceWhitelistUsersNotFound{}, nil
		}
		return nil, err
	}

	// Load users by IDs
	var users []model.User
	if err := h.db.WithContext(ctx).Where("id IN ?", req.UserIds).Find(&users).Error; err != nil {
		return nil, err
	}

	// Remove user association
	if err := h.db.WithContext(ctx).Model(&place).Association("WhitelistUsers").Delete(&users); err != nil {
		return nil, err
	}

	return &gen.RemovePlaceWhitelistUsersNoContent{}, nil
}

// GetPlaceQrCode generates a QR code.
// GET /places/{placeId}/qrCode
func (h *PlaceHandler) GetPlaceQrCode(ctx context.Context, params gen.GetPlaceQrCodeParams) (gen.GetPlaceQrCodeRes, error) {
	var place model.Place
	if err := h.db.WithContext(ctx).First(&place, "id = ?", params.PlaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gen.GetPlaceQrCodeNotFound{}, nil
		}
		return nil, err
	}

	// TODO: Generate QR code
	return &gen.GetPlaceQrCodeOKHeaders{}, nil
}
