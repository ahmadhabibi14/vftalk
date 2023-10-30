package conf

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetJWTasCookie(c *fiber.Ctx, tokenString string, expiration time.Time) {
	c.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    tokenString,
		Expires:  expiration,
		SameSite: "Lax",
		HTTPOnly: false,
	})
	return
}
