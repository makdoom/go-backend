package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
}

func MustLoad() Config {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		panic("Port is not configured")
	}

	env := os.Getenv("ENV")
	if env == "" {
		panic("Env is not configured")
	}

	return Config{
		Port: port,
		Env:  env,
	}

}
