package app

import (
	"ZeroAgencyTest/config"
	"ZeroAgencyTest/internal/db_connection"
	newsHandler "ZeroAgencyTest/internal/http_server/api/v1/handlers/news_handler"
	"ZeroAgencyTest/internal/http_server/login"
	"ZeroAgencyTest/internal/repository/news_repository"
	newsService "ZeroAgencyTest/internal/service/news_service"
	"ZeroAgencyTest/lib/logger/sl"
	"github.com/go-playground/validator/v10"
	"gopkg.in/reform.v1"
	"log/slog"
	"os"
)

type DiContainer struct {
	log            *slog.Logger
	cfg            *config.Config
	db             *reform.DB
	auth           login.LoginHandler
	newsRepository newsService.Repository
	newsService    newsHandler.Service
	newsHandler    newsHandler.NewsHandler
	validator      *validator.Validate
}

func NewDiContainer(log *slog.Logger, cfg *config.Config) *DiContainer {
	return &DiContainer{
		log: log,
		cfg: cfg,
	}
}

func (d *DiContainer) Config() *config.Config {
	return d.cfg
}

func (d *DiContainer) Auth() login.LoginHandler {
	if d.auth == nil {
		d.auth = login.NewLoginHandler(d.log, d.Config().Auth)
	}

	return d.auth
}

func (d *DiContainer) NewsHandler() newsHandler.NewsHandler {
	if d.newsHandler == nil {
		d.newsHandler = newsHandler.NewNewsHandler(d.log, d.NewsService(), d.Validator())
	}

	return d.newsHandler
}

func (d *DiContainer) NewsService() newsHandler.Service {
	if d.newsService == nil {
		d.newsService = newsService.NewNewsService(d.NewsRepository())
	}

	return d.newsService
}

func (d *DiContainer) NewsRepository() newsService.Repository {
	if d.newsRepository == nil {
		d.newsRepository = news_repository.NewNewsRepository(d.DbConn())
	}

	return d.newsRepository
}

func (d *DiContainer) DbConn() *reform.DB {
	if d.db == nil {
		var err error

		d.db, err = db_connection.Connect(d.Config().DB, d.log)
		if err != nil {
			d.log.Error("Failed to initialize db connection: ", sl.Err(err))
			os.Exit(1)
		}
	}

	return d.db
}

func (d *DiContainer) Validator() *validator.Validate {
	if d.validator == nil {
		d.validator = validator.New()
	}

	return d.validator
}
