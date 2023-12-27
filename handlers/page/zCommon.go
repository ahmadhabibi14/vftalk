package page

import (
	"database/sql"
	"vftalk/configs"

	"github.com/rs/zerolog"
)

type PageHandler struct {
	Log   *zerolog.Logger
	Db    *sql.DB
	OAuth configs.OAuthConf
}
