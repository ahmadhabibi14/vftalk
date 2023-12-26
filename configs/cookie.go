package configs

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

var AUTH_COOKIE = `auth`

func SetJWTasCookie(c *fiber.Ctx, tokenString string, expiration time.Time) {
	c.Cookie(&fiber.Cookie{
		Name:     AUTH_COOKIE,
		Value:    tokenString,
		Expires:  expiration,
		SameSite: "Lax",
		HTTPOnly: false,
	})
	return
}

func RemoveCookie(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:    AUTH_COOKIE,
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
	})
}
