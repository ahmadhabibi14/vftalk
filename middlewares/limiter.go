package middlewares

import (
	"time"
	"vftalk/handlers/apis"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var Limiter = limiter.Config{
	Max:        300,
	Expiration: 30 * time.Second,
	KeyGenerator: func(c *fiber.Ctx) string {
		return c.IP()
	},
	LimitReached: func(c *fiber.Ctx) error {
		response := apis.HTTPResponse{
			Code:   fiber.StatusTooManyRequests,
			Status: apis.STATUS_TOOMANYREQUEST,
			Errors: "You have exceeded your rate limit. Please try again after the specified time.",
			Data:   "",
		}
		return c.Status(fiber.StatusTooManyRequests).JSON(response)
	},
}
