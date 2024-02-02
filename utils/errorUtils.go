package utils

import "github.com/go-playground/validator/v10"

func ValidatorErrors(err error) map[string]interface{} {
	errors := make(map[string]interface{})
	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = err.Value()
	}
	return errors
}
