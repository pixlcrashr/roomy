package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type APIKey struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey"`
	UserID     uuid.UUID  `gorm:"type:uuid;not null;index"`
	Name       string     `gorm:"not null;size:255"`
	Value      string     `gorm:"not null;size:128"`
	LastUsedAt *time.Time `gorm:"type:timestamp"`
	ExpiresAt  *time.Time `gorm:"type:timestamp"`
	CreatedAt  time.Time  `gorm:"not null;default:now()"`
	UpdatedAt  time.Time  `gorm:"not null"`

	// Relations
	User *User `gorm:"foreignKey:UserID"`
}

func (APIKey) TableName() string { return "api_keys" }

func (m *APIKey) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *APIKey) Exists() bool {
	return m != nil && m.ID != uuid.Nil
}
