package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils/v2"
)

// ref: https://github.com/gofiber/fiber/issues/41
func setupApp() *fiber.App {
	app := fiber.New()

	return app
}

const bearerToken = "Bearer o8SMWsPx.9wJ7nZ86.XUyQV3yu"

func TestGetUser(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/api/users/1", nil)
	req.Header.Set("Authorization", bearerToken)

	res, _ := app.Test(req)

	if res.StatusCode == 200 {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body)) // => Hello, World!
	}
}

// ref: https://github.com/gofiber/fiber/issues/756
func Test_Handler(t *testing.T) {
	app := New()

	app.Get("/test", func(c *Ctx) {
		c.SendStatus(400)
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/test", nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 400, resp.StatusCode, "Status code")
}
