package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Update the migration source path
	m, err := migrate.New(
		"file://./migrations",
		"postgres://postgres:postgres@localhost:5433/ecom-1?sslmode=disable",
	)
	if err != nil {
		log.Fatalf("Failed to initialize migration: %v", err)
	}

	// Check for user command
	if len(os.Args) < 2 {
		log.Fatalf("Please specify a command: 'up' or 'down'")
	}

	cmd := os.Args[1]

	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
		log.Println("Migrations applied successfully.")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to roll back migrations: %v", err)
		}
		log.Println("Migrations rolled back successfully.")
	default:
		log.Fatalf("Unknown command: %s. Use 'up' or 'down'.", cmd)
	}
}
