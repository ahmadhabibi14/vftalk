package apis

import (
	"vftalk/services"

	"github.com/gofiber/fiber/v2"
)

func (a *ApisHandler) Debug(c *fiber.Ctx) error {
	ctx := c.Context()
	response := HTTPResponse{}

	UserID := "6f935d5c-1f55-4e6c-bd24-13e6ef6fb129"
	user := services.NewUser(a.Db, a.Log)
	if !user.Debug(ctx, UserID) {
		response = NewHTTPResponse(fiber.StatusBadRequest, "User ID not found", "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response = NewHTTPResponse(fiber.StatusOK, "", "User ID found")
	return c.Status(fiber.StatusOK).JSON(response)
}
