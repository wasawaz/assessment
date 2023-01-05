package router

import (
	"github.com/labstack/echo/v4"
	handler "github.com/wasawaz/assessment/controller/handler"
)

func newExpenseRoute(e *echo.Echo, createExpenseHandler *handler.CreateExpenseHandler, getExpenseHandler *handler.GetExpenseHandler) {
	e.POST("/expenses", createExpenseHandler.CreateExpense)
	e.GET("/expenses:id", getExpenseHandler.GetExpense)
}
