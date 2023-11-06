package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(myStruct interface{}) error {
	validate := validator.New()
	err := validate.Struct(myStruct)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, err := range validationErrors {
			errorMessages[i] = fmt.Sprintf("Error %s validation", err.Field())
		}
		msg := strings.Join(errorMessages, ", ")
		return errors.New(msg)
	}
	return nil
}
