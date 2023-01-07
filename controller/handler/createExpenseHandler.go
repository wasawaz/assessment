// Package handler contains functions for http request handling separate route name
package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wasawaz/assessment/entity"
	"github.com/wasawaz/assessment/usecase"
)

// CreateExpenseHandler is used for create expense handler
type CreateExpenseHandler struct {
	createExpenseUsecase usecase.ICreateExpenseUsecase
}
// createExpense is used for create expense payload
type createExpense struct {
	Title  string   `json:"title" validate:"required"`
	Amount float32  `json:"amount" validate:"required",gt:0`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

// NewCreateExpenseHandler is used for new CreateExpenseHandler instance
func NewCreateExpenseHandler(createExpenseUsecase usecase.ICreateExpenseUsecase) *CreateExpenseHandler {
	return &CreateExpenseHandler{createExpenseUsecase}
}

// CreateExpense is used for handle http request for create new expense
func (e *CreateExpenseHandler) CreateExpense(c echo.Context) error {
	expense := &createExpense{}
	err := c.Bind(expense)
	if err != nil {
		log.Printf("cannot binding payload")
	}

	if err = c.Validate(expense); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	entity := &entity.Expense{
		Title:  expense.Title,
		Amount: expense.Amount,
		Note:   expense.Note,
		Tags:   expense.Tags,
	}
	err = e.createExpenseUsecase.Execute(entity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, entity)
}
