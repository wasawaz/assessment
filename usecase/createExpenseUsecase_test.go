package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wasawaz/assessment/entity"
)

type mockExpenseRepository struct{}

func (r *mockExpenseRepository) Add(entity *entity.Expense) error {
	entity.Id = 1
	return nil
}

func TestCreateExpenseUsecase(t *testing.T) {

	newExpense := &entity.Expense{}
	mockExpenseRepository := &mockExpenseRepository{}
	createExpenseUsecase := NewCreateExpenseUsecase(mockExpenseRepository)

	if assert.NoError(t, createExpenseUsecase.Execute(newExpense)) {
		assert.NotEqual(t, 0, newExpense.Id, "expect id not equal 0")
	}
}
