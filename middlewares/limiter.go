package middlewares

import (
	"fmt"
	"net/http"
	"time"
	"vftalk/handlers/apis"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var Limiter = limiter.Config{
	Max:        300,
	Expiration: 2 * time.Minute,
	KeyGenerator: func(c *fiber.Ctx) string {
		return c.IP()
	},
	LimitReached: func(c *fiber.Ctx) error {
		httpMethod := string(c.Request().Header.Method())
		if httpMethod == fiber.MethodGet {
			c.Status(fiber.StatusTooManyRequests)
			return c.Render("404", fiber.Map{
				"Title":   fmt.Sprintf("%d - %s", fiber.StatusTooManyRequests, http.StatusText(fiber.StatusTooManyRequests)),
				"Code":    fiber.StatusTooManyRequests,
				"Status":  http.StatusText(fiber.StatusTooManyRequests),
				"Message": "You have exceeded your rate limit. Please try again a few minutes later.",
			})
		}
		response := apis.HTTPResponse{
			Code:   fiber.StatusTooManyRequests,
			Status: http.StatusText(fiber.StatusTooManyRequests),
			Errors: "You have exceeded your rate limit. Please try again a few minutes later.",
			Data:   "",
		}
		return c.Status(fiber.StatusTooManyRequests).JSON(response)
	},
}
