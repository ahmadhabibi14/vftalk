package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	ErrInvalidPassword = `Please use a stronger password with at least 8 characters, including numbers and symbols.`
	ErrInvalidEmail    = `Invalid email address. Please check and try again.`
)

func ValidateStruct(myStruct interface{}) error {
	validate := validator.New()
	err := validate.Struct(myStruct)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, err := range validationErrors {
			if strings.Contains(err.Field(), `assword`) {
				errorMessages[i] = ErrInvalidPassword
			} else if strings.Contains(err.Field(), `mail`) {
				errorMessages[i] = ErrInvalidEmail
			} else {
				errorMessages[i] = fmt.Sprintf("Error when validation %s", err.Field())
			}
		}
		msg := errorMessages[0]
		return errors.New(msg)
	}
	return nil
}
