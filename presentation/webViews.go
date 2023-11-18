package presentation

import (
	"fmt"
	"vftalk/conf"
	"vftalk/handlers"
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
		}, "layouts/main")
	})

	app.Get("/profile", middlewares.AuthJWT, func(c *fiber.Ctx) error {
		userData, err := handlers.GetUserDataByUsername(c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.Render("profile", fiber.Map{
			"Title":    "Profile",
			"UserData": userData,
		}, "layouts/main")
	})

	app.Get("/login", middlewares.IsLoggedIn, func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{
			"Title": "Login",
			"Desc":  "Hi, Welcome back 👋",
		})
	})

	app.Get("/register", middlewares.IsLoggedIn, func(c *fiber.Ctx) error {
		return c.Render("register", fiber.Map{
			"Title": "Register",
			"Desc":  "Welcome, please create your account",
		})
	})
}
