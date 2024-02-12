package page

import (
	"vftalk/configs"
	"vftalk/services"
	"vftalk/utils"

	"github.com/gofiber/fiber/v2"
)

func (p *PageHandler) Rooms(c *fiber.Ctx) error {
	userId, err := configs.GetUserIdFromJWTfunc(c)
	LogoutIfError(c, err)

	in := services.InUser_FindById{
		UserID: userId.(string),
	}
	user := services.NewUser(p.Db, p.Log)
	userOut, err := user.FindById(c.UserContext(), in)
	LogoutIfError(c, err)

	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return c.Render("rooms/index", fiber.Map{
		"Title":  "VFtalk | Rooms",
		"User":   userOut,
		"JoinAt": utils.FormatTime(userOut.JoinAt),
	})
}
