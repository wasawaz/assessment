// Package handler contains functions for http request handling separate route name
package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wasawaz/assessment/usecase"
)

// GetExpenseHandler is used for get expense by id handler
type GetExpenseHandler struct {
	getExpenseUsecase usecase.IGetExpenseUsecase
}

// NewGetExpenseHandler is used for new GetExpenseHandler instance
func NewGetExpenseHandler(getExpenseUsecase usecase.IGetExpenseUsecase) *GetExpenseHandler {
	return &GetExpenseHandler{getExpenseUsecase}
}

// GetExpense is used for handle http request for get expense by id
func (e *GetExpenseHandler) GetExpense(c echo.Context) error {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 32)
	if err != nil {
		c.NoContent(http.StatusNotFound)
	}
	expense, err := e.getExpenseUsecase.Execute(int(id))
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return c.NoContent(http.StatusNotFound)
		default:
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusOK, expense)
}
