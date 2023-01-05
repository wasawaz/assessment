package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	testmoq "github.com/wasawaz/assessment/moq"
)

func TestGetExpenseUsecase(t *testing.T) {
	id := 1
	mockExpenseRepository := &testmoq.MockExpenseRepository{}
	getExpenseUsecase := NewGetExpenseUsecase(mockExpenseRepository)
	expense, err := getExpenseUsecase.Execute(id)
	if assert.NoError(t, err) {
		assert.NotEqual(t, 0, expense.Id, "expect id not equal 0")
	}
}
