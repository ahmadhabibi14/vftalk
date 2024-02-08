package middlewares

import (
	"strings"
	"vftalk/handlers/apis"

	"github.com/gofiber/fiber/v2"
)

func ContentJSON(c *fiber.Ctx) error {
	ctype := c.Get(fiber.HeaderContentType)
	if ctype != fiber.MIMEApplicationJSON {
		resp := apis.NewHTTPResponse(
			fiber.StatusBadRequest,
			"Invalid Content-Type",
			"Try to use "+fiber.MIMEApplicationJSON+" for valid Content-Type",
		)
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}
	return c.Next()
}

func ContentMultipartForm(c *fiber.Ctx) error {
	ctype := c.Get(fiber.HeaderContentType)
	if !strings.Contains(ctype, fiber.MIMEMultipartForm) {
		resp := apis.NewHTTPResponse(
			fiber.StatusBadRequest,
			"Invalid Content-Type",
			"Try to use "+fiber.MIMEMultipartForm+" for valid Content-Type",
		)
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}
	return c.Next()
}
