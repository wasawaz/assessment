// Package repository contains functions for CRUD expense to database

package repository

import (
	"github.com/lib/pq"
	"github.com/wasawaz/assessment/entity"
	"github.com/wasawaz/assessment/pkg/postgresql"
)

type IExpenseRepository interface {
	Add(entity *entity.Expense) error
	Get(id int) (entity.Expense, error)
	Update(entity entity.Expense) error
	GetAll() ([]entity.Expense, error)
}

type expenseRepository struct {
	dbContext *postgresql.Postgres
}

// NewExpenseRepository -.
func NewExpenseRepository(dbContext *postgresql.Postgres) *expenseRepository {
	return &expenseRepository{dbContext}
}

// Add - add new expense.
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

// Get - get expense by Id.
func (e *expenseRepository) Get(id int) (entity.Expense, error) {
	expense := entity.Expense{}
	stmt, err := e.dbContext.Db.Prepare(`SELECT ID, TITLE, AMOUNT, NOTE, TAGS FROM EXPENSES WHERE ID = $1`)
	if err != nil {
		return expense, err
	}
	err = stmt.QueryRow(id).Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
	if err != nil {
		return expense, err
	}
	return expense, nil
}

// Update - update expense by Id.
func (e *expenseRepository) Update(entity entity.Expense) error {
	stmt, err := e.dbContext.Db.Prepare(`UPDATE EXPENSES SET TITLE = $1, AMOUNT = $2, NOTE = $3, TAGS = $4 WHERE ID = $5`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(entity.Title, entity.Amount, entity.Note, pq.Array(&entity.Tags), entity.Id)
	if err != nil {
		return err
	}
	return nil
}

// GetAll - get a;; expense.
func (e *expenseRepository) GetAll() ([]entity.Expense, error) {
	expenses := []entity.Expense{}
	expense := entity.Expense{}
	stmt, err := e.dbContext.Db.Prepare(`SELECT ID, TITLE, AMOUNT, NOTE, TAGS FROM EXPENSES`)
	if err != nil {
		return expenses, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return expenses, err
	}

	for rows.Next() {
		rows.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
		expenses = append(expenses, expense)
	}
	return expenses, nil
}
