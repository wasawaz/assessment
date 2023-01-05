package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wasawaz/assessment/usecase"
)

type GetExpenseHandler struct {
	getExpenseUsecase usecase.IGetExpenseUsecase
}

func NewGetExpenseHandler(getExpenseUsecase usecase.IGetExpenseUsecase) *GetExpenseHandler {
	return &GetExpenseHandler{getExpenseUsecase}
}

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
