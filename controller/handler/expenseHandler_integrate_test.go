//go:build integration
// +build integration

package handler

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/wasawaz/assessment/controller/handler/customvalidator"
	expense_middleware "github.com/wasawaz/assessment/middleware"
	"github.com/wasawaz/assessment/pkg/postgresql"
	"github.com/wasawaz/assessment/repository"
	"github.com/wasawaz/assessment/usecase"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestExpense(t *testing.T) {
	t.Run("should created return http status 201", func(t *testing.T) {
		// Setup
		os.Setenv("DATABASE_URL", "postgres://expenses:P@ssw0rd@expense_mock_db/expenses?sslmode=disable")
		expenseJson := `{"title":"strawberry smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}`
		expectedExpenseJson := `{"id":1,"title":"strawberry smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}` + "\n"
		e := echo.New()
		e.Validator = customvalidator.NewCustomValidator(validator.New())
		req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(expenseJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "November 10, 2009")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		pg, err := postgresql.New(os.Getenv("DATABASE_URL"))
		assert.NoError(t, err)
		defer pg.Close()
		expenseRepository := repository.NewExpenseRepository(pg)
		createExpenseUsecase := usecase.NewCreateExpenseUsecase(expenseRepository)
		h := NewCreateExpenseHandler(createExpenseUsecase)
		wrappedHandler := expense_middleware.AuthMiddleware(h.CreateExpense)
		// Assertions
		if assert.NoError(t, wrappedHandler(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, expectedExpenseJson, rec.Body.String())
		}
	})

	t.Run("should get return http status 200", func(t *testing.T) {
		// Setup
		os.Setenv("DATABASE_URL", "postgres://expenses:P@ssw0rd@expense_mock_db/expenses?sslmode=disable")
		e := echo.New()
		e.Validator = customvalidator.NewCustomValidator(validator.New())
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "November 10, 2009")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/expenses/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		pg, err := postgresql.New(os.Getenv("DATABASE_URL"))
		assert.NoError(t, err)
		defer pg.Close()
		expenseRepository := repository.NewExpenseRepository(pg)
		getExpenseUsecase := usecase.NewGetExpenseUsecase(expenseRepository)
		h := NewGetExpenseHandler(getExpenseUsecase)
		wrappedHandler := expense_middleware.AuthMiddleware(h.GetExpense)
		// Assertions
		if assert.NoError(t, wrappedHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("should updated return http status 202", func(t *testing.T) {
		// Setup
		os.Setenv("DATABASE_URL", "postgres://expenses:P@ssw0rd@expense_mock_db/expenses?sslmode=disable")
		expenseJson := `{"title":"strawberry smoothie","amount":80,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}`
		expectedExpenseJson := `{"id":1,"title":"strawberry smoothie","amount":80,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}` + "\n"
		e := echo.New()
		e.Validator = customvalidator.NewCustomValidator(validator.New())
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(expenseJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "November 10, 2009")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/expenses/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		pg, err := postgresql.New(os.Getenv("DATABASE_URL"))
		assert.NoError(t, err)
		defer pg.Close()
		expenseRepository := repository.NewExpenseRepository(pg)
		updateExpenseUsecase := usecase.NewUpdateExpenseUsecase(expenseRepository)
		h := NewUpdateExpenseHandler(updateExpenseUsecase)
		wrappedHandler := expense_middleware.AuthMiddleware(h.UpdateExpense)

		// Assertions
		if assert.NoError(t, wrappedHandler(c)) {
			assert.Equal(t, http.StatusAccepted, rec.Code)
			assert.Equal(t, expectedExpenseJson, rec.Body.String())
		}
	})

	t.Run("should get all return http status 200", func(t *testing.T) {
		// Setup
		os.Setenv("DATABASE_URL", "postgres://expenses:P@ssw0rd@expense_mock_db/expenses?sslmode=disable")
		e := echo.New()
		e.Validator = customvalidator.NewCustomValidator(validator.New())
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "November 10, 2009")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/expenses")
		pg, err := postgresql.New(os.Getenv("DATABASE_URL"))
		assert.NoError(t, err)
		defer pg.Close()
		expenseRepository := repository.NewExpenseRepository(pg)
		getAllExpenseUsecase := usecase.NewGetAllExpenseUsecase(expenseRepository)
		h := NewGetAllExpenseHandler(getAllExpenseUsecase)
		wrappedHandler := expense_middleware.AuthMiddleware(h.GetAllExpense)
		// Assertions
		if assert.NoError(t, wrappedHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
