//go:build unit
// +build unit

package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wasawaz/assessment/entity"
	testmoq "github.com/wasawaz/assessment/usecase/mock"
)

func TestGetAllExpenseUsecase(t *testing.T) {
	// Setup
	mockExpenseRepository := &testmoq.MockExpenseRepository{
		Expenses: []entity.Expense{
			{
				Title:  "",
				Amount: 79,
				Note:   "night market promotion discount 10 bath",
				Tags:   []string{"food", "beverage"},
			},
		},
	}
	getAllExpenseUsecase := NewGetAllExpenseUsecase(mockExpenseRepository)

	// Arrange
	expenses, err := getAllExpenseUsecase.Execute()

	// Assertions
	if assert.NoError(t, err) {
		assert.NotEqual(t, 0, len(expenses), "expect expense length not equal 0")
	}
}
