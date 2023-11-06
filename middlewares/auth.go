package middlewares

import (
	"vftalk/conf"

	"github.com/gofiber/fiber/v2"
)

func AuthJWT(c *fiber.Ctx) error {
	err := conf.TokenValid(c)
	httpMethod := string(c.Request().Header.Method())
	if err != nil {
		if string(httpMethod) == fiber.MethodGet {
			if c.Route().Path == "/login" {
				return c.Next()
			} else if c.Route().Path == "/register" {
				return c.Next()
			} else {
				return c.Redirect("/login", fiber.StatusTemporaryRedirect)
			}
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized Access",
			})
		}
	}
	return c.Next()
}

// Only use in login and register handler, don't use in other places
func IsLoggedIn(c *fiber.Ctx) error {
	err := conf.TokenValid(c)
	if err != nil {
		return c.Next()
	}
	return c.Redirect("/", fiber.StatusPermanentRedirect)
}
