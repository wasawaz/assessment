// Package router contains functions for expense domain route
package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	handler "github.com/wasawaz/assessment/controller/handler"
)

// New is used for create init application route
func New(e *echo.Echo, createExpenseHandler *handler.CreateExpenseHandler,
	getExpenseHandler *handler.GetExpenseHandler, updateExpenseHandler *handler.UpdateExpenseHandler,
	getAllExpenseHandler *handler.GetAllExpenseHandler) {
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	newExpenseRoute(e, createExpenseHandler, getExpenseHandler, updateExpenseHandler, getAllExpenseHandler)
}

// newExpenseRoute is used for initial expense route
func newExpenseRoute(e *echo.Echo, createExpenseHandler *handler.CreateExpenseHandler, getExpenseHandler *handler.GetExpenseHandler,
	updateExpenseHandler *handler.UpdateExpenseHandler, getAllExpenseHandler *handler.GetAllExpenseHandler) {
	e.POST("/expenses", createExpenseHandler.CreateExpense)
	e.GET("/expenses/:id", getExpenseHandler.GetExpense)
	e.PUT("/expenses/:id", updateExpenseHandler.UpdateExpense)
	e.GET("/expenses", getAllExpenseHandler.GetAllExpense)
}
