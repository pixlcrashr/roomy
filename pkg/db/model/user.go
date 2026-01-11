package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email          string    `gorm:"not null;uniqueIndex;size:255"`
	Username       string    `gorm:"not null;size:100"`
	Name           string    `gorm:"not null;size:255"`
	ProfilePicture *string   `gorm:"size:512"`
	OAuthProvider  string    `gorm:"not null;size:50;default:'gitlab'"`
	OAuthID        string    `gorm:"not null;size:255;index"`
	IsActive       bool      `gorm:"not null;default:true"`
	CreatedAt      time.Time `gorm:"not null;default:now()"`
	UpdatedAt      time.Time `gorm:"not null"`

	// Notification Preferences (embedded)
	NotifyReservationConfirmed bool `gorm:"not null;default:true"`
	NotifyReservationCancelled bool `gorm:"not null;default:true"`
	NotifyReservationReminder  bool `gorm:"not null;default:true"`
	ReminderMinutesBefore      int  `gorm:"not null;default:15"`
	NotifyCheckInWarning       bool `gorm:"not null;default:true"`

	// Relations
	Groups         []*Group       `gorm:"many2many:user_groups;"`
	Reservations   []*Reservation `gorm:"foreignKey:UserID"`
	FavoritePlaces []*Place       `gorm:"many2many:user_favorites;"`
	APIKeys        []*APIKey      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (User) TableName() string { return "users" }

func (m *User) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *User) Exists() bool {
	return m != nil && m.ID != uuid.Nil
}
