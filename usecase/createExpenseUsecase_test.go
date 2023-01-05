package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wasawaz/assessment/entity"
	testmoq "github.com/wasawaz/assessment/moq"
)

func TestCreateExpenseUsecase(t *testing.T) {

	newExpense := &entity.Expense{}
	mockExpenseRepository := &testmoq.MockExpenseRepository{}
	createExpenseUsecase := NewCreateExpenseUsecase(mockExpenseRepository)

	if assert.NoError(t, createExpenseUsecase.Execute(newExpense)) {
		assert.NotEqual(t, 0, newExpense.Id, "expect id not equal 0")
	}
}
