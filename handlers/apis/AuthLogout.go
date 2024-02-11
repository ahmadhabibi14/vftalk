package apis

import (
	"vftalk/configs"

	"github.com/gofiber/fiber/v2"
)

func (a *ApisHandler) AuthLogout(c *fiber.Ctx) error {
	c.ClearCookie(configs.AUTH_COOKIE)
	return c.Redirect("/login", fiber.StatusTemporaryRedirect)
}
