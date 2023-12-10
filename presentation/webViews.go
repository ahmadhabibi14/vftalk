package presentation

import (
	"vftalk/conf"
	"vftalk/handlers"
	"vftalk/middlewares"
	"vftalk/utils"

	"github.com/gofiber/fiber/v2"
)

func WebViews(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		err := conf.TokenValid(c)
		if err != nil {
			return c.Render("landingpage", fiber.Map{
				"Title": "VFTalk",
			})
		} else {
			userData, err := handlers.GetUserDataByUsername(c)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err,
				})
			}
			c.Set("Content-Type", "text/html; charset=utf-8")
			return c.Render("index", fiber.Map{
				"Title":    "VFtalk",
				"Username": userData.Username,
				"UserData": userData,
				"JoinAt":   utils.FormatTime(userData.JoinAt),
			}, "layouts/main")
		}
	})

	app.Get("/direct", middlewares.AuthJWT, func(c *fiber.Ctx) error {
		userData, err := handlers.GetUserDataByUsername(c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}
		c.Set("Content-Type", "text/html; charset=utf-8")
		return c.Render("direct", fiber.Map{
			"Title":    "Direct Messages",
			"Username": userData.Username,
			"UserData": userData,
			"JoinAt":   utils.FormatTime(userData.JoinAt),
		}, "layouts/main")
	})

	app.Get("/profile", middlewares.AuthJWT, func(c *fiber.Ctx) error {
		userData, err := handlers.GetUserDataByUsername(c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}
		c.Set("Content-Type", "text/html; charset=utf-8")
		return c.Render("profile", fiber.Map{
			"Title":    "Profile",
			"UserData": userData,
			"JoinAt":   utils.FormatTime(userData.JoinAt),
		}, "layouts/main")
	})

	app.Get("/login", middlewares.IsLoggedIn, func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html; charset=utf-8")
		return c.Render("login", fiber.Map{
			"Title": "Login",
			"Desc":  "Hi, Welcome back 👋",
		})
	})

	app.Get("/register", middlewares.IsLoggedIn, func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html; charset=utf-8")
		return c.Render("register", fiber.Map{
			"Title": "Register",
			"Desc":  "Welcome, please create your account",
		})
	})
}
