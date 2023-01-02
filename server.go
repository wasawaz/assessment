package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/wasawaz/assessment/pkg/httpserver"
	"github.com/wasawaz/assessment/pkg/postgresql"
)

func main() {
	databaseConnectionString := os.Getenv("DATABASE_URL")
	pg, err := postgresql.New(databaseConnectionString)
	if err != nil {
		log.Fatalf("cannot connect to db cause %v", err)
	}
	defer pg.Close()

	appPort := os.Getenv("PORT")
	e := echo.New()
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	httpserver.New(e, appPort)

}
