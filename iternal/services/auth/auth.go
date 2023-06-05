package services

import (
	"fmt"
	facade "go-backend/iternal/database/facade"
	"go-backend/iternal/database/models"
	"go-backend/iternal/schemas"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	Facade facade.DBFacadeInterface
}

// SignUp sign up user
// @Summary Sign up user
// @Description Sign up user
// @ID sign-up
// @Tags auth
// @Accept json
// @Produce json
// @Param body body schemas.UserSignUp true "Request body"
// @Success 201 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Router /auth/signup [post]
func (a *AuthService) SignUp(c *fiber.Ctx) error {
	user := new(models.User)

	err := c.BodyParser(user)
	if err != nil {
		logrus.WithError(err).Printf("Cannot parse body")
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid body",
			"data":    err,
		})
	}

	err = a.Facade.CreateUser(user)
	if err != nil {
		logrus.WithError(err).Printf("Cannot create user")
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create user",
			"data":    err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User has created",
		"data":    user,
	})
}

// SignIn sign in user
// @Summary Sign in user
// @Description Sign in user
// @ID sign-in
// @Tags auth
// @Accept json
// @Produce json
// @Param body body schemas.UserSignIn true "Request body"
// @Success 201 {object} schemas.AccessToken
// @Failure 409 {object} map[string]interface{}
// @Router /auth/login [post]
func (a *AuthService) SignIn(c *fiber.Ctx) error {
	user := new(schemas.UserSignIn)

	err := c.BodyParser(user)
	if err != nil {
		logrus.WithError(err).Printf("Cannot parse body")
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid body",
			"data":    err,
		})
	}
	fmt.Print(user)
	userInfo, ok := a.checkUser(user)

	if !ok {
		logrus.WithError(err).Printf("Wrong email or password")
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "error",
			"message": "Wrong email or password",
			"data":    err,
		})
	}

	token, err := Token(userInfo)
	if err != nil {
		logrus.WithError(err).Printf("Cannot sign in user")
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not sign in user",
			"data":    err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully sign in",
		"data":    schemas.AccessToken{Token: token},
	})
}

func (a *AuthService) checkUser(user *schemas.UserSignIn) (*schemas.UserInfo, bool) {
	userDB, err := a.Facade.GetUserByEmail(user.Email)

	if err != nil || user.Password != userDB.Password {
		return nil, false
	}

	return &schemas.UserInfo{
		ID:       userDB.ID,
		Username: userDB.Username,
		Email:    userDB.Email,
	}, true
}
