package repository

import (
	"github.com/lib/pq"
	"github.com/wasawaz/assessment/entity"
	"github.com/wasawaz/assessment/pkg/postgresql"
)

type IExpenseRepository interface {
	Add(entity *entity.Expense) error
}

type expenseRepository struct {
	dbContext *postgresql.PostgresqlDB
}

func NewExpenseRepository(dbContext *postgresql.PostgresqlDB) *expenseRepository {
	return &expenseRepository{dbContext}
}

func (e *expenseRepository) Add(entity *entity.Expense) error {
	stmt, err := e.dbContext.Db.Prepare(`INSERT INTO EXPENSES(TITLE, AMOUNT, NOTE, TAGS) VALUES($1, $2, $3, $4) RETURNING id`)
	if err != nil {
		return err
	}
	var id int
	err = stmt.QueryRow(entity.Title, entity.Amount, entity.Note, pq.Array(&entity.Tags)).Scan(&id)
	if err != nil {
		return err
	}
	entity.Id = id
	return nil
}
