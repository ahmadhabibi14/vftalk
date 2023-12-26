package apis

import (
	"vftalk/services"

	"github.com/gofiber/fiber/v2"
)

func (a *ApisHandler) Debug(c *fiber.Ctx) error {
	id := "A4qJdgJ6H3aoLir7WiqxMDJXqRM"
	response := HTTPResponse{}
	user := services.NewUser(a.Db, a.Log)
	userOut, err := user.FindById(id)
	if err != nil {
		response = HTTPResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_BADREQUEST,
			Errors: err.Error(),
			Data:   "",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if mailErr := a.Mailer.SendUserRegisterEmail(userOut.Email); mailErr != nil {
		a.Log.Error().Str("Error: ", mailErr.Error()).Msg("Canno send email")
	}

	response = HTTPResponse{
		Code:   fiber.StatusOK,
		Status: STATUS_OK,
		Errors: "",
		Data:   userOut,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
