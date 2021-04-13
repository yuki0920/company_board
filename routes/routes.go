package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuki0920/company_board/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
}
