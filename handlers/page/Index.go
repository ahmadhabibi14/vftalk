package page

import (
	"vftalk/configs"
	"vftalk/services"
	"vftalk/utils"

	"github.com/gofiber/fiber/v2"
)

func (p *PageHandler) Index(c *fiber.Ctx) error {
	err := configs.TokenValid(c)
	if err != nil {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.Render("landingpage/index", fiber.Map{
			"Title": "VFTalk | Chat App",
		})
	} else {
		userId, err := configs.GetUserIdFromJWTfunc(c)
		LogoutIfError(c, err)

		in := services.InUser_FindById{
			UserID: userId.(string),
		}
		user := services.NewUser(p.Db, p.Log)
		userOut, err := user.FindById(c.UserContext(), in)
		LogoutIfError(c, err)

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.Render("index", fiber.Map{
			"Title":  "VFtalk | Home",
			"User":   userOut,
			"JoinAt": utils.FormatTime(userOut.JoinAt),
		})
	}
}
