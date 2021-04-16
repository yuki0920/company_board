package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuki0920/company_board/util"
)

func IsAuthenticated(c *fiber.Ctx) error {
	// クッキー取得
	cookie := c.Cookies("jwt")

	if _, err := util.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	return c.Next()
}
