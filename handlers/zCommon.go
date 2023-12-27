package handlers

import (
	"database/sql"
	"vftalk/configs"
	"vftalk/models/mailer"

	"github.com/rs/zerolog"
)

type ApisHandler struct {
	Mailer *mailer.Mailer
	Log    *zerolog.Logger
	Db     *sql.DB
	OAuth  configs.OAuthConf
}

type HTTPResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Errors string      `json:"errors"`
	Data   interface{} `json:"data"`
}
