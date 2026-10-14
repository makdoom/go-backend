package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Env         string
	DatabaseURL string
}

func MustLoad() Config {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		panic("Port is not configured")
	}

	env := os.Getenv("ENV")
	if env == "" {
		panic("Environment is not configured")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		panic("Database URL is not configured")
	}

	return Config{
		Port:        port,
		Env:         env,
		DatabaseURL: databaseURL,
	}

}
