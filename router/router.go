package router

import (
	"github.com/heshan-g/go-api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes (app *fiber.App) {
	api := app.Group("/api", logger.New())

	// User
	user := api.Group("/user")
	user.Get("/", handler.GetUsers)
	user.Post("/", handler.CreateUser)
}
