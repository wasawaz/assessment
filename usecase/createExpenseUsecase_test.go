//go:build unit
// +build unit

package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wasawaz/assessment/entity"
	testmoq "github.com/wasawaz/assessment/usecase/mock"
)

func TestCreateExpenseUsecase(t *testing.T) {
	// Setup
	newExpense := &entity.Expense{}
	mockExpenseRepository := &testmoq.MockExpenseRepository{}
	createExpenseUsecase := NewCreateExpenseUsecase(mockExpenseRepository)

	// Arrange
	err:= createExpenseUsecase.Execute(newExpense)

	// Assertions
	if assert.NoError(t,err) {
		assert.NotEqual(t, 0, newExpense.Id, "expect id not equal 0")
	}
}
