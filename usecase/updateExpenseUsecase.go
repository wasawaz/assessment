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

func NewUpdateExpenseUsecase(expenseRepository repository.IExpenseRepository) *updateExpenseUsecase {
	return &updateExpenseUsecase{expenseRepository}
}

func (u *updateExpenseUsecase) Execute(expense entity.Expense) error {
	err := u.expenseRepository.Update(expense)
	if err != nil {
		return err
	}
	return nil
}
