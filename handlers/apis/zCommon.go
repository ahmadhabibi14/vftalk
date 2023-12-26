package apis

import (
	"database/sql"
	"vftalk/models/mailer"

	"github.com/rs/zerolog"
)

type ApisHandler struct {
	Mailer mailer.Mailer
	Log    *zerolog.Logger
	Db     *sql.DB
}

type HTTPResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Errors string      `json:"errors"`
	Data   interface{} `json:"data"`
}

const (
	STATUS_OK                  = "OK"
	STATUS_BADREQUEST          = "BAD REQUEST"
	STATUS_INTERNALSERVERERROR = "INTERNAL SERVER ERROR"
)
