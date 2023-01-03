package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wasawaz/assessment/entity"
	"github.com/wasawaz/assessment/usecase"
)

type CreateExpenseHandler struct {
	createExpenseUsecase usecase.ICreateExpenseUsecase
}

type createExpense struct {
	Title  string   `json:"title"`
	Amount float32  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

func NewCreateExpenseHandler(createExpenseUsecase usecase.ICreateExpenseUsecase) *CreateExpenseHandler {
	return &CreateExpenseHandler{createExpenseUsecase}
}

func (e *CreateExpenseHandler) CreateExpense(c echo.Context) error {
	expense := &createExpense{}
	err := c.Bind(expense)
	if err != nil {
		log.Printf("cannot binding payload")
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
