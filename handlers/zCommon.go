package handlers

import (
	"vftalk/models/mailer"

	"github.com/rs/zerolog"
)

type Handler struct {
	Mailer mailer.Mailer
	Log    zerolog.Logger
}

type HTTPResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Errors string      `json:"errors"`
	Data   interface{} `json:"data"`
}
