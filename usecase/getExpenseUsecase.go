package usecase

import (
	"github.com/wasawaz/assessment/entity"
	"github.com/wasawaz/assessment/repository"
)

type IGetExpenseUsecase interface {
	Execute(id int) (entity.Expense, error)
}

type getExpenseUsecase struct {
	expenseRepository repository.IExpenseRepository
}

func NewGetExpenseUsecase(expenseRepository repository.IExpenseRepository) *getExpenseUsecase {
	return &getExpenseUsecase{expenseRepository}
}

func (u *getExpenseUsecase) Execute(id int) (entity.Expense, error) {
	expense, err := u.expenseRepository.Get(id)
	if err != nil {
		return entity.Expense{}, err
	}
	return expense, nil
}
