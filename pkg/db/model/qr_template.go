package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QRTemplate struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name         string    `gorm:"not null;size:255"`
	HTMLTemplate string    `gorm:"not null;type:text"`
	CreatedAt    time.Time `gorm:"not null;default:now()"`
	UpdatedAt    time.Time `gorm:"not null"`
}

func (QRTemplate) TableName() string { return "qr_templates" }

func (m *QRTemplate) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *QRTemplate) Exists() bool {
	return m != nil && m.ID != uuid.Nil
}
