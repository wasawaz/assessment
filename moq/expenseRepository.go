package testmoq

import "github.com/wasawaz/assessment/entity"

type MockExpenseRepository struct{}

func (r *MockExpenseRepository) Add(entity *entity.Expense) error {
	entity.Id = 1
	return nil
}

func (r *MockExpenseRepository) Get(id int) (entity.Expense, error) {
	return entity.Expense{Id: 1}, nil
}
