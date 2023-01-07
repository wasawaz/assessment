//go:build unit
// +build unit

package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	testmoq "github.com/wasawaz/assessment/usecase/mock"
)

func TestGetExpenseUsecase(t *testing.T) {
	// Setup
	id := 1
	mockExpenseRepository := &testmoq.MockExpenseRepository{}
	getExpenseUsecase := NewGetExpenseUsecase(mockExpenseRepository)

	// Arrange
	expense, err := getExpenseUsecase.Execute(id)

	// Assertions
	if assert.NoError(t, err) {
		assert.NotEqual(t, 0, expense.Id, "expect id not equal 0")
	}
}
