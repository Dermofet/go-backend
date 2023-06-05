package router

import (
	"fmt"
	_ "go-backend/docs"
	"go-backend/iternal/config"
	"go-backend/iternal/database"
	"go-backend/iternal/database/dao"
	facade "go-backend/iternal/database/facade"
	"go-backend/iternal/middleware"
	"go-backend/iternal/services"
	auth "go-backend/iternal/services/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	userGroup := app.Group("/user")
	authGroup := app.Group("/auth")

	userGroup.Use(middleware.AuthUser)

	app.Get("/", services.HealthCheck)

	app.Get("/docs/*", swagger.New(swagger.Config{
		URL: fmt.Sprintf("http://%s:%d/swagger", config.Config.BACKEND_HOST, config.Config.BACKEND_PORT),
	}))

	app.Get("/swagger", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

	facade := &facade.DBFacade{
		UserDao: &dao.UserDAO{DB: database.DB},
	}

	userService := services.UserService{Facade: facade}
	userGroup.Get("/email/:email", userService.GetUserByEmail)
	userGroup.Get("/id/:id", userService.GetUserByID)
	userGroup.Put("/id/:id", userService.UpdateUser)
	userGroup.Delete("/id/:id", userService.DeleteUser)

	authService := auth.AuthService{Facade: facade}
	authGroup.Post("/signup", authService.SignUp)
	authGroup.Post("/login", authService.SignIn)
}
