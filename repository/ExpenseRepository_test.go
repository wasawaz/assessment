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

	mockExpense := &entity.Expense{
		Title:  "mock title",
		Amount: 200.0,
		Note:   "mock note",
		Tags:   []string{"mock"},
	}
	db, mock, err := sqlmock.New()
	assert.Nil(t, err, "an error was not expected when opening a stub database connection")
	defer db.Close()
	mockDB := &postgresql.PostgresqlDB{Db: db}
	expenseRepository := NewExpenseRepository(mockDB)

	mockRows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectPrepare("INSERT INTO EXPENSES").ExpectQuery().WithArgs("mock title", 200.0, "mock note", pq.Array([]string{"mock"})).WillReturnRows(mockRows)
	err = expenseRepository.Add(mockExpense)
	assert.Nil(t, err, "an error was not expected when add new expense")
	assert.NotEqual(t, 0, mockExpense.Id, "expect expenseId equal 0")

}
