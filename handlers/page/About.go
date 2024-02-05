package page

import (
	"encoding/json"
	"vftalk/services"

	"github.com/gofiber/fiber/v2"
)

func (p *PageHandler) About(c *fiber.Ctx) error {
	ctx := c.Context()

	var users string = ``
	var jsonBytes []byte

	user := services.NewUser(p.Db, p.Log)
	userLists, err := user.UserLists(ctx)
	if err == nil {
		jsonBytes, _ = json.Marshal(userLists)
		users = string(jsonBytes)
	}

	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return c.Render("about/index", fiber.Map{
		"Title":     "VFtalk | About",
		"UserLists": users,
	})
}
