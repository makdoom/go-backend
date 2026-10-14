package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/makdoom/go-backend/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: migrate <up | down>")
	}

	cfg := config.MustLoad()

	m, err := migrate.New("file://migrations", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to load DB migrations: %v", err)
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatalf("Error running UP migrations: %v\n", err)
			return
		}
		log.Println("Running UP migrations")
	case "down":
		if err := m.Down(); err != nil {
			log.Fatalf("Error running DOWN migrations: %v\n", err)
			return
		}
		log.Println("Running DOWN migrations")
	default:
		log.Fatalf("Unknown command '%s',\nUsage: migrate <up | down>", os.Args[1])
	}
}
