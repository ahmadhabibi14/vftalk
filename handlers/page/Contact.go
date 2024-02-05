package page

import (
	"github.com/gofiber/fiber/v2"
)

func (p *PageHandler) Contact(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return c.Render("contact/index", fiber.Map{
		"Title": "VFtalk | Contact",
	})
}
