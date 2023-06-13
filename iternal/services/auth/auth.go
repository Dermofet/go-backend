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
// @Success 201 {object} schemas.Response
// @Failure 400 {object} schemas.Response
// @Failure 409 {object} schemas.Response
// @Router /auth/signup [post]
func (a *AuthService) SignUp(c *fiber.Ctx) error {
	user := new(models.User)

	err := c.BodyParser(user)
	if err != nil {
		logrus.WithError(err).Printf("Cannot parse body")
		resp := schemas.Response{Status: "error", Message: "Invalid body", Data: nil}
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	errors := schemas.ValidateStruct(*user)
	if errors != nil {
		logrus.Printf("Validation error")
		resp := schemas.Response{Status: "error", Message: "Validation error", Data: errors[0]}
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	err = a.Facade.CreateUser(user)
	if err != nil {
		logrus.WithError(err).Printf("Cannot create user")
		resp := schemas.Response{Status: "error", Message: "Could not create user", Data: nil}
		return c.Status(fiber.StatusConflict).JSON(resp)
	}

	userInfo := schemas.UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	logrus.Print("Create user ", userInfo.ID)
	resp := schemas.Response{Status: "success", Message: "User has created", Data: userInfo}
	return c.Status(fiber.StatusCreated).JSON(resp)
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
// @Failure 409 {object} schemas.Response
// @Router /auth/login [post]
func (a *AuthService) SignIn(c *fiber.Ctx) error {
	user := new(schemas.UserSignIn)

	err := c.BodyParser(user)
	if err != nil {
		logrus.WithError(err).Printf("Cannot parse body")
		resp := schemas.Response{Status: "error", Message: "Invalid body", Data: nil}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
	}

	fmt.Print(user)

	errors := schemas.ValidateStruct(*user)
	if errors != nil {
		logrus.Printf("Validation error")
		resp := schemas.Response{Status: "error", Message: "Validation error", Data: errors[0]}
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	userInfo, ok := a.checkUser(user)
	if !ok {
		logrus.WithError(err).Printf("Wrong email or password")
		resp := schemas.Response{Status: "error", Message: "Wrong email or password", Data: nil}
		return c.Status(fiber.StatusConflict).JSON(resp)
	}

	token, err := Token(userInfo)
	if err != nil {
		logrus.WithError(err).Printf("Cannot sign in user")
		resp := schemas.Response{Status: "error", Message: "Could not sign in user", Data: nil}
		return c.Status(fiber.StatusConflict).JSON(resp)
	}

	logrus.Print("Sign in user ", userInfo.ID)
	resp := schemas.Response{Status: "success", Message: "Successfully sign in", Data: schemas.AccessToken{Token: token}}
	return c.Status(fiber.StatusOK).JSON(resp)
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
