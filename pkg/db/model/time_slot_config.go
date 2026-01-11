package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TimeSlotConfig struct {
	ID               uuid.UUID  `gorm:"type:uuid;primaryKey"`
	PlaceID          uuid.UUID  `gorm:"type:uuid;not null;uniqueIndex"`
	IntervalMinutes  int        `gorm:"not null;default:30"`
	EarliestStartTime *string   `gorm:"size:10"` // "HH:MM" format
	LatestEndTime    *string    `gorm:"size:10"` // "HH:MM" format
	CreatedAt        time.Time  `gorm:"not null;default:now()"`
	UpdatedAt        time.Time  `gorm:"not null"`
}

func (TimeSlotConfig) TableName() string { return "time_slot_configs" }

func (m *TimeSlotConfig) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *TimeSlotConfig) Exists() bool {
	return m != nil && m.ID != uuid.Nil
}
