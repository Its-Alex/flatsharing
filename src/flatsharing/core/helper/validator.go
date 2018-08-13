package helper

import (
	validator "gopkg.in/go-playground/validator.v9"
)

var (
	// Validate is used to valide struct
	Validate *validator.Validate
)

func init() {
	Validate = validator.New()
}
