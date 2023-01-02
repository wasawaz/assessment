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
	expenseJson := "{}"
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(expenseJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := newCreateExpenseHandler()

	// Assertions
	if assert.NoError(t, h.createExpense(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
