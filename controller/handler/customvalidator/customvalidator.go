// Package customvalidator contains functions for validate struct of http request payload
// This use with echo framework
package customvalidator

import (
	"github.com/go-playground/validator"
)

// expenseValidator is used to assign echo validator which implement from go-playground/validator
type expenseValidator struct {
	validator *validator.Validate
}
// NewCustomValidator is used to create new expenseValidator validator
func NewCustomValidator(validator *validator.Validate) *expenseValidator {
	return &expenseValidator{validator}
}
// Validate is implement of go-playground/validator.validate function
func (cv *expenseValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return err
	}
	return nil
}
