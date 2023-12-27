package apis

import (
	"vftalk/services"

	"github.com/gofiber/fiber/v2"
)

func (a *ApisHandler) Debug(c *fiber.Ctx) error {
	ctx := c.Context()
	response := HTTPResponse{}
	in := services.InUser_FindById{
		UserID: "6f935d5c-1f55-4e6c-bd24-13e6ef6fb129",
	}

	a.Log.Info().Str("Host: ", c.Hostname()).Msg("Hostname for this request")

	user := services.NewUser(a.Db, a.Log)
	userOut, err := user.FindById(ctx, in)
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
		a.Log.Error().Str("Error", mailErr.Error()).Msg("Cannot send email")
	}

	response = HTTPResponse{
		Code:   fiber.StatusOK,
		Status: STATUS_OK,
		Errors: "",
		Data:   userOut,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
