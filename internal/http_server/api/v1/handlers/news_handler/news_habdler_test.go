package news_handler

import (
	"ZeroAgencyTest/internal/DTOs"

	"ZeroAgencyTest/internal/http_server/api/v1/handlers/news_handler/mocks"
	"ZeroAgencyTest/internal/service/news_service/models"
	"bytes"
	"encoding/json"
	valid "github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"log/slog"
)

func newTestHandler() (*Handler, *mocks.Service) {
	mockService := new(mocks.Service)
	mockLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	validator := valid.New()
	handler := NewNewsHandler(mockLogger, mockService, validator)
	return handler, mockService
}

func setupApp(handler *Handler) *fiber.App {
	app := fiber.New()
	app.Get("/list", handler.GetNews)
	app.Put("/edit/:id", handler.UpdateNews)
	return app
}

func TestGetNews_Successful(t *testing.T) {
	handler, service := newTestHandler()
	a := setupApp(handler)

	mockNews := []models.News{
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

	service.On("GetNews", mock.Anything).Return(mockNews, nil)

	req := httptest.NewRequest(http.MethodGet, "/list", nil)
	resp, _ := a.Test(req)

	require.Equal(t, http.StatusOK, resp.StatusCode)
	service.AssertExpectations(t)

	var response NewsResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(t, err)
	require.Equal(t, len(mockNews), len(response.News))
	for i, news := range mockNews {
		require.Equal(t, news.Title, response.News[i].Title)
		require.Equal(t, news.Content, response.News[i].Content)
		require.Equal(t, news.Categories, response.News[i].Categories)
	}
}

func TestGetNews_ServiceError(t *testing.T) {
	handler, service := newTestHandler()
	a := setupApp(handler)

	service.On("GetNews", mock.Anything).Return(nil, errors.New("service error"))

	req := httptest.NewRequest(http.MethodGet, "/list", nil)
	resp, _ := a.Test(req)

	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	service.AssertExpectations(t)
}

func TestUpdateNews_Successful(t *testing.T) {
	handler, service := newTestHandler()
	a := setupApp(handler)

	updateRequest := DTOs.News{
		Id:         1,
		Title:      "Updated Title",
		Content:    "Updated Content",
		Categories: []int64{1, 2, 3},
	}
	body, _ := json.Marshal(updateRequest)

	service.On("UpdateNews", mock.Anything, mock.AnythingOfType("models.News")).Return(nil)

	req := httptest.NewRequest(http.MethodPut, "/edit/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := a.Test(req)

	require.Equal(t, http.StatusOK, resp.StatusCode)
	service.AssertExpectations(t)

	var response map[string]bool
	err := json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(t, err)
	require.True(t, response["Success"])
}

func TestUpdateNews_InvalidRequestBody(t *testing.T) {
	handler, _ := newTestHandler()
	a := setupApp(handler)

	req := httptest.NewRequest(http.MethodPut, "/edit/1", bytes.NewReader([]byte("{invalid json")))
	req.Header.Set("Content-Type", "application/json")
	resp, err := a.Test(req)
	require.NoError(t, err)

	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUpdateNews_ServiceError(t *testing.T) {
	handler, service := newTestHandler()
	a := setupApp(handler)

	updateRequest := DTOs.News{
		Id:         1,
		Title:      "Updated Title",
		Content:    "Updated Content",
		Categories: []int64{1, 2, 3},
	}
	body, _ := json.Marshal(updateRequest)

	service.On("UpdateNews", mock.Anything, mock.AnythingOfType("models.News")).Return(errors.New("service error"))

	req := httptest.NewRequest(http.MethodPut, "/edit/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := a.Test(req)
	require.NoError(t, err)

	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	service.AssertExpectations(t)
}
