package apis

import (
	"database/sql"
	"errors"
	"net/http"
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
