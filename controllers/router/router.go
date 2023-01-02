package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	handler "github.com/wasawaz/assessment/controllers/handlers"
)

func New(e *echo.Echo, createExpenseHandler *handler.CreateExpenseHandler) {
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	newExpenseRoute(e, createExpenseHandler)
}
