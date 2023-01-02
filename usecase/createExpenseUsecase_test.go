package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wasawaz/assessment/entity"
	"github.com/wasawaz/assessment/repository"
)

func TestCreateExpenseUsecase(t *testing.T) {

	newExpense := &entity.Expense{}
	expenseRepository := &repository.ExpenseRepository{}
	createExpenseUsecase := NewCreateExpenseUsecase(expenseRepository)

	if assert.NoError(t, createExpenseUsecase.Execute(newExpense)) {
		assert.NotEqual(t, 0, newExpense.Id, "expect id not equal 0 but ")
	}
}
