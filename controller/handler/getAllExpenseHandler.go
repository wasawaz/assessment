// Package handler contains functions for http request handling separate route name
package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wasawaz/assessment/usecase"
)

// GetAllExpenseHandler is used for get all expense handler
type GetAllExpenseHandler struct {
	getAllExpenseUsecase usecase.IGetAllExpenseUsecase
}

// NewGetAllExpenseHandler is used for new GetAllExpenseHandler instance
func NewGetAllExpenseHandler(getAllExpenseUsecase usecase.IGetAllExpenseUsecase) *GetAllExpenseHandler {
	return &GetAllExpenseHandler{getAllExpenseUsecase}
}

// GetAllExpense is used for handle http request for get all expense
func (e *GetAllExpenseHandler) GetAllExpense(c echo.Context) error {
	expenses, err := e.getAllExpenseUsecase.Execute()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, expenses)
}
