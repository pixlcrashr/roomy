package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationStatus string

const (
	ReservationStatusPending   ReservationStatus = "pending"
	ReservationStatusConfirmed ReservationStatus = "confirmed"
	ReservationStatusCheckedIn ReservationStatus = "checkedIn"
	ReservationStatusCancelled ReservationStatus = "cancelled"
)

type Reservation struct {
	ID               uuid.UUID         `gorm:"type:uuid;primaryKey"`
	PlaceID          uuid.UUID         `gorm:"type:uuid;not null;index"`
	UserID           uuid.UUID         `gorm:"type:uuid;not null;index"`
	StartTime        time.Time         `gorm:"not null;index"`
	EndTime          time.Time         `gorm:"not null;index"`
	Status           ReservationStatus `gorm:"not null;size:20;default:'pending'"`
	CheckInTime      *time.Time        `gorm:"type:timestamp"`
	CancelReason     *string           `gorm:"type:text"`
	CancelTime       *time.Time        `gorm:"type:timestamp"`
	IsRecurring      bool              `gorm:"not null;default:false"`
	RecurringGroupID *uuid.UUID        `gorm:"type:uuid;index"`
	CreatedAt        time.Time         `gorm:"not null;default:now()"`
	UpdatedAt        time.Time         `gorm:"not null"`

	// Relations
	Place *Place `gorm:"foreignKey:PlaceID"`
	User  *User  `gorm:"foreignKey:UserID"`
}

func (Reservation) TableName() string { return "reservations" }

func (m *Reservation) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *Reservation) Exists() bool {
	return m != nil && m.ID != uuid.Nil
}
