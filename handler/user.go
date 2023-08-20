package handler

import "github.com/gofiber/fiber/v2"

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success",
		"message": "All users",
		"data": fiber.Map{
			"userId": "hard-coded-user-id",
		},
	})
}
