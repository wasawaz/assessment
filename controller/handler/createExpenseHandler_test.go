//go:build unit
// +build unit

package handler

import (
	expense_middleware "github.com/wasawaz/assessment/middleware"
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
		req.Header.Set(echo.HeaderAuthorization, "November 10, 2009")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := NewCreateExpenseHandler(&mockCreateExpenseUsecase{})
		wrappedHandler := expense_middleware.AuthMiddleware(h.CreateExpense)

		// Arrange
		err := wrappedHandler(c)

		// Assertions
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, expectedExpenseJson, rec.Body.String())
		}
	})

	t.Run("should return http status 400", func(t *testing.T) {
		// Setup
		tests := []struct {
			Name                   string
			Input                  string
			ExpectedHttpStatusCode int
			ExpectedJson           string
		}{
			{Name: "title empty", Input: `{"title":"","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}`, ExpectedHttpStatusCode: http.StatusBadRequest, ExpectedJson: ""},
			{Name: "amount less or equal 0", Input: `{"title":"strawberry smoothie","amount":0,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}`, ExpectedHttpStatusCode: http.StatusBadRequest, ExpectedJson: ""},
		}
		for _, test := range tests {
			t.Run(test.Name, func(t *testing.T) {
				e := echo.New()
				e.Validator = customvalidator.NewCustomValidator(validator.New())
				req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(test.Input))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				req.Header.Set(echo.HeaderAuthorization, "November 10, 2009")
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				h := NewCreateExpenseHandler(&mockCreateExpenseUsecase{})
				wrappedHandler := expense_middleware.AuthMiddleware(h.CreateExpense)

				// Arrange
				err := wrappedHandler(c)

				// Assertions
				if assert.NoError(t, err) {
					assert.Equal(t, test.ExpectedHttpStatusCode, rec.Code)
					assert.Equal(t, test.ExpectedJson, rec.Body.String())
				}
			})
		}

	})

	t.Run("should return http status 401", func(t *testing.T) {
		// Setup
		expenseJson := `{"title":"strawberry smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}`
		expectedExpenseJson := ``
		e := echo.New()
		e.Validator = customvalidator.NewCustomValidator(validator.New())
		req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(expenseJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "November 10, 2009wrong_token")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := NewCreateExpenseHandler(&mockCreateExpenseUsecase{})
		wrappedHandler := expense_middleware.AuthMiddleware(h.CreateExpense)

		// Arrange
		err := wrappedHandler(c)

		// Assertions
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
			assert.Equal(t, expectedExpenseJson, rec.Body.String())
		}
	})
}
