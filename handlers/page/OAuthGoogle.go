package page

import (
	"vftalk/utils"

	"github.com/gofiber/fiber/v2"
)

func (p *PageHandler) OAuthGoogle(c *fiber.Ctx) error {
	stateString := utils.GenerateRandomID(40)
	url := p.OAuth.Google.AuthCodeURL(stateString)
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}
