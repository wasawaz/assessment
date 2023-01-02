package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpense(t *testing.T) {
	// Setup
	expenseJson := `{"title":"strawberry smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}`
	expectedExpenseJson := `{"title":"strawberry smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}` + "\n"
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(expenseJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewCreateExpenseHandler()

	// Assertions
	if assert.NoError(t, h.CreateExpense(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expectedExpenseJson, rec.Body.String())
	}
}
