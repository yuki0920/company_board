package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuki0920/company_board/models"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "Yuki",
		LastName:  "Watanabe",
		Email:     "test@example.com",
	}

	return c.JSON(user)
}
