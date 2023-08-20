package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/heshan-g/go-api/database"
	"github.com/heshan-g/go-api/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("Server is healthy")
	})

	database.Connect()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
