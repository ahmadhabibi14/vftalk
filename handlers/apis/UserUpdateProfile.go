package apis

import (
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
		response = NewHTTPResponse(fiber.StatusUnauthorized, ERROR_UNAUTHORIZED, "")
		c.ClearCookie(configs.AUTH_COOKIE)
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	in.UserID = userId.(string)
	user := services.NewUser(a.Db, a.Log)
	updateProfile := user.UpdateProfile(ctx, in)
	if updateProfile != nil {
		response = NewHTTPResponse(fiber.StatusBadRequest, updateProfile.Error(), "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response = NewHTTPResponse(fiber.StatusOK, "", "Profile updated !!")
	return c.Status(fiber.StatusOK).JSON(response)
}
