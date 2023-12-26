package middlewares

import (
	"vftalk/configs"

	"github.com/gofiber/fiber/v2"
)

func AuthJWT(c *fiber.Ctx) error {
	err := configs.TokenValid(c)
	httpMethod := string(c.Request().Header.Method())
	if err != nil {
		if string(httpMethod) == fiber.MethodGet {
			if c.Route().Path == "/login" || c.Route().Path == "/register" {
				return c.Next()
			} else {
				return c.Redirect("/login", fiber.StatusTemporaryRedirect)
			}
		} else {
			c.ClearCookie(`auth`)
			c.Status(fiber.StatusUnauthorized)
			return c.Render("401", fiber.Map{
				"Title":   "401 - Unauthorized",
				"Message": "Unauthorized",
			})
		}
	}
	return c.Next()
}

// Only use in login and register handler, don't use in other places
func IsLoggedIn(c *fiber.Ctx) error {
	err := configs.TokenValid(c)
	if err != nil {
		return c.Next()
	}
	return c.Redirect("/", fiber.StatusPermanentRedirect)
}
