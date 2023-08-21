package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heshan-g/go-api/config"
	"github.com/heshan-g/go-api/model"
)

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success",
		"message": "All users",
		"data": fiber.Map{
			"userId": "hard-coded-user-id",
		},
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)
	err := c.BodyParser(user)
	if err != nil {
		return err
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return c.Status(201).JSON(fiber.Map{ "id": user.ID })
}
