package migrations

import (
	"database/sql"
	"errors"
	"fmt"
	"io/fs"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// Run executes database migrations using the embedded PostgreSQL migration files.
// It applies all pending up migrations.
func Run(db *sql.DB) error {
	m, err := newMigrator(db)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}

// Rollback rolls back the last applied migration.
func Rollback(db *sql.DB) error {
	m, err := newMigrator(db)
	if err != nil {
		return err
	}

	if err := m.Steps(-1); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("rollback failed: %w", err)
	}

	return nil
}

// RollbackAll rolls back all applied migrations.
func RollbackAll(db *sql.DB) error {
	m, err := newMigrator(db)
	if err != nil {
		return err
	}

	if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("rollback all failed: %w", err)
	}

	return nil
}

// Version returns the current migration version and dirty state.
func Version(db *sql.DB) (uint, bool, error) {
	m, err := newMigrator(db)
	if err != nil {
		return 0, false, err
	}

	version, dirty, err := m.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		return 0, false, fmt.Errorf("failed to get version: %w", err)
	}

	return version, dirty, nil
}

// Steps applies n migrations. If n > 0, applies up migrations. If n < 0, applies down migrations.
func Steps(db *sql.DB, n int) error {
	m, err := newMigrator(db)
	if err != nil {
		return err
	}

	if err := m.Steps(n); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("steps migration failed: %w", err)
	}

	return nil
}

// Force sets a specific migration version without running migrations.
// This is useful for fixing a dirty state.
func Force(db *sql.DB, version int) error {
	m, err := newMigrator(db)
	if err != nil {
		return err
	}

	if err := m.Force(version); err != nil {
		return fmt.Errorf("force version failed: %w", err)
	}

	return nil
}

// newMigrator creates a new migrate instance using the embedded PostgreSQL migrations.
func newMigrator(db *sql.DB) (*migrate.Migrate, error) {
	// Create a sub-filesystem for the postgresql directory
	subFS, err := fs.Sub(postgresFS, "postgresql")
	if err != nil {
		return nil, fmt.Errorf("failed to create sub filesystem: %w", err)
	}

	// Create the source driver from the embedded filesystem
	sourceDriver, err := iofs.New(subFS, ".")
	if err != nil {
		return nil, fmt.Errorf("failed to create source driver: %w", err)
	}

	// Create the database driver
	dbDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create database driver: %w", err)
	}

	// Create the migrator
	m, err := migrate.NewWithInstance("iofs", sourceDriver, "postgres", dbDriver)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrator: %w", err)
	}

	return m, nil
}
