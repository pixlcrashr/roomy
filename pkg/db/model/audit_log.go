package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuditAction string

const (
	AuditActionCreate AuditAction = "create"
	AuditActionUpdate AuditAction = "update"
	AuditActionDelete AuditAction = "delete"
)

type AuditLogEntry struct {
	ID         uuid.UUID   `gorm:"type:uuid;primaryKey"`
	EntityType string      `gorm:"not null;size:50;index"` // building, area, place, reservation, user, group
	EntityID   uuid.UUID   `gorm:"type:uuid;not null;index"`
	Action     AuditAction `gorm:"not null;size:20"`
	UserID     uuid.UUID   `gorm:"type:uuid;not null;index"`
	Changes    *string     `gorm:"type:jsonb"` // JSON with before/after values
	Timestamp  time.Time   `gorm:"not null;default:now();index"`

	// Relations
	User *User `gorm:"foreignKey:UserID"`
}

func (AuditLogEntry) TableName() string { return "audit_log" }

func (m *AuditLogEntry) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *AuditLogEntry) Exists() bool {
	return m != nil && m.ID != uuid.Nil
}
