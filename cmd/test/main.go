package main

import (
	"ZeroAgencyTest/config"
	application "ZeroAgencyTest/internal/app"
	"ZeroAgencyTest/lib/logger/sl"
	"os"
)

func main() {
	log := sl.SetupLogger(config.LoadProjectMode())

	cfg := config.MustLoad(log)

	if err := application.NewApp(log, cfg).Run(); err != nil {
		log.Error("Failed to start app: ", sl.Err(err))
		os.Exit(1)
	}
}
