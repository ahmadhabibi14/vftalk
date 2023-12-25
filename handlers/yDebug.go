package handlers

import "github.com/gofiber/fiber/v2"

func (h Handler) Debug(c *fiber.Ctx) error {
	h.Log.Info().Msg(`Ignore it, its just an info`)
	h.Mailer.SendUserRegisterEmail(`ahmadhabibi7159@gmail.com`)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		`success`: `Its for debug`,
	})
}
