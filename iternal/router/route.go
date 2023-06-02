package router

import (
	"go-backend/iternal/database/dao"
	database "go-backend/iternal/database/facade"
	"go-backend/iternal/services"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/user")
	// routes
	userService := services.UserService{Facade: database.DBFacade{UserDao: &dao.UserDAO{}}}
	v1.Post("/", userService.CreateUser)
	v1.Get("/", userService.GetUserByEmail)
	v1.Get("/:id", userService.GetUserByID)
	v1.Put("/:id", userService.UpdateUser)
	v1.Delete("/:id", userService.DeleteUser)
}
