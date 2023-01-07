//go:build unit
// +build unit

package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wasawaz/assessment/entity"
	testmoq "github.com/wasawaz/assessment/usecase/mock"
)

func TestUpdateExpenseUsecase(t *testing.T) {
	// Setup
	newExpense := entity.Expense{}
	mockExpenseRepository := &testmoq.MockExpenseRepository{}
	updateExpenseUsecase := NewUpdateExpenseUsecase(mockExpenseRepository)

	// Arrange
	err := updateExpenseUsecase.Execute(newExpense)

	// Assertions
	assert.NoError(t, err)
}
