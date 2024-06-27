package app

import (
	"ZeroAgencyTest/config"
	"ZeroAgencyTest/internal/middlewares"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"log/slog"
)

type App struct {
	log *slog.Logger
	app *fiber.App
	di  *DiContainer
}

func NewApp(log *slog.Logger, cfg *config.Config) *App {
	return &App{
		log: log,
		di:  NewDiContainer(log, cfg),
	}
}

func (a App) Run() error {
	a.app = fiber.New(
		fiber.Config{
			//Prefork: a.di.Config().HTTPServer.Prefork,
		})

	a.app.Get("/login", func(c *fiber.Ctx) error {
		return a.di.Auth().Login(c)
	})

	API := a.app.Group("/api/v1")
	API.Use(middlewares.Authentication(a.log, a.di.Config().Auth))
	NewsAPIV1 := API.Group("news")

	NewsAPIV1.Get("/list", a.di.NewsHandler().GetNews)
	NewsAPIV1.Put("/edit/:id", a.di.NewsHandler().UpdateNews)

	err := a.app.Listen(fmt.Sprintf(":%s", a.di.Config().HTTPServer.Port))
	if err != nil {
		return errors.Wrap(err, "failed to start the server")
	}
	return nil
}
