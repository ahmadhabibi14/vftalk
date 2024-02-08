package apis

import (
	"net/http"
	"vftalk/configs"
	"vftalk/services"

	"github.com/gofiber/fiber/v2"
)

func (a *ApisHandler) UpdateProfile(c *fiber.Ctx) error {
	ctx := c.Context()
	response := HTTPResponse{}

	in, err := ReadJSON[services.InUser_UpdateProfile](c, c.Body())
	if err != nil {
		response = NewHTTPResponse(fiber.StatusBadRequest, err.Error(), "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	userId, err := configs.GetUserIdFromJWTfunc(c)
	if err != nil {
		response = HTTPResponse{
			Code:   fiber.StatusUnauthorized,
			Status: http.StatusText(fiber.StatusUnauthorized),
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
			Code:   fiber.StatusBadRequest,
			Status: http.StatusText(fiber.StatusBadRequest),
			Errors: updateProfile.Error(),
			Data:   "",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response = HTTPResponse{
		Code:   fiber.StatusOK,
		Status: http.StatusText(fiber.StatusOK),
		Errors: "",
		Data:   "Profile updated !!",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
