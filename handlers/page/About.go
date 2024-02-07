package page

import (
	"github.com/gofiber/fiber/v2"
)

func (p *PageHandler) About(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return c.Render("about/index", fiber.Map{
		"Title":     "VFtalk | About",
		"Permalink": p.Domain + "/about",
	})
}
