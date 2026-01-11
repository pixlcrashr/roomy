package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Area struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	BuildingID  uuid.UUID `gorm:"type:uuid;not null;index"`
	Name        string    `gorm:"not null;size:255"`
	Description *string   `gorm:"type:text"`
	Location    *string   `gorm:"size:255"`
	RoomPlanURL *string   `gorm:"size:1024"`
	CreatedAt   time.Time `gorm:"not null;default:now()"`
	UpdatedAt   time.Time `gorm:"not null"`

	// Relations
	Building     *Building      `gorm:"foreignKey:BuildingID"`
	Places       []*Place       `gorm:"foreignKey:AreaID;constraint:OnDelete:CASCADE"`
	PlaceMarkers []*PlaceMarker `gorm:"foreignKey:AreaID;constraint:OnDelete:CASCADE"`
	Blockings    []*Blocking    `gorm:"polymorphic:Entity;polymorphicValue:area"`
}

func (Area) TableName() string { return "areas" }

func (m *Area) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *Area) Exists() bool {
	return m != nil && m.ID != uuid.Nil
}
