package apis

import (
	"net/http"
	"vftalk/configs"
	"vftalk/services"

	"github.com/gofiber/fiber/v2"
)

func (a *ApisHandler) AuthRegister(c *fiber.Ctx) error {
	ctx := c.Context()
	response := HTTPResponse{}
	in := services.InUser_Create{}

	if err := c.BodyParser(&in); err != nil {
		response = HTTPResponse{
			Code:   fiber.StatusBadRequest,
			Status: http.StatusText(fiber.StatusBadRequest),
			Errors: ERROR_INVALIDPAYLOAD,
			Data:   "",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	user := services.NewUser(a.Db, a.Log)
	token, err := user.CreateUser(ctx, in)
	if err != nil {
		response = HTTPResponse{
			Code:   fiber.StatusBadRequest,
			Status: http.StatusText(fiber.StatusBadRequest),
			Errors: err.Error(),
			Data:   "",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	configs.SetJWTasCookie(c, token)
	response = HTTPResponse{
		Code:   fiber.StatusOK,
		Status: http.StatusText(fiber.StatusOK),
		Errors: "",
		Data: struct {
			Msg      string `json:"message"`
			Username string `json:"username"`
		}{
			Msg:      "Register successful !",
			Username: in.Username,
		},
	}

	err = a.Mailer.SendUserRegisterEmail(in.Email)
	if err != nil {
		a.Log.Error().
			Str(`ERROR`, err.Error()).
			Msg(`Cannot send email when user` + in.Username + `register`)
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
