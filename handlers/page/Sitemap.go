package page

import (
	"github.com/gofiber/fiber/v2"
)

func (p *PageHandler) Sitemap(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationXMLCharsetUTF8)
	c.Status(fiber.StatusOK)
	return c.SendFile("views/pages/dist/sitemap-index.xml")
}

func (p *PageHandler) Sitemap0(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationXMLCharsetUTF8)
	c.Status(fiber.StatusOK)
	return c.SendFile("views/pages/dist/sitemap-0.xml")
}
