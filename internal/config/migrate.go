package config

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func MigrateDatabase(fs embed.FS) (*sql.DB, error) {
	maxOpenConnections := 3
	connMaxxTime := time.Hour
	db, err := sql.Open("sqlite3", "file:test.db?cache=shared&mode=memory")
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(connMaxxTime)
	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(maxOpenConnections)

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, fmt.Errorf("error creating db driver: %w", err)
	}

	d, err := iofs.New(fs, "resources/migrations")
	if err != nil {
		return nil, fmt.Errorf("error getting the migration files: %w", err)
	}
	m, err := migrate.NewWithInstance("iofs", d,
		"test", driver)
	if err != nil {
		return nil, fmt.Errorf("error creating the migration instance: %w", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("error migrating database: %w", err)
	}

	return db, nil
}
