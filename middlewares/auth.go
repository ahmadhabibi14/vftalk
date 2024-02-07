package middlewares

import (
	"net/http"
	"vftalk/configs"
	"vftalk/handlers/apis"

	"github.com/gofiber/fiber/v2"
)

func AuthJWT(c *fiber.Ctx) error {
	err := configs.TokenValid(c)
	httpMethod := string(c.Request().Header.Method())
	if err != nil {
		c.ClearCookie(configs.AUTH_COOKIE)
		if string(httpMethod) == fiber.MethodGet {
			if c.Route().Path == "/login" || c.Route().Path == "/register" {
				return c.Next()
			} else {
				return c.Redirect("/login", fiber.StatusTemporaryRedirect)
			}
		} else {
			resp := apis.HTTPResponse{
				Code:   fiber.StatusUnauthorized,
				Status: http.StatusText(fiber.StatusUnauthorized),
				Errors: "You are unauthorized to process this operation",
				Data:   "",
			}
			return c.Status(fiber.StatusUnauthorized).JSON(resp)
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
