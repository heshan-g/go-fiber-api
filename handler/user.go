package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heshan-g/go-api/config"
	"github.com/heshan-g/go-api/model"
)

func GetUsers(c *fiber.Ctx) error {
	type GetUser struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
		DeletedAt string `json:"deletedAt"`
	}

	var users []GetUser

	q := "SELECT id, name, email, created_at, updated_at, deleted_at FROM users"

	result := config.DB.Raw(q).Scan(&users)
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

	return c.Status(201).JSON(fiber.Map{"id": user.ID})
}
