package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) error {
	validate := validator.New()
	err := validate.Struct(s)
	errMsgs := make([]string, 0)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"Error when validating %s: '%v'",
				err.Field(),
				err.Value(),
			))
		}
		return errors.New(errMsgs[0])
	}
	return nil
}
