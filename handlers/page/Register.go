package page

import (
	"github.com/gofiber/fiber/v2"
)

func (p *PageHandler) Register(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, "text/html; charset=utf-8")
	return c.Render("register/index", fiber.Map{
		"Title": "VFtalk | Register",
	})
}
