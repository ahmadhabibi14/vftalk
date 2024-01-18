package apis

import (
	"vftalk/configs"
	"vftalk/services"

	"github.com/gofiber/fiber/v2"
)

func (a *ApisHandler) UpdateProfile(c *fiber.Ctx) error {
	ctx := c.Context()
	response := HTTPResponse{}
	in := services.InUser_UpdateProfile{}

	if err := c.BodyParser(&in); err != nil {
		response = HTTPResponse{
			Code:   fiber.StatusBadRequest,
			Status: STATUS_BADREQUEST,
			Errors: ERROR_INVALIDPAYLOAD,
			Data:   "",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	userId, err := configs.GetUserIdFromJWTfunc(c)
	if err != nil {
		response = HTTPResponse{
			Code:   fiber.StatusUnauthorized,
			Status: STATUS_UNAUTHORIZED,
			Errors: ERROR_UNAUTHORIZED,
			Data:   "",
		}
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	user := services.NewUser(a.Db, a.Log)

	in.UserID = userId.(string)
	updateProfile := user.UpdateProfile(ctx, in)
	if updateProfile != nil {
		response = HTTPResponse{
			Code:   fiber.StatusInternalServerError,
			Status: STATUS_INTERNALSERVERERROR,
			Errors: updateProfile.Error(),
			Data:   "",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response = HTTPResponse{
		Code:   fiber.StatusOK,
		Status: STATUS_OK,
		Errors: "",
		Data:   "Profile updated !!",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
