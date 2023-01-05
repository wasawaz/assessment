package router

import (
	"github.com/labstack/echo/v4"
	handler "github.com/wasawaz/assessment/controller/handler"
)

func newExpenseRoute(e *echo.Echo, createExpenseHandler *handler.CreateExpenseHandler, getExpenseHandler *handler.GetExpenseHandler,
	updateExpenseHandler *handler.UpdateExpenseHandler) {
	e.POST("/expenses", createExpenseHandler.CreateExpense)
	e.GET("/expenses/:id", getExpenseHandler.GetExpense)
	e.PUT("/expenses/:id", updateExpenseHandler.UpdateExpense)
}
