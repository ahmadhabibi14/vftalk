package page

import (
	"github.com/gofiber/fiber/v2"
)

func (p *PageHandler) RickRoll(c *fiber.Ctx) error {
	return c.Redirect("https://www.youtube.com/watch?v=dQw4w9WgXcQ", fiber.StatusPermanentRedirect)
}
