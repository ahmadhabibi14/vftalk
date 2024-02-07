package apis

import (
	"vftalk/configs"
	"vftalk/services"

	"github.com/gofiber/fiber/v2"
)

func (a *ApisHandler) AuthLogin(c *fiber.Ctx) error {
	ctx := c.Context()
	response := HTTPResponse{}

	in, err := ReadJSON[services.InUser_AuthLogin](c, c.Body())
	if err != nil {
		response = JSONResponse(fiber.StatusBadRequest, err.Error(), "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	user := services.NewUser(a.Db, a.Log)
	token, username, err := user.AuthLogin(ctx, in)
	if err != nil {
		response = JSONResponse(fiber.StatusBadRequest, err.Error(), "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	configs.SetJWTasCookie(c, token)

	var data = struct {
		Msg      string `json:"message"`
		Username string `json:"username"`
	}{
		Msg:      "Login successful !",
		Username: username,
	}
	response = JSONResponse(fiber.StatusOK, "", data)
	return c.Status(fiber.StatusOK).JSON(response)
}
