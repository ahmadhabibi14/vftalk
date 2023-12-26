package apis

import (
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
			Status: STATUS_BADREQUEST,
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
			Status: STATUS_BADREQUEST,
			Errors: err.Error(),
			Data:   "",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	configs.SetJWTasCookie(c, token)
	response = HTTPResponse{
		Code:   fiber.StatusOK,
		Status: STATUS_OK,
		Errors: "",
		Data:   "Register successful !",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
