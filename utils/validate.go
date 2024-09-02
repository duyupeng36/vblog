package utils

import "github.com/go-playground/validator/v10"

func Validate(v any) error {
	validator := validator.New()

	return validator.Struct(v)
}
