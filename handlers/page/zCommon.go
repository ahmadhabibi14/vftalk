package page

import (
	"database/sql"
	"vftalk/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type PageHandler struct {
	Log   *zerolog.Logger
	Db    *sql.DB
	OAuth configs.OAuthConf
}

func LogoutIfError(c *fiber.Ctx, err error) error {
	if err != nil {
		c.ClearCookie(`auth`)
		return c.Redirect("/", fiber.StatusTemporaryRedirect)
	}

	return nil
}
