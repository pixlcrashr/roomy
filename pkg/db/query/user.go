package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

type UserQuery interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)

	// SELECT * FROM @@table WHERE email = @email
	GetByEmail(ctx context.Context, email string) (*model.User, error)

	// SELECT * FROM @@table WHERE oauth_provider = @provider AND oauth_id = @oauthID
	GetByOAuth(ctx context.Context, provider string, oauthID string) (*model.User, error)

	// SELECT * FROM @@table
	// {{where}}
	//   {{if search != nil}} (name ILIKE @search OR email ILIKE @search OR username ILIKE @search) {{end}}
	//   {{if isActive != nil}} AND is_active = @isActive {{end}}
	// {{end}}
	// ORDER BY created_at DESC
	// LIMIT @limit OFFSET @offset
	List(ctx context.Context, limit int, offset int, search *string, isActive *bool) ([]*model.User, error)

	// SELECT COUNT(*) FROM @@table
	// {{where}}
	//   {{if search != nil}} (name ILIKE @search OR email ILIKE @search OR username ILIKE @search) {{end}}
	//   {{if isActive != nil}} AND is_active = @isActive {{end}}
	// {{end}}
	CountAll(ctx context.Context, search *string, isActive *bool) (int64, error)

	// INSERT INTO @@table (
	//   id, email, username, name, profile_picture, oauth_provider, oauth_id,
	//   is_active, notify_reservation_confirmed, notify_reservation_cancelled,
	//   notify_reservation_reminder, reminder_minutes_before, notify_check_in_warning,
	//   created_at, updated_at
	// ) VALUES (
	//   @id, @email, @username, @name, @profilePicture, @oauthProvider, @oauthID,
	//   true, true, true, true, 15, true, NOW(), NOW()
	// )
	Insert(
		ctx context.Context,
		id uuid.UUID,
		email string,
		username string,
		name string,
		profilePicture *string,
		oauthProvider string,
		oauthID string,
	) error

	// UPDATE @@table
	// {{set}}
	//   {{if name != nil}} name = @name, {{end}}
	//   {{if username != nil}} username = @username, {{end}}
	//   {{if profilePicture != nil}} profile_picture = @profilePicture, {{end}}
	//   updated_at = NOW()
	// {{end}}
	// WHERE id = @id
	Save(ctx context.Context, id uuid.UUID, name *string, username *string, profilePicture *string) error

	// UPDATE @@table
	// {{set}}
	//   {{if notifyReservationConfirmed != nil}} notify_reservation_confirmed = @notifyReservationConfirmed, {{end}}
	//   {{if notifyReservationCancelled != nil}} notify_reservation_cancelled = @notifyReservationCancelled, {{end}}
	//   {{if notifyReservationReminder != nil}} notify_reservation_reminder = @notifyReservationReminder, {{end}}
	//   {{if reminderMinutesBefore != nil}} reminder_minutes_before = @reminderMinutesBefore, {{end}}
	//   {{if notifyCheckInWarning != nil}} notify_check_in_warning = @notifyCheckInWarning, {{end}}
	//   updated_at = NOW()
	// {{end}}
	// WHERE id = @id
	UpdateNotificationPreferences(
		ctx context.Context,
		id uuid.UUID,
		notifyReservationConfirmed *bool,
		notifyReservationCancelled *bool,
		notifyReservationReminder *bool,
		reminderMinutesBefore *int,
		notifyCheckInWarning *bool,
	) error

	// UPDATE @@table SET is_active = true, updated_at = NOW() WHERE id = @id
	Enable(ctx context.Context, id uuid.UUID) error

	// UPDATE @@table SET is_active = false, updated_at = NOW() WHERE id = @id
	Disable(ctx context.Context, id uuid.UUID) error

	// DELETE FROM @@table WHERE id = @id
	Remove(ctx context.Context, id uuid.UUID) error
}
