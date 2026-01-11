package database

import (
	"fmt"

	"github.com/pixlcrashr/roomy/pkg/db/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Building{},
		&model.Area{},
		&model.Place{},
		&model.PlaceMarker{},
		&model.Blocking{},
		&model.Reservation{},
		&model.User{},
		&model.Group{},
		&model.Equipment{},
		&model.APIKey{},
		&model.AuditLogEntry{},
		&model.QRTemplate{},
	)
}
