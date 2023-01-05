package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	handler "github.com/wasawaz/assessment/controller/handler"
)

func New(e *echo.Echo, createExpenseHandler *handler.CreateExpenseHandler,
	getExpenseHandler *handler.GetExpenseHandler) {
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	newExpenseRoute(e, createExpenseHandler, getExpenseHandler)
}
