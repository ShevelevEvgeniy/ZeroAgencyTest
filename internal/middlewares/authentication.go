package middlewares

import "C"
import (
	"ZeroAgencyTest/config"
	"ZeroAgencyTest/lib/logger/sl"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"strings"
)

func Authentication(log *slog.Logger, cfg config.Auth) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		s := ctx.Get("Authorization")
		tokenString := strings.TrimPrefix(s, "Bearer ")

		if err := validateToken(tokenString, cfg); err != nil {
			log.Error("Authorisation Error: ", sl.Err(err))
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Authorisation Error",
			})
		}
		return ctx.Next()
	}
}

func validateToken(tokenString string, cfg config.Auth) error {
	if tokenString == "" {
		return fmt.Errorf("token not found")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(cfg.Token), nil
	})

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	} else {
		return fmt.Errorf("unable to successfully authenticate your request")
	}
}
