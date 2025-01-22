package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AshFire1/config"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func NewPostgresStorage() (*sql.DB, error) {
	// Construct the PostgreSQL connection string
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Envs.Host, config.Envs.Port, config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBName, config.Envs.SSLMode,
	)

	// Open the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Error opening database connection:", err)
		return nil, err
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		log.Println("Error connecting to database:", err)
		return nil, err
	}

	return db, nil
}
