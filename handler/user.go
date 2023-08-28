package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heshan-g/go-api/config"
	"github.com/heshan-g/go-api/model"
)

func GetUsers(c *fiber.Ctx) error {
	var users []model.User

	result := config.DB.Find(&users)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(users)
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
