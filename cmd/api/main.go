package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/makdoom/go-backend/internal/config"
	"github.com/makdoom/go-backend/internal/db"
	"github.com/makdoom/go-backend/internal/handlers"
)

func main() {
	// Load configurations
	cfg := config.MustLoad()

	// Custom logger
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	// Database connection
	db, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}

	log.Printf("✅ Database connected successfully")

	// Dependency Injection
	lh := handlers.NewListingHandler(db, logger)

	// Mux router
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.Health)
	mux.HandleFunc("GET /listings", lh.List)
	mux.HandleFunc("DELETE /listing/{id}", lh.DeleteListing)

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
