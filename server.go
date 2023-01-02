package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/wasawaz/assessment/pkg/httpserver"
)

func main() {
	appPort := os.Getenv("PORT")
	e := echo.New()
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	httpserver.New(e, appPort)

}
