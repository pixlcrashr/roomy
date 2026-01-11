package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookingMethod string

const (
	BookingMethodSelfService BookingMethod = "selfService"
	BookingMethodManual      BookingMethod = "manual"
)

type Place struct {
	ID              uuid.UUID     `gorm:"type:uuid;primaryKey"`
	AreaID          uuid.UUID     `gorm:"type:uuid;not null;index"`
	Name            string        `gorm:"not null;size:255"`
	Description     *string       `gorm:"type:text"`
	Location        *string       `gorm:"size:255"`
	Capacity        int           `gorm:"not null;default:1"`
	IsBookable      bool          `gorm:"not null;default:true"`
	BookingMethod   BookingMethod `gorm:"not null;default:'selfService';size:50"`
	IsDisabled      bool          `gorm:"not null;default:false"`
	RequiresCheckIn bool          `gorm:"not null;default:false"`
	CreatedAt       time.Time     `gorm:"not null;default:now()"`
	UpdatedAt       time.Time     `gorm:"not null"`

	// Relations
	Area             *Area           `gorm:"foreignKey:AreaID"`
	TimeSlotConfig   *TimeSlotConfig `gorm:"foreignKey:PlaceID;constraint:OnDelete:CASCADE"`
	Equipment        []*Equipment    `gorm:"many2many:place_equipment;"`
	WhitelistedUsers []*User         `gorm:"many2many:place_whitelist;"`
	Reservations     []*Reservation  `gorm:"foreignKey:PlaceID;constraint:OnDelete:CASCADE"`
	Blockings        []*Blocking     `gorm:"polymorphic:Entity;polymorphicValue:place"`
	FavoritedBy      []*User         `gorm:"many2many:user_favorites;"`
}

func (Place) TableName() string { return "places" }

func (m *Place) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *Place) Exists() bool {
	return m != nil && m.ID != uuid.Nil
}
