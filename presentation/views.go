package presentation

import (
	"fmt"
	"vftalk/conf"
	"vftalk/middlewares"

	"github.com/gofiber/fiber/v2"
)

func WebViews(app *fiber.App) {
	app.Get("/", middlewares.AuthJWT, func(c *fiber.Ctx) error {
		u, _ := conf.GetUsernameFromJWT(c)
		username := fmt.Sprintf("%v", u)
		return c.Render("index", fiber.Map{
			"Title":    "VFtalk",
			"Username": username,
		})
	})
	app.Get("/login", middlewares.IsLoggedIn, func(c *fiber.Ctx) error {
		return c.Render("login/index", fiber.Map{
			"Title": "Login",
			"Desc":  "Welcome, please use your username",
		})
	})
}
