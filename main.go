package main

import (
	"example/server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(helmet.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${status} - ${method} ${path}\n",
	}))

	apiV1 := app.Group("/api/v1")
	routes.SignInRoutes(apiV1)

	app.Listen(":3000")
}
