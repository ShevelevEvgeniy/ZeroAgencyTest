package login

import (
	"ZeroAgencyTest/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"time"
)

type LoginHandler interface {
	Login(c *fiber.Ctx) error
}

type Handler struct {
	token string
	log   *slog.Logger
}

func NewLoginHandler(log *slog.Logger, cfg config.Auth) *Handler {
	return &Handler{
		log:   log,
		token: cfg.Token,
	}
}

// Login @Summary Generate JWT token
// @Description Generate JWT token with 72 hours expiration
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {string} string "JWT token"
// @Router /login [get]
func (h *Handler) Login(c *fiber.Ctx) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
	})

	ss, err := token.SignedString([]byte(h.token))
	if err != nil {
		return err
	}

	h.log.Info("JWT token generated successfully.")

	return c.SendString(ss)
}
