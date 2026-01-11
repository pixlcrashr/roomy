package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Equipment struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"not null;uniqueIndex;size:100"`
	Icon        *string   `gorm:"size:50"`
	Description *string   `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"not null;default:now()"`
	UpdatedAt   time.Time `gorm:"not null"`

	// Relations
	Places []*Place `gorm:"many2many:place_equipment;"`
}

func (Equipment) TableName() string { return "equipment" }

func (m *Equipment) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *Equipment) Exists() bool {
	return m != nil && m.ID != uuid.Nil
}
