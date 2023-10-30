package middlewares

import (
	"log"
	"vftalk/conf"

	"github.com/gofiber/fiber/v2"
)

func AuthJWT(c *fiber.Ctx) error {
	err := conf.TokenValid(c)
	httpMethod := string(c.Request().Header.Method())
	log.Println("HTTP Method =", string(httpMethod))
	if err != nil {
		if string(httpMethod) == fiber.MethodGet {
			return c.Redirect("/login", fiber.StatusTemporaryRedirect)
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized Access",
			})
		}
	}
	return c.Next()
}

func IsLoggedIn(c *fiber.Ctx) error {
	err := conf.TokenValid(c)
	if err != nil {
		if c.Route().Path == "/login" {
			return c.Next()
		}
	}
	return c.Redirect("/", fiber.StatusPermanentRedirect)
}
