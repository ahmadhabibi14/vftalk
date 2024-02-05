package page

import (
	"encoding/json"
	"fmt"
	"vftalk/services"

	"github.com/gofiber/fiber/v2"
)

func (p *PageHandler) About(c *fiber.Ctx) error {
	ctx := c.Context()

	var users string
	var isError bool
	var jsonBytes []byte

	user := services.NewUser(p.Db, p.Log)
	userLists, err := user.UserLists(ctx)
	if err != nil {
		users = fmt.Sprintf("%v", userLists)
		isError = true
	}
	if !isError {
		jsonBytes, _ = json.Marshal(userLists)
		users = string(jsonBytes)
	}

	return c.Render("about/index", fiber.Map{
		"Title":     "About | VFtalk",
		"UserLists": users,
	})
}
