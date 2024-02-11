package configs

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

var AUTH_COOKIE = `auth`

func SetJWTasCookie(c *fiber.Ctx, tokenString string) {
	expiration := time.Now().AddDate(0, 2, 0)
	c.Cookie(&fiber.Cookie{
		Name:     AUTH_COOKIE,
		Value:    tokenString,
		Expires:  expiration,
		SameSite: "Lax",
		Secure:   os.Getenv("WEB_ENV") == "prod",
		HTTPOnly: true,
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
