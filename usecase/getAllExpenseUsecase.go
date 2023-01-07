// Package usecase contains business logic for get all expenses
package usecase

import (
	"github.com/wasawaz/assessment/entity"
	"github.com/wasawaz/assessment/repository"
)

type IGetAllExpenseUsecase interface {
	Execute() ([]entity.Expense, error)
}

type getAllExpenseUsecase struct {
	expenseRepository repository.IExpenseRepository
}

// NewGetAllExpenseUsecase -.
func NewGetAllExpenseUsecase(expenseRepository repository.IExpenseRepository) *getAllExpenseUsecase {
	return &getAllExpenseUsecase{expenseRepository}
}

// Execute - execute business logic for get all expense.
func (u *getAllExpenseUsecase) Execute() ([]entity.Expense, error) {
	expense, err := u.expenseRepository.GetAll()
	if err != nil {
		return []entity.Expense{}, err
	}
	return expense, nil
}
