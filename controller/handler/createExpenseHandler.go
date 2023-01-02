package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateExpenseHandler struct{}

type createExpense struct {
	Title  string   `json:"title"`
	Amount float32  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

func NewCreateExpenseHandler() *CreateExpenseHandler {
	return &CreateExpenseHandler{}
}

func (e *CreateExpenseHandler) CreateExpense(c echo.Context) error {
	expense := &createExpense{}
	err := c.Bind(expense)
	if err != nil {
		log.Printf("cannot biding payload")
	}
	return c.JSON(http.StatusCreated, expense)
}
