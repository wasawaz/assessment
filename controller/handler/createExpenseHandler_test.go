package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/wasawaz/assessment/controller/handler/customvalidator"
	"github.com/wasawaz/assessment/entity"
)

type mockCreateExpenseUsecase struct{}

func (u *mockCreateExpenseUsecase) Execute(expense *entity.Expense) error {
	return nil
}

func TestCreateExpense(t *testing.T) {

	t.Run("should return http status 201", func(t *testing.T) {
		// Setup
		expenseJson := `{"title":"strawberry smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}`
		expectedExpenseJson := `{"id":0,"title":"strawberry smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}` + "\n"
		e := echo.New()
		e.Validator = customvalidator.NewCustomValidator(validator.New())
		req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(expenseJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := NewCreateExpenseHandler(&mockCreateExpenseUsecase{})

		// Assertions
		if assert.NoError(t, h.CreateExpense(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, expectedExpenseJson, rec.Body.String())
		}
	})

	t.Run("should return http status 400", func(t *testing.T) {
		// Setup
		expenseJson := `{"title":"","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}`
		e := echo.New()
		e.Validator = customvalidator.NewCustomValidator(validator.New())
		req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(expenseJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := NewCreateExpenseHandler(&mockCreateExpenseUsecase{})

		// Assertions
		if assert.NoError(t, h.CreateExpense(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}
