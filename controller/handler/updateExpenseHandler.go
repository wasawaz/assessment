package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wasawaz/assessment/entity"
	"github.com/wasawaz/assessment/usecase"
)

type UpdateExpenseHandler struct {
	updateExpenseUsecase usecase.IUpdateExpenseUsecase
}

type updateExpense struct {
	Title  string   `json:"title" validate:"required"`
	Amount float32  `json:"amount" validate:"required",gt:0`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

func NewUpdateExpenseHandler(updateExpenseUsecase usecase.IUpdateExpenseUsecase) *UpdateExpenseHandler {
	return &UpdateExpenseHandler{updateExpenseUsecase}
}

func (h *UpdateExpenseHandler) UpdateExpense(c echo.Context) error {
	expense := &updateExpense{}
	paramId := c.Param("id")

	id, err := strconv.ParseInt(paramId, 10, 32)
	if err != nil {
		c.NoContent(http.StatusNotFound)
	}

	err = c.Bind(expense)
	if err != nil {
		log.Printf("cannot binding payload")
	}

	if err = c.Validate(expense); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	entity := entity.Expense{
		Id:     int(id),
		Title:  expense.Title,
		Amount: expense.Amount,
		Note:   expense.Note,
		Tags:   expense.Tags,
	}

	err = h.updateExpenseUsecase.Execute(entity)

	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return c.NoContent(http.StatusNotFound)
		default:
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusAccepted, entity)
}
