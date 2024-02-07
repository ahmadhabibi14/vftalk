package apis

import (
	"vftalk/configs"
	"vftalk/services"

	"github.com/gofiber/fiber/v2"
)

func (a *ApisHandler) AuthRegister(c *fiber.Ctx) error {
	ctx := c.Context()
	response := HTTPResponse{}

	in, err := ReadJSON[services.InUser_Create](c, c.Body())
	if err != nil {
		response = JSONResponse(fiber.StatusBadRequest, err.Error(), "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	user := services.NewUser(a.Db, a.Log)
	token, err := user.CreateUser(ctx, in)
	if err != nil {
		response = JSONResponse(fiber.StatusBadRequest, err.Error(), "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	err = a.Mailer.SendUserRegisterEmail(in.Email)
	if err != nil {
		a.Log.Error().
			Str(`ERROR`, err.Error()).
			Msg(`Cannot send email when user` + in.Username + `register`)
	}

	configs.SetJWTasCookie(c, token)
	var data = struct {
		Msg      string `json:"message"`
		Username string `json:"username"`
	}{
		Msg:      "Register successful !",
		Username: in.Username,
	}
	response = JSONResponse(fiber.StatusOK, "", data)
	return c.Status(fiber.StatusOK).JSON(response)
}
