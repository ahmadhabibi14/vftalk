package handlers

import "github.com/gofiber/fiber/v2"

type (
	updateProfileIn struct {
		Username string `json:"username"`
		FullName string `json:"full_name"`
		Email    string `json:"email" validate:"email"`
		Location string `json:"location"`
		Website  string `json:"website"`
	}
	updateProfileOut struct {
		Ok       bool   `json:"ok"`
		Username string `json:"username"`
		Message  string `json:"message"`
	}
	updateProfileError struct {
		Ok       bool   `json:"ok"`
		ErrorMsg string `json:"error"`
	}
)

func UpdateProfile(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ok":      true,
		"message": "profile updated!",
	})
}
