package news_service

import (
	repoModel "ZeroAgencyTest/internal/repository/models"
	"ZeroAgencyTest/internal/service/news_service/mocks"
	servModel "ZeroAgencyTest/internal/service/news_service/models"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetNews_Success(t *testing.T) {
	mockRepository := new(mocks.Repository)
	service := NewNewsService(mockRepository)

	expectedNews := []servModel.News{
		{
			Id:         1,
			Title:      "Test News 1",
			Content:    "Test content 1",
			Categories: []int64{1, 2, 3},
		},
		{
			Id:         2,
			Title:      "Test News 2",
			Content:    "Test content 2",
			Categories: []int64{4, 5},
		},
	}
	mockRepository.On("GetNews", mock.Anything).Return(expectedNews, nil)

	ctx := context.Background()
	news, err := service.GetNews(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, news)
	assert.Equal(t, expectedNews, news)

	mockRepository.AssertExpectations(t)
}

func TestGetNews_ErrorFromRepository(t *testing.T) {
	mockRepository := new(mocks.Repository)
	service := NewNewsService(mockRepository)

	expectedError := errors.New("repository error")
	mockRepository.On("GetNews", mock.Anything).Return(nil, expectedError)

	ctx := context.Background()
	news, err := service.GetNews(ctx)

	assert.Error(t, err)
	assert.Nil(t, news)
	assert.EqualError(t, err, "failed to retrieve news from repository: repository error")

	mockRepository.AssertExpectations(t)
}

func TestUpdateNews_Success(t *testing.T) {
	mockRepository := new(mocks.Repository)
	service := NewNewsService(mockRepository)

	updateNews := servModel.News{
		Id:         1,
		Title:      "Updated Title",
		Content:    "Updated Content",
		Categories: []int64{1, 2, 3},
	}

	mockNews, mockNewsCategories := prepareMockConverterBehavior(updateNews)

	mockRepository.On("UpdateNews", mock.Anything, mockNews, mockNewsCategories).Return(nil)

	ctx := context.Background()
	err := service.UpdateNews(ctx, updateNews)

	assert.NoError(t, err)

	mockRepository.AssertExpectations(t)
}

func TestUpdateNews_ErrorFromRepository(t *testing.T) {
	mockRepository := new(mocks.Repository)
	service := NewNewsService(mockRepository)

	updateNews := servModel.News{
		Id:         1,
		Title:      "Updated Title",
		Content:    "Updated Content",
		Categories: []int64{1, 2, 3},
	}

	mockNews, mockNewsCategories := prepareMockConverterBehavior(updateNews)

	expectedError := errors.New("repository error")
	mockRepository.On("UpdateNews", mock.Anything, mockNews, mockNewsCategories).Return(expectedError)

	ctx := context.Background()
	err := service.UpdateNews(ctx, updateNews)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to update news in repository: repository error")

	mockRepository.AssertExpectations(t)
}

func prepareMockConverterBehavior(updateNews servModel.News) (*repoModel.News, *repoModel.NewsCategories) {
	mockNews := &repoModel.News{
		Id:      updateNews.Id,
		Title:   &updateNews.Title,
		Content: &updateNews.Content,
	}

	var categories *int64

	if len(updateNews.Categories) == 0 {
		return mockNews, nil
	} else {
		combinedValue := int64(0)
		for _, category := range updateNews.Categories {
			combinedValue = combinedValue*10 + category
		}
		categories = &combinedValue
	}

	mockNewsCategories := &repoModel.NewsCategories{
		NewsId:     updateNews.Id,
		CategoryId: categories,
	}

	return mockNews, mockNewsCategories
}
