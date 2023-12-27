package apis

import (
	"database/sql"
	"vftalk/configs"
	"vftalk/models/mailer"

	"github.com/rs/zerolog"
)

type ApisHandler struct {
	Mailer mailer.Mailer
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

const (
	STATUS_OK                  = "OK"
	STATUS_BADREQUEST          = "BAD REQUEST"
	STATUS_INTERNALSERVERERROR = "INTERNAL SERVER ERROR"
	ERROR_INVALIDPAYLOAD       = "The payload or input provided is invalid. Please check your request and try again."
	ERROR_GENERATETOKEN        = "Error generate session token"
)
