package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/heshan-g/go-api/router"
)

func main() {
	app := fiber.New()

	app.Get("/health-check", func (c *fiber.Ctx) error {
		return c.SendString("Server is healthy")
	})

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
