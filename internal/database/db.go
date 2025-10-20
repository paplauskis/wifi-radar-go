package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"wifi-radar-go/internal/config"
)

func Connect(cfg config.DatabaseConfig) *sql.DB {
	db, err := sql.Open(cfg.DatabaseDriver, cfg.DatabaseConnectionString)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Printf("Connected to %s database", cfg.DatabaseDriver)
	return db
}
