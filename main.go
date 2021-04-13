package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuki0920/company_board/database"
	"github.com/yuki0920/company_board/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
