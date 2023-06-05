package services

import (
	"go-backend/iternal/config"
	"go-backend/iternal/schemas"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

func Token(user *schemas.UserInfo) (string, error) {
	payload := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(config.Config.JWT_SECRET_KEY)
	if err != nil {
		logrus.WithError(err).Error("JWT token signing")
		return "", err
	}

	return token, nil
}
