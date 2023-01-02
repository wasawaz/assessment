package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ExpenseHandler struct{}

func newCreateExpenseHandler() *ExpenseHandler {
	return &ExpenseHandler{}
}

func (e *ExpenseHandler) createExpense(c echo.Context) error {
	return c.NoContent(http.StatusCreated)
}
