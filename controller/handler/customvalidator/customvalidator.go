package customvalidator

import (
	"github.com/go-playground/validator"
)

type expenseValidator struct {
	validator *validator.Validate
}

func NewCustomValidator(validator *validator.Validate) *expenseValidator {
	return &expenseValidator{validator}
}

func (cv *expenseValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return err
	}
	return nil
}
