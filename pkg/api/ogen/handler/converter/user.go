package converter

import (
	"net/url"

	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

func UserToAPI(m *model.User) *gen.User {
	if m == nil {
		return nil
	}
	u := &gen.User{
		ID:            m.ID,
		Email:         m.Email,
		Username:      m.Username,
		Name:          m.Name,
		OauthProvider: m.OAuthProvider,
		IsActive:      m.IsActive,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
	}
	if m.ProfilePicture != nil {
		if parsed, err := url.Parse(*m.ProfilePicture); err == nil {
			u.ProfilePicture.SetTo(*parsed)
		}
	}
	if m.OAuthID != "" {
		u.OauthId.SetTo(m.OAuthID)
	}
	if len(m.Groups) > 0 {
		u.Groups = GroupsToAPI(m.Groups)
	}
	return u
}

func UsersToAPI(models []*model.User) []gen.User {
	result := make([]gen.User, len(models))
	for i, m := range models {
		result[i] = *UserToAPI(m)
	}
	return result
}

func UserReferenceToAPI(m *model.User) *gen.UserReference {
	if m == nil {
		return nil
	}
	return &gen.UserReference{
		ID:    m.ID,
		Email: m.Email,
		Name:  m.Name,
	}
}

func UserReferencesToAPI(models []model.User) []gen.UserReference {
	result := make([]gen.UserReference, len(models))
	for i, m := range models {
		result[i] = *UserReferenceToAPI(&m)
	}
	return result
}

func NotificationPreferencesToAPI(m *model.User) *gen.NotificationPreferences {
	if m == nil {
		return nil
	}
	return &gen.NotificationPreferences{
		ReservationConfirmed:  gen.NewOptBool(m.NotifyReservationConfirmed),
		ReservationCancelled:  gen.NewOptBool(m.NotifyReservationCancelled),
		ReservationReminder:   gen.NewOptBool(m.NotifyReservationReminder),
		ReminderMinutesBefore: gen.NewOptInt(m.ReminderMinutesBefore),
		CheckInWarning:        gen.NewOptBool(m.NotifyCheckInWarning),
	}
}

func UpdateNotificationPreferencesToModel(req *gen.NotificationPreferences, user *model.User) {
	if req == nil || user == nil {
		return
	}
	if req.ReservationConfirmed.IsSet() {
		user.NotifyReservationConfirmed = req.ReservationConfirmed.Value
	}
	if req.ReservationCancelled.IsSet() {
		user.NotifyReservationCancelled = req.ReservationCancelled.Value
	}
	if req.ReservationReminder.IsSet() {
		user.NotifyReservationReminder = req.ReservationReminder.Value
	}
	if req.ReminderMinutesBefore.IsSet() {
		user.ReminderMinutesBefore = req.ReminderMinutesBefore.Value
	}
	if req.CheckInWarning.IsSet() {
		user.NotifyCheckInWarning = req.CheckInWarning.Value
	}
}
