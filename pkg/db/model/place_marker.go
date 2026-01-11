package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceMarkerShape string

const (
	PlaceMarkerShapeRectangle PlaceMarkerShape = "rectangle"
	PlaceMarkerShapeCircle    PlaceMarkerShape = "circle"
)

type PlaceMarker struct {
	ID        uuid.UUID        `gorm:"type:uuid;primaryKey"`
	AreaID    uuid.UUID        `gorm:"type:uuid;not null;index"`
	PlaceID   uuid.UUID        `gorm:"type:uuid;not null;index"`
	X         float64          `gorm:"not null"`
	Y         float64          `gorm:"not null"`
	Width     float64          `gorm:"not null"`
	Height    float64          `gorm:"not null"`
	Shape     PlaceMarkerShape `gorm:"not null;size:20"`
	CreatedAt time.Time        `gorm:"not null;default:now()"`
	UpdatedAt time.Time        `gorm:"not null"`

	// Relations
	Area  *Area  `gorm:"foreignKey:AreaID"`
	Place *Place `gorm:"foreignKey:PlaceID"`
}

func (PlaceMarker) TableName() string { return "place_markers" }

func (m *PlaceMarker) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *PlaceMarker) Exists() bool {
	return m != nil && m.ID != uuid.Nil
}
