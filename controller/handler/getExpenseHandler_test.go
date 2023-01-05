package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/wasawaz/assessment/controller/handler/customvalidator"
)

func TestGetExpense(t *testing.T) {
	t.Run("should return http status 200", func(t *testing.T) {
		e := echo.New()
		e.Validator = customvalidator.NewCustomValidator(validator.New())
		req := httptest.NewRequest(http.MethodGet, "/expenses/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := NewGetExpenseHandler()

		// Assertions
		if assert.NoError(t, h.GetExpense(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
