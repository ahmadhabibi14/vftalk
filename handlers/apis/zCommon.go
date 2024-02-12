package apis

import (
	"database/sql"
	"errors"
	"net/http"
	"time"
	"vftalk/configs"
	"vftalk/models/mailer"
	"vftalk/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type ApisHandler struct {
	Mailer mailer.Mailer
	Log    *zerolog.Logger
	Db     *sql.DB
	OAuth  configs.OAuthConf
}

const (
	ERROR_INVALIDPAYLOAD = "The payload or input provided is invalid. Please check your request and try again."
	ERROR_GENERATETOKEN  = "Error generate session token"
	ERROR_UNAUTHORIZED   = "You are unauthorized to do this operation"
)

const (
	CHAT_SENDER_SYSTEM = "system"

	CHAT_TYPE_TEXT  = "text"
	CHAT_TYPE_INFO  = "info"
	CHAT_TYPE_ERROR = "error"
)

type (
	ChatIn struct {
		Type    string `json:"type" validate:"required"`
		Content string `json:"content" validate:"required,max=200"`
	}
	ChatOut struct {
		Sender    string    `json:"sender"`
		Type      string    `json:"type"`
		Content   string    `json:"content"`
		Timestamp time.Time `json:"datetime"`
	}
)

type HTTPResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Errors string      `json:"errors"`
	Data   interface{} `json:"data"`
}

func NewHTTPResponse(code int, errors string, data any) HTTPResponse {
	return HTTPResponse{
		Code:   code,
		Status: http.StatusText(code),
		Errors: errors,
		Data:   data,
	}
}

func ReadJSON[T any](c *fiber.Ctx, b []byte) (T, error) {
	var body T
	err := c.BodyParser(&body)
	if err != nil {
		return body, errors.New(ERROR_INVALIDPAYLOAD)
	}

	errvalid := utils.ValidateStruct(body)
	if errvalid != nil {
		return body, errvalid
	}

	return body, nil
}
