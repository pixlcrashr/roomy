package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Example struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"not null"`
	Description *string
	CreatedAt   time.Time `gorm:"not null;default:now()"`
	UpdatedAt   time.Time `gorm:"not null"`
}

// TableName overrides the default table name.
func (Example) TableName() string { return "examples" }

// BeforeCreate ensures an ID is set using google/uuid when creating the record.
func (m *Example) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *Example) Exists() bool {
	return m != nil && m.ID != uuid.Nil
}
