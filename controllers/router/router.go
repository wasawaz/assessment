package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func New(e *echo.Echo) {
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	newExpenseRoute(e)
}
