package main

import (
	"log"
	"wifi-radar-go/internal/config"
	"wifi-radar-go/internal/database"
	"wifi-radar-go/internal/router"
)

func main() {
	cfg := config.NewConfig()

	db := database.Connect(cfg.Database)

	rtr := router.SetupRoutes(db)

	if err := rtr.Run(cfg.Server.Address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
