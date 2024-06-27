package config

import (
	"ZeroAgencyTest/lib/logger/sl"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

type ProjectMode struct {
	Env string `envconfig:"ENV" env-default:"development"`
}

func LoadProjectMode() string {
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Failed to load .env file: ", sl.Err(err))
	}

	return os.Getenv("ENV")
}
