package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuki0920/company_board/database"
	"github.com/yuki0920/company_board/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	// リクエストボディをdataにパースする
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		return c.JSON(fiber.Map{
			"message": "password do not match",
		})
	}

	// パスワードのハッシュ化
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}
