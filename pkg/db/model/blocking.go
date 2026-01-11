package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlockingType string

const (
	BlockingTypeClosedHours BlockingType = "closedHours"
	BlockingTypeWeekend     BlockingType = "weekend"
	BlockingTypeHoliday     BlockingType = "holiday"
	BlockingTypeMaintenance BlockingType = "maintenance"
	BlockingTypeEvent       BlockingType = "event"
	BlockingTypeDisabled    BlockingType = "disabled"
	BlockingTypeCustom      BlockingType = "custom"
)

type Blocking struct {
	ID                 uuid.UUID    `gorm:"type:uuid;primaryKey"`
	EntityType         string       `gorm:"not null;size:20;index"` // building, area, place
	EntityID           uuid.UUID    `gorm:"type:uuid;not null;index"`
	BlockingType       BlockingType `gorm:"not null;size:20"`
	Name               *string      `gorm:"size:255"`
	Reason             *string      `gorm:"type:text"`
	StartTime          time.Time    `gorm:"not null;index"`
	EndTime            time.Time    `gorm:"not null;index"`
	IsRecurring        bool         `gorm:"not null;default:false"`
	RecurrenceRule     *string      `gorm:"type:text"` // RRULE format
	RecurrenceDuration *int64       `gorm:"type:bigint"` // nanoseconds
	RecurrenceEnd      *time.Time   `gorm:"type:date"`
	CreatedAt          time.Time    `gorm:"not null;default:now()"`
	UpdatedAt          time.Time    `gorm:"not null"`
}

func (Blocking) TableName() string { return "blockings" }

func (m *Blocking) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *Blocking) Exists() bool {
	return m != nil && m.ID != uuid.Nil
}