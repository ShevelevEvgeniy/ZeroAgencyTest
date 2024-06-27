package news_service

import (
	newsConverter "ZeroAgencyTest/internal/converter/news_converter"
	repoModel "ZeroAgencyTest/internal/repository/models"
	servModel "ZeroAgencyTest/internal/service/news_service/models"
	"context"
	"github.com/pkg/errors"
)

type Service struct {
	repository Repository ``
}

type Repository interface {
	GetNews(ctx context.Context) ([]servModel.News, error)
	UpdateNews(ctx context.Context, news *repoModel.News, newsCategories *repoModel.NewsCategories) error
}

func NewNewsService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetNews(ctx context.Context) ([]servModel.News, error) {
	news, err := s.repository.GetNews(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve news from repository")
	}
	return news, err
}

func (s *Service) UpdateNews(ctx context.Context, servModel servModel.News) error {
	news, newsCategories := newsConverter.ServToRepoModel(servModel)

	err := s.repository.UpdateNews(ctx, news, newsCategories)
	if err != nil {
		return errors.Wrap(err, "failed to update news in repository")
	}

	return nil
}
