package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateExpenseHandler struct{}

func NewCreateExpenseHandler() *CreateExpenseHandler {
	return &CreateExpenseHandler{}
}

func (e *CreateExpenseHandler) CreateExpense(c echo.Context) error {
	return c.NoContent(http.StatusCreated)
}
