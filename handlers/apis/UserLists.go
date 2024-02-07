package apis

import (
	"vftalk/configs"
	"vftalk/services"

	"github.com/gofiber/fiber/v2"
)

func (a *ApisHandler) UserLists(c *fiber.Ctx) error {
	ctx := c.Context()
	response := HTTPResponse{}

	_, err := configs.GetUserIdFromJWTfunc(c)
	if err != nil {
		response = JSONResponse(fiber.StatusUnauthorized, ERROR_UNAUTHORIZED, "")
		c.ClearCookie(configs.AUTH_COOKIE)
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	user := services.NewUser(a.Db, a.Log)
	userLists, err := user.UserLists(ctx)
	if err != nil {
		response = JSONResponse(fiber.StatusInternalServerError, "Something went wrong", "")
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	var data = struct {
		Msg   string                  `json:"message"`
		Users []services.OutUserLists `json:"users"`
	}{
		Msg:   "User lists fetched !",
		Users: userLists,
	}
	response = JSONResponse(fiber.StatusOK, "", data)
	return c.Status(fiber.StatusOK).JSON(response)
}
