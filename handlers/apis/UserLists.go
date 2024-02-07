package apis

import (
	"net/http"
	"vftalk/services"

	"github.com/gofiber/fiber/v2"
)

func (a *ApisHandler) UserLists(c *fiber.Ctx) error {
	ctx := c.Context()
	response := HTTPResponse{}

	user := services.NewUser(a.Db, a.Log)
	userLists, err := user.UserLists(ctx)
	if err != nil {
		response = HTTPResponse{
			Code:   fiber.StatusInternalServerError,
			Status: http.StatusText(fiber.StatusInternalServerError),
			Errors: "Something went wrong",
			Data:   "",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response = HTTPResponse{
		Code:   fiber.StatusOK,
		Status: http.StatusText(fiber.StatusOK),
		Errors: "",
		Data: struct {
			Msg   string                  `json:"message"`
			Users []services.OutUserLists `json:"users"`
		}{
			Msg:   "User lists fetched !",
			Users: userLists,
		},
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
