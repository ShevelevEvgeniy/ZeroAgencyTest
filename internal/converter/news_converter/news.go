package news_converter

import (
	"ZeroAgencyTest/internal/DTOs"
	repoModel "ZeroAgencyTest/internal/repository/models"
	servModel "ZeroAgencyTest/internal/service/news_service/models"
)

func ServToRepoModel(servNews servModel.News) (*repoModel.News, *repoModel.NewsCategories) {
	news := repoModel.News{
		Id:      servNews.Id,
		Title:   nil,
		Content: nil,
	}

	if servNews.Title != "" {
		news.Title = &servNews.Title
	}

	if servNews.Content != "" {
		news.Content = &servNews.Content
	}

	var categories *int64

	if len(servNews.Categories) == 0 {
		return &news, nil
	} else {
		combinedValue := int64(0)
		for _, category := range servNews.Categories {
			combinedValue = combinedValue*10 + category
		}
		categories = &combinedValue
	}

	newsCategories := repoModel.NewsCategories{
		NewsId:     servNews.Id,
		CategoryId: categories,
	}

	return &news, &newsCategories
}

func RepoToServModels(repoNews []repoModel.NewsWithCategories) []servModel.News {
	servNews := make([]servModel.News, len(repoNews))

	for i, repo := range repoNews {
		serv := servModel.News{
			Id:         repo.Id,
			Title:      repo.Title,
			Content:    repo.Content,
			Categories: repo.Categories,
		}

		servNews[i] = serv
	}

	return servNews
}

func DTOsToServModel(DTO DTOs.News) servModel.News {
	return servModel.News{
		Id:         DTO.Id,
		Title:      DTO.Title,
		Content:    DTO.Content,
		Categories: DTO.Categories,
	}
}
