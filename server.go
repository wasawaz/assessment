package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(middleware.Logger())
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	httpServer := httpserver.New(e, appPort)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Print("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Fatalf("app - Run - httpServer.Notify: %v", err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Fatalf("app - Run - httpServer.Shutdown: %v", err)
	}

}
