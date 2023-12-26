package page

import (
	"database/sql"

	"github.com/rs/zerolog"
)

type PageHandler struct {
	Log *zerolog.Logger
	Db  *sql.DB
}
