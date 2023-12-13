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
				"Title": "VFTalk | Chat App",
			})
		} else {
			userData, err := handlers.GetUserDataByUsername(c)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err,
				})
			}
			c.Set(fiber.HeaderContentType, "text/html; charset=utf-8")
			return c.Render("index", fiber.Map{
				"Title":    "VFtalk | Home",
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
		c.Set(fiber.HeaderContentType, "text/html; charset=utf-8")
		return c.Render("direct", fiber.Map{
			"Title":    "VFtalk | Direct Messages",
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
		c.Set(fiber.HeaderContentType, "text/html; charset=utf-8")
		return c.Render("profile", fiber.Map{
			"Title":    "VFtalk | Profile",
			"UserData": userData,
			"JoinAt":   utils.FormatTime(userData.JoinAt),
		}, "layouts/main")
	})

	app.Get("/explore", middlewares.AuthJWT, func(c *fiber.Ctx) error {
		userData, err := handlers.GetUserDataByUsername(c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}
		c.Set(fiber.HeaderContentType, "text/html; charset=utf-8")
		return c.Render("explore", fiber.Map{
			"Title":    "VFtalk | Explore",
			"UserData": userData,
			"JoinAt":   utils.FormatTime(userData.JoinAt),
		}, "layouts/main")
	})

	app.Get("/term-of-service", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, "text/html; charset=utf-8")
		return c.Render("termOfService", fiber.Map{
			"Title": "VFtalk | Term of Service",
		})
	})

	app.Get("/privacy-policy", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, "text/html; charset=utf-8")
		return c.Render("privacyPolicy", fiber.Map{
			"Title": "VFtalk | Privacy Policy",
		})
	})

	app.Get("/login", middlewares.IsLoggedIn, func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, "text/html; charset=utf-8")
		return c.Render("login", fiber.Map{
			"Title": "Login",
			"Desc":  "Hi, Welcome back 👋",
		})
	})

	app.Get("/register", middlewares.IsLoggedIn, func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, "text/html; charset=utf-8")
		return c.Render("register", fiber.Map{
			"Title": "Register",
			"Desc":  "Welcome, please create your account",
		})
	})

	app.Get("/oauth/google", middlewares.IsLoggedIn, func(c *fiber.Ctx) error {
		stateString := utils.GenerateRandomID(40)
		url := handlers.GoogleOauthConfig.AuthCodeURL(stateString)
		return c.Redirect(url, fiber.StatusTemporaryRedirect)
	})

	app.Get("/PRIVACY_POLICY", func(c *fiber.Ctx) error {
		content, err := utils.ReadFile("./PRIVACY_POLICY.md")
		if err != nil {
			return fiber.NewError(fiber.StatusServiceUnavailable)
		}
		c.Set(fiber.HeaderContentType, "text/markdown; charset=utf-8")
		return c.Send(content)
	})

	app.Get("/TERM_OF_SERVICE", func(c *fiber.Ctx) error {
		content, err := utils.ReadFile("./TERM_OF_SERVICE.md")
		if err != nil {
			return fiber.NewError(fiber.StatusServiceUnavailable)
		}
		c.Set(fiber.HeaderContentType, "text/markdown; charset=utf-8")
		return c.Send(content)
	})
}
