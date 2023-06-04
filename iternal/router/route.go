package router

import (
	"fmt"
	_ "go-backend/docs"
	"go-backend/iternal/config"
	"go-backend/iternal/database/dao"
	database "go-backend/iternal/database/facade"
	"go-backend/iternal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/user")
	// routes
	app.Get("/", services.HealthCheck)

	// swaggerConfig := "./docs/swagger.json"
	app.Get("/docs/*", swagger.New(swagger.Config{
		URL: fmt.Sprintf("http://%s:%d/swagger", config.Config.BACKEND_HOST, config.Config.BACKEND_PORT),
	}))

	app.Get("/swagger", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

	userService := services.UserService{Facade: database.DBFacade{UserDao: &dao.UserDAO{}}}
	v1.Post("/", userService.CreateUser)
	v1.Get("/{email}", userService.GetUserByEmail)
	v1.Get("/{id}", userService.GetUserByID)
	v1.Put("/{id}", userService.UpdateUser)
	v1.Delete("/{id}", userService.DeleteUser)
}
