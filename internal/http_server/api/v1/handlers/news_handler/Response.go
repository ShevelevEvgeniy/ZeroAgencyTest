package news_handler

import (
	newsService "ZeroAgencyTest/internal/service/news_service/models"
)

type NewsResponse struct {
	Success bool
	News    []newsService.News
}

func NewNewsResponse(news []newsService.News) NewsResponse {
	return NewsResponse{
		Success: true,
		News:    news,
	}
}
