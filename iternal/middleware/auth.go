package middleware

import (
	"errors"
	"go-backend/iternal/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

func AuthUser(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		logrus.Printf("Missing authentication token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing authentication token",
		})
	}

	// Извлекаем часть токена без "Bearer "
	tokenString := strings.Replace(token, "Bearer ", "", 1)

	// Проверяем валидность токена
	_, err := validateToken(tokenString)
	if err != nil {
		logrus.WithError(err).Printf("Invalid authentication token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid authentication token",
		})
	}

	// Продолжаем выполнение следующих обработчиков
	return c.Next()
}

func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Верните ваш секретный ключ для проверки подписи токена
		return config.Config.JWT_SECRET_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid authentication token")
	}

	return token, nil
}
