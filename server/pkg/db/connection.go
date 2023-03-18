package db

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectAndMigrate(port int) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", port, "postgres", "admin", "postgres"))

	if err != nil {
		return nil, err
	}
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/db/migrations",
		"postgres", driver)
	if err != nil {
		return nil, err
	}

	err = m.Up()
	if !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	return db, nil
}
