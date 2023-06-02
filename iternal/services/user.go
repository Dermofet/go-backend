package services

import (
	facade "go-backend/iternal/database/facade"
	"go-backend/iternal/database/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserService struct {
	Facade facade.DBFacade
}

func (s *UserService) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	err := c.BodyParser(user)
	if err != nil {
		log.Print("UserService.CreateUser wrong input")
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	err = s.Facade.CreateUser(user)
	if err != nil {
		log.Print("UserService.CreateUser error: ", err)
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})
}

func (s *UserService) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	user, err := s.Facade.GetUserByEmail(email)
	if err != nil {
		log.Print("UserService.GetUserByEmail error: ", err)
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Could not find user", "data": err})
	}

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User has found", "data": user})
}

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
