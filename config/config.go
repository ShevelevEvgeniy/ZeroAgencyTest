package config

import (
	"ZeroAgencyTest/lib/logger/sl"
	"github.com/kelseyhightower/envconfig"
	"log/slog"
	"os"
)

type Config struct {
	Mode       ProjectMode
	HTTPServer HTTPServer
	DB         DB
	Auth       Auth
}

func MustLoad(log *slog.Logger) *Config {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		log.Error("Failed to load config: ", sl.Err(err))
		os.Exit(1)
	}

	log.Info("Loaded config")
	log.Info("Mode.Env:", "value", cfg.Mode.Env)
	log.Info("HTTPServer.Port:", "value", cfg.HTTPServer.Port)
	log.Info("DB.Host:", "value", cfg.DB.Host)
	log.Info("DB.Port:", "value", cfg.DB.Port)
	log.Info("DB.DBName:", "value", cfg.DB.DBName)
	log.Info("DB.SslMode:", "value", cfg.DB.SslMode)
	log.Info("DB.DriverName:", "value", cfg.DB.DriverName)
	log.Info("DB.MigrationUrl:", "value", cfg.DB.MigrationUrl)

	return &cfg
}
