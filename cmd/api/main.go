package main

import (
	"log"
	"net/http"
	"time"

	"github.com/makdoom/go-backend/internal/config"
	"github.com/makdoom/go-backend/internal/db"
	"github.com/makdoom/go-backend/internal/handlers"
)

func main() {
	cfg := config.MustLoad()

	_, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}

	log.Printf("✅ Database connected successfully")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.Health)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("💻 Server up and running at %v", "http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
