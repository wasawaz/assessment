package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func newExpenseRoute(e *echo.Echo) {
	e.POST("/expenses", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
