package services

import (
	facade "go-backend/iternal/database/facade"
	"go-backend/iternal/database/models"
	_ "go-backend/iternal/schemas"
	"log"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	Facade facade.DBFacadeInterface
}

// GetUserByEmail gets a user by Email.
// @Summary Get user by Email
// @Description Get a user by their Email.
// @ID get-user-by-email
// @Tags users
// @Accept json
// @Produce json
// @Param email path string true "User Email"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Router /user/email/{email} [get]
// @Security JWT
func (s *UserService) GetUserByEmail(c *fiber.Ctx) error {
	email, err := url.QueryUnescape(c.Params("email"))
	if err != nil {
		logrus.WithError(err).Print("Bad unescape query")

	}

	user, err := s.Facade.GetUserByEmail(email)
	if err != nil {
		logrus.WithError(err).Print("Could not find user")
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "error", "message": "Could not find user", "data": err})
	}

	if user.ID == uuid.Nil {
		logrus.WithError(err).Print("User not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "ok", "message": "User not found", "data": nil})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User has found", "data": user})
}

// GetUserByID gets a user by ID.
// @Summary Get user by ID
// @Description Get a user by their ID.
// @ID get-user-by-id
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /user/id/{id} [get]
// @Security JWT
func (s *UserService) GetUserByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		log.Print("UserService.GetUserByID error: Failed to parse UUID: ", err)
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Invalid ID", "data": err})
	}

	user, err := s.Facade.GeUsertByID(id)
	if err != nil {
		log.Print("UserService.GetUserByID error: ", err)
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Could not find user", "data": err})
	}

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User has found", "data": user})
}

// UpdateUser update user.
// @Summary Update user
// @Description Update user.
// @ID update-user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Router /user/id/{id} [put]
// @Security JWT
func (s *UserService) UpdateUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		log.Print("UserService.UpdateUser error: Failed to parse UUID: ", err)
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "ID is not valid", "data": err})
	}

	user := new(models.User)

	err = c.BodyParser(user)
	if err != nil {
		log.Print("UserService.UpdateUser wrong input")
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "User's fields are invalid", "data": err})
	}

	user, err = s.Facade.UpdateUser(id, user)
	if err != nil {
		log.Print("UserService.UpdateUser error: ", err)
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Could not update user", "data": err})
	}

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User has updated", "data": user})
}

// DeleteUser delete user.
// @Summary Delete user
// @Description Delete user.
// @ID delete-user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Router /user/id/{id} [delete]
// @Security JWT
func (s *UserService) DeleteUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		log.Print("UserService.DeleteUser error: Failed to parse UUID: ", err)
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "ID is not valid", "data": err})
	}

	err = s.Facade.DeleteUser(id)
	if err != nil {
		log.Print("UserService.DeleteUser error: ", err)
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Could not delete user", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User has deleted", "data": nil})
}
