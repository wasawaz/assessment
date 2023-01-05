package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wasawaz/assessment/usecase"
)

type GetAllExpenseHandler struct {
	getAllExpenseUsecase usecase.IGetAllExpenseUsecase
}

func NewGetAllExpenseHandler(getAllExpenseUsecase usecase.IGetAllExpenseUsecase) *GetAllExpenseHandler {
	return &GetAllExpenseHandler{getAllExpenseUsecase}
}

func (e *GetAllExpenseHandler) GetAllExpense(c echo.Context) error {
	expenses, err := e.getAllExpenseUsecase.Execute()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, expenses)
}
