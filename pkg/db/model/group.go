package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"not null;uniqueIndex;size:100"`
	Description *string   `gorm:"type:text"`
	IsSystem    bool      `gorm:"not null;default:false"` // True for immutable "system" group
	IsDefault   bool      `gorm:"not null;default:false"` // True for "default" group
	CreatedAt   time.Time `gorm:"not null;default:now()"`
	UpdatedAt   time.Time `gorm:"not null"`

	// Relations
	Users []*User `gorm:"many2many:user_groups;"`
}

func (Group) TableName() string { return "groups" }

func (m *Group) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *Group) Exists() bool {
	return m != nil && m.ID != uuid.Nil
}
