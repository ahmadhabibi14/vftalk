package page

import (
	"github.com/gofiber/fiber/v2"
)

func (p *PageHandler) Register(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return c.Render("register/index", fiber.Map{
		"Title": "VFtalk | Register",
	})
}
