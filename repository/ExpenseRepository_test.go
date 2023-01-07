//go:build unit
// +build unit

package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/wasawaz/assessment/entity"
	"github.com/wasawaz/assessment/pkg/postgresql"
)

// a successful case
func TestShouldAddExpense(t *testing.T) {
	// Setup
	mockExpense := &entity.Expense{
		Title:  "mock title",
		Amount: 200.0,
		Note:   "mock note",
		Tags:   []string{"mock"},
	}
	db, mock, err := sqlmock.New()
	assert.Nil(t, err, "an error was not expected when opening a stub database connection")
	defer db.Close()
	mockDB := &postgresql.Postgres{Db: db}
	expenseRepository := NewExpenseRepository(mockDB)

	mockRows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectPrepare("INSERT INTO EXPENSES").ExpectQuery().WithArgs("mock title", 200.0, "mock note", pq.Array([]string{"mock"})).WillReturnRows(mockRows)

	// Arrange
	err = expenseRepository.Add(mockExpense)

	// Assertions
	assert.Nil(t, err, "an error was not expected when add new expense")
	assert.NotEqual(t, 0, mockExpense.Id, "expect expenseId equal 0")

}

func TestShouldGetExpense(t *testing.T) {
	// Setup
	expenseId := 1
	db, mock, err := sqlmock.New()
	assert.Nil(t, err, "an error was not expected when opening a stub database connection")
	defer db.Close()
	mockDB := &postgresql.Postgres{Db: db}
	expenseRepository := NewExpenseRepository(mockDB)
	mockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).AddRow(1, "mock title", 200.0, "mock note", pq.Array([]string{"mock"}))
	mock.ExpectPrepare("SELECT (.+) EXPENSES").ExpectQuery().WithArgs(expenseId).WillReturnRows(mockRows)

	// Arrange
	queryExpense, err := expenseRepository.Get(expenseId)

	// Assertions
	assert.Nil(t, err, "an error was not expected when add new expense")
	assert.Equal(t, 1, queryExpense.Id, "expect expenseId not equal 1")
}

func TestShouldUpdateExpense(t *testing.T) {
	// Setup
	mockExpense := entity.Expense{
		Id:     1,
		Title:  "mock title",
		Amount: 200.0,
		Note:   "mock note",
		Tags:   []string{"mock"},
	}
	db, mock, err := sqlmock.New()
	assert.Nil(t, err, "an error was not expected when opening a stub database connection")
	defer db.Close()
	mockDB := &postgresql.Postgres{Db: db}
	expenseRepository := NewExpenseRepository(mockDB)
	mock.ExpectPrepare("UPDATE EXPENSES").ExpectExec().WithArgs(mockExpense.Title, mockExpense.Amount, mockExpense.Note, pq.Array(mockExpense.Tags), mockExpense.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	// Arrange
	err = expenseRepository.Update(mockExpense)

	// Assertions
	assert.Nil(t, err, "an error was not expected when update exist expense")

}

func TestShouldGetAllExpense(t *testing.T) {
	// Setup
	db, mock, err := sqlmock.New()
	assert.Nil(t, err, "an error was not expected when opening a stub database connection")
	defer db.Close()
	mockDB := &postgresql.Postgres{Db: db}
	expenseRepository := NewExpenseRepository(mockDB)
	mockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).AddRow(1, "mock title", 200.0, "mock note", pq.Array([]string{"mock"}))
	mock.ExpectPrepare("SELECT (.+) EXPENSES").ExpectQuery().WillReturnRows(mockRows)

	// Arrange
	queryExpenses, err := expenseRepository.GetAll()

	// Assertions
	assert.Nil(t, err, "an error was not expected when add new expense")
	assert.NotEqual(t, 0, len(queryExpenses), "expect expenses length greater than 0")
}
