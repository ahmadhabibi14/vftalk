package page

import "github.com/gofiber/fiber/v2"

func (p *PageHandler) Login(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, "text/html; charset=utf-8")
	return c.Render("login", fiber.Map{
		"Title": "Login",
		"Desc":  "Hi, Welcome back 👋",
	})
}
