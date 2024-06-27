package news_handler

import (
	"ZeroAgencyTest/internal/DTOs"
	newsConverter "ZeroAgencyTest/internal/converter/news_converter"
	servModel "ZeroAgencyTest/internal/service/news_service/models"
	"ZeroAgencyTest/lib/logger/sl"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"strconv"
)

type Handler struct {
	log       *slog.Logger
	service   Service
	validator *validator.Validate
}

type NewsHandler interface {
	GetNews(c *fiber.Ctx) error
	UpdateNews(c *fiber.Ctx) error
}

type Service interface {
	GetNews(ctx context.Context) ([]servModel.News, error)
	UpdateNews(ctx context.Context, news servModel.News) error
}

func NewNewsHandler(log *slog.Logger, service Service, validator *validator.Validate) *Handler {
	return &Handler{
		log:       log,
		service:   service,
		validator: validator,
	}
}

func (h *Handler) GetNews(c *fiber.Ctx) error {
	h.log.Info("Starting GetNews handler", slog.String("op", "GetNewsHandler"))

	news, err := h.service.GetNews(c.Context())
	if err != nil {
		h.log.Error("Error while executing GetNews service", sl.Err(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error while executing request",
		})
	}

	h.log.Info("Successfully retrieved news", slog.String("op", "GetNewsHandler"), slog.Int("news_count", len(news)))

	return c.JSON(NewNewsResponse(news))
}

func (h *Handler) UpdateNews(c *fiber.Ctx) error {
	h.log.Info("Starting UpdateNews handler", slog.String("op", "UpdateNewsHandler"))

	id := c.Params("id")

	var request DTOs.News
	if err := c.BodyParser(&request); err != nil {
		h.log.Error("Error parsing request body", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if request.Id == 0 {
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			h.log.Error("Invalid ID parameter", sl.Err(err))
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID parameter",
			})
		}

		request.Id = intId
	}

	if err := h.validator.Struct(request); err != nil {
		h.log.Error("Validation error", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation error",
		})
	}

	err := h.service.UpdateNews(c.Context(), newsConverter.DTOsToServModel(request))
	if err != nil {
		h.log.Error("Error while executing UpdateNews service", sl.Err(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update news",
		})
	}

	h.log.Info("Successfully updated news", slog.String("op", "UpdateNewsHandler"), slog.Int("news_id", int(request.Id)))

	return c.JSON(fiber.Map{
		"Success": true,
	})
}
