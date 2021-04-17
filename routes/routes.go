package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuki0920/company_board/controllers"
	"github.com/yuki0920/company_board/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	// beforeフィルター的役割
	app.Use(middleware.IsAuthenticated)

	app.Get("api/user", controllers.User)
	app.Post("api/logout", controllers.Logout)
	app.Get("api/users", controllers.AllUsers)
	app.Post("api/users", controllers.CreateUser)
	app.Get("api/users/:id", controllers.GetUser)
}
