// Package usecase contains business logic for create new expense

package usecase

import (
	"github.com/wasawaz/assessment/entity"
	"github.com/wasawaz/assessment/repository"
)

type ICreateExpenseUsecase interface {
	Execute(expense *entity.Expense) error
}

type createExpenseUsecase struct {
	expenseRepository repository.IExpenseRepository
}

// NewCreateExpenseUsecase -.
func NewCreateExpenseUsecase(expenseRepository repository.IExpenseRepository) *createExpenseUsecase {
	return &createExpenseUsecase{expenseRepository}
}
// Execute - execute business logic for create new expense.
func (u *createExpenseUsecase) Execute(expense *entity.Expense) error {
	err := u.expenseRepository.Add(expense)
	if err != nil {
		return err
	}
	return nil
}
