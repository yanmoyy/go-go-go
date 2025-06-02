// cmd/api/main.go
package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/yanmoyy/go-go-go/internal/api"
	"github.com/yanmoyy/go-go-go/internal/config"
	"github.com/yanmoyy/go-go-go/internal/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbConn, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	dbQueries := database.New(dbConn)

	server := api.NewServer(cfg, dbQueries)

	log.Printf("Starting server on port %s", cfg.Port)
	if err := server.Start(); err != nil {
		log.Fatal("Server Error: %w", err)
	}
}
