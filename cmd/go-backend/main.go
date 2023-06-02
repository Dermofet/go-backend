package main

import (
	"fmt"
	"go-backend/iternal/config"
	"go-backend/iternal/database"
	"go-backend/iternal/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(fmt.Sprint(config.Config.BACKEND_HOST, ":", config.Config.BACKEND_PORT))
}
