package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wasawaz/assessment/entity"
	testmoq "github.com/wasawaz/assessment/moq"
)

func TestUpdateExpenseUsecase(t *testing.T) {

	newExpense := entity.Expense{}
	mockExpenseRepository := &testmoq.MockExpenseRepository{}
	updateExpenseUsecase := NewUpdateExpenseUsecase(mockExpenseRepository)
	assert.NoError(t, updateExpenseUsecase.Execute(newExpense))
}
