package presentation

import (
	"github.com/gofiber/fiber/v2"
)

func WebViews(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "VFtalk",
		})
	})
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login/index", fiber.Map{
			"Title": "Login",
			"Desc":  "Welcome, please use your username",
		})
	})
}
