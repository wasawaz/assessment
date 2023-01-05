package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handler "github.com/wasawaz/assessment/controller/handler"
	"github.com/wasawaz/assessment/controller/handler/customvalidator"
	"github.com/wasawaz/assessment/controller/router"
	"github.com/wasawaz/assessment/pkg/httpserver"
	"github.com/wasawaz/assessment/pkg/postgresql"
	"github.com/wasawaz/assessment/repository"
	"github.com/wasawaz/assessment/usecase"
)

func main() {

	// init database
	pg, err := initDatabase()
	if err != nil {
		log.Fatalf("cannot init db cause %v", err)
	}
	defer pg.Close()

	expenseRepository := repository.NewExpenseRepository(pg)
	createExpenseUsecase := usecase.NewCreateExpenseUsecase(expenseRepository)

	// init httpserver
	httpServer := initHttpServer(createExpenseUsecase)

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

func initHttpServer(createExpenseUsecase usecase.ICreateExpenseUsecase) *httpserver.Server {
	appPort := os.Getenv("PORT")
	e := echo.New()
	e.Validator = customvalidator.NewCustomValidator(validator.New())
	e.Use(middleware.Logger())
	createExpenseHandler := handler.NewCreateExpenseHandler(createExpenseUsecase)
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
