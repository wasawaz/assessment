package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handler "github.com/wasawaz/assessment/controllers/handlers"
	"github.com/wasawaz/assessment/controllers/router"
	"github.com/wasawaz/assessment/pkg/httpserver"
	"github.com/wasawaz/assessment/pkg/postgresql"
)

func main() {

	pg, err := initDatabase()
	if err != nil {
		log.Fatalf("cannot init db cause %v", err)
	}
	defer pg.Close()
	httpServer := initHttpServer()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Print("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		log.Fatalf("app - Run - httpServer.Notify: %v", err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Fatalf("app - Run - httpServer.Shutdown: %v", err)
	}

}

func initHttpServer() *httpserver.Server {
	appPort := os.Getenv("PORT")
	e := echo.New()
	e.Use(middleware.Logger())
	createExpenseHandler := handler.NewCreateExpenseHandler()
	router.New(e, createExpenseHandler)
	httpServer := httpserver.New(e, appPort)
	return httpServer
}

func initDatabase() (*postgresql.PostgresqlDB, error) {
	databaseConnectionString := os.Getenv("DATABASE_URL")
	pg, err := postgresql.New(databaseConnectionString)
	if err != nil {
		return pg, err
	}
	_, err = pg.Db.Exec(`CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);`)
	return pg, err
}
