// Package usecase contains business logic for update expense by id
package usecase

import (
	"github.com/wasawaz/assessment/entity"
	"github.com/wasawaz/assessment/repository"
)

type IUpdateExpenseUsecase interface {
	Execute(expense entity.Expense) error
}

type updateExpenseUsecase struct {
	expenseRepository repository.IExpenseRepository
}

// NewUpdateExpenseUsecase -
func NewUpdateExpenseUsecase(expenseRepository repository.IExpenseRepository) *updateExpenseUsecase {
	return &updateExpenseUsecase{expenseRepository}
}

// Execute - execute business logic for update expense by id.
func (u *updateExpenseUsecase) Execute(expense entity.Expense) error {
	_, err := u.expenseRepository.Get(expense.Id)
	if err != nil {
		return err
	}
	err = u.expenseRepository.Update(expense)
	if err != nil {
		return err
	}
	return nil
}
