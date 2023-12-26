package apis

import (
	"time"
	"vftalk/configs"
	"vftalk/services"

	"github.com/gofiber/fiber/v2"
)

func (a *ApisHandler) AuthRegister(c *fiber.Ctx) error {
	var in services.InUser_Create
	response := HTTPResponse{}
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
	token, err := user.CreateUser(in)
	if err != nil {
		response = HTTPResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_BADREQUEST,
			Errors: err.Error(),
			Data:   "",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	configs.SetJWTasCookie(c, token, time.Now().AddDate(0, 2, 0))
	response = HTTPResponse{
		Code:   fiber.StatusOK,
		Status: STATUS_OK,
		Errors: "",
		Data:   "Register successful !",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
