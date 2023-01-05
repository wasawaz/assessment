package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetExpenseHandler struct {
}

func NewGetExpenseHandler() *GetExpenseHandler {
	return &GetExpenseHandler{}
}

func (e *GetExpenseHandler) GetExpense(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
