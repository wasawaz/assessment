package testmock

import "github.com/wasawaz/assessment/entity"

type MockExpenseRepository struct {
	Expense  entity.Expense
	Expenses []entity.Expense
}

func (r *MockExpenseRepository) Add(entity *entity.Expense) error {
	entity.Id = 1
	return nil
}

func (r *MockExpenseRepository) Get(id int) (entity.Expense, error) {
	return entity.Expense{Id: 1}, nil
}

func (r *MockExpenseRepository) Update(entity entity.Expense) error {
	return nil
}

func (r *MockExpenseRepository) GetAll() ([]entity.Expense, error) {
	return r.Expenses, nil
}
