package main

import (
	"ZeroAgencyTest/config"
	application "ZeroAgencyTest/internal/app"
	"ZeroAgencyTest/lib/logger/sl"
	"os"
)

// @title News API
// @version 1.0
// @description This is a sample server for managing news.
// @host localhost:8080
// @BasePath /

func main() {
	log := sl.SetupLogger(config.LoadProjectMode())

	cfg := config.MustLoad(log)

	if err := application.NewApp(log, cfg).Run(); err != nil {
		log.Error("Failed to start app: ", sl.Err(err))
		os.Exit(1)
	}
}
