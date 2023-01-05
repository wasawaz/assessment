package handler

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/wasawaz/assessment/controller/handler/customvalidator"
	"github.com/wasawaz/assessment/entity"
)

type mockGetExpenseUsecase struct {
	expense entity.Expense
	err     error
}

func (u *mockGetExpenseUsecase) Execute(id int) (entity.Expense, error) {
	return u.expense, u.err
}

func TestGetExpense(t *testing.T) {
	t.Run("should return http status 200", func(t *testing.T) {
		e := echo.New()
		e.Validator = customvalidator.NewCustomValidator(validator.New())
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/expense/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		h := NewGetExpenseHandler(&mockGetExpenseUsecase{expense: entity.Expense{Id: 1}})

		// Assertions
		if assert.NoError(t, h.GetExpense(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("should return http status 400 id not number", func(t *testing.T) {
		e := echo.New()
		e.Validator = customvalidator.NewCustomValidator(validator.New())
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/expense/:id")
		c.SetParamNames("id")
		c.SetParamValues("dmdemkd")
		h := NewGetExpenseHandler(&mockGetExpenseUsecase{expense: entity.Expense{Id: 1}})

		// Assertions
		if assert.NoError(t, h.GetExpense(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})

	t.Run("should return http status 400 id not found", func(t *testing.T) {
		e := echo.New()
		e.Validator = customvalidator.NewCustomValidator(validator.New())
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/expense/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		h := NewGetExpenseHandler(&mockGetExpenseUsecase{expense: entity.Expense{}, err: sql.ErrNoRows})

		// Assertions
		if assert.NoError(t, h.GetExpense(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})
}
