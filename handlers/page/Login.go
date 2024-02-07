package page

import "github.com/gofiber/fiber/v2"

func (p *PageHandler) Login(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return c.Render("login/index", fiber.Map{
		"Title":     "VFtalk | Login",
		"Permalink": "https://vftalk.my.id/login",
	})
}
