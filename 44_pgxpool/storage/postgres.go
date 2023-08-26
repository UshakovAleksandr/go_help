package storage

import (
	"context"
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	// import db migrations engine.
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewDB - Конструктор, принимает строку подключения к БД.
func NewDB() (*pgxpool.Pool, error) {
	dsn := "postgres://user:password@localhost:5432/db?sslmode=disable"

	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(context.Background()); err != nil {
		return nil, err
	}

	log.Println("✓ Successful connection to the database...")

	makeMigrations(dsn, "down")
	makeMigrations(dsn, "up")

	return db, nil
}

// makeMigrationsUp - make db migrations.
func makeMigrations(dsn string, flag string) {
	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if flag == "up" {
		err = m.Up()
		if err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Printf("✓ No changes in database...\n")
				return
			}
			log.Fatal(err)
		}
		log.Printf("✓ Database UP migration is done...\n")
	} else {
		err = m.Down()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("✓ Database DOWN migration is done...\n")
	}
}
