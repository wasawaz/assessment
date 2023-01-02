package usecase

import (
	"github.com/wasawaz/assessment/entity"
	"github.com/wasawaz/assessment/repository"
)

type ICreateExpenseUsecase interface {
	Execute(expense *entity.Expense) error
}

type CreateExpenseUsecase struct {
	expenseRepository repository.IExpenseRepository
}

func NewCreateExpenseUsecase(expenseRepository repository.IExpenseRepository) *CreateExpenseUsecase {
	return &CreateExpenseUsecase{expenseRepository}
}

func (u *CreateExpenseUsecase) Execute(expense *entity.Expense) error {
	expense.Id = 1
	return nil
}
