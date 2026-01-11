package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Building struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"not null;size:255"`
	Description *string   `gorm:"type:text"`
	Location    *string   `gorm:"size:255"`
	CreatedAt   time.Time `gorm:"not null;default:now()"`
	UpdatedAt   time.Time `gorm:"not null"`

	// Relations
	Areas     []*Area     `gorm:"foreignKey:BuildingID;constraint:OnDelete:CASCADE"`
	Blockings []*Blocking `gorm:"polymorphic:Entity;polymorphicValue:building"`
}

func (Building) TableName() string { return "buildings" }

func (m *Building) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *Building) Exists() bool {
	return m != nil && m.ID != uuid.Nil
}
