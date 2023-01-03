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

func NewCreateExpenseUsecase(expenseRepository repository.IExpenseRepository) *createExpenseUsecase {
	return &createExpenseUsecase{expenseRepository}
}

func (u *createExpenseUsecase) Execute(expense *entity.Expense) error {
	err := u.expenseRepository.Add(expense)
	if err != nil {
		return err
	}
	return nil
}
