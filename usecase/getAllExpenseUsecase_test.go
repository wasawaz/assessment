package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wasawaz/assessment/entity"
	testmoq "github.com/wasawaz/assessment/moq"
)

func TestGetAllExpenseUsecase(t *testing.T) {
	mockExpenseRepository := &testmoq.MockExpenseRepository{
		Expenses: []entity.Expense{},
	}
	getAllExpenseUsecase := NewGetAllExpenseUsecase(mockExpenseRepository)
	expenses, err := getAllExpenseUsecase.Execute()
	if assert.NoError(t, err) {
		assert.NotEqual(t, 0, len(expenses), "expect expense length not equal 0")
	}
}
