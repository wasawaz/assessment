//go:build unit
// +build unit

package handler

import (
	expense_middleware "github.com/wasawaz/assessment/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/wasawaz/assessment/controller/handler/customvalidator"
	"github.com/wasawaz/assessment/entity"
)

type mockGetAllExpenseUsecase struct {
	expenses []entity.Expense
	err      error
}

func (u *mockGetAllExpenseUsecase) Execute() ([]entity.Expense, error) {
	return u.expenses, u.err
}

func TestGetAllExpense(t *testing.T) {
	t.Run("should return http status 200", func(t *testing.T) {
		e := echo.New()
		e.Validator = customvalidator.NewCustomValidator(validator.New())
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization,"November 10, 2009")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/expenses")
		h := NewGetAllExpenseHandler(&mockGetAllExpenseUsecase{expenses: []entity.Expense{{Id: 1}}})
		wrappedHandler := expense_middleware.AuthMiddleware(h.GetAllExpense)
		// Assertions
		if assert.NoError(t, wrappedHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("should return http status 401", func(t *testing.T) {
		e := echo.New()
		e.Validator = customvalidator.NewCustomValidator(validator.New())
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization,"November 10, 2009wrong_token")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/expenses")
		h := NewGetAllExpenseHandler(&mockGetAllExpenseUsecase{expenses: []entity.Expense{{Id: 1}}})
		wrappedHandler := expense_middleware.AuthMiddleware(h.GetAllExpense)
		// Assertions
		if assert.NoError(t, wrappedHandler(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})
}
