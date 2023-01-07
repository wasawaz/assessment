// Package main configures and runs application.
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
	expense_middleware "github.com/wasawaz/assessment/middleware"
	"github.com/wasawaz/assessment/pkg/httpserver"
	"github.com/wasawaz/assessment/pkg/postgresql"
	"github.com/wasawaz/assessment/repository"
	"github.com/wasawaz/assessment/usecase"
)

func main() {

	// Database
	pg, err := initDatabase()
	if err != nil {
		log.Fatalf("cannot init db cause %v", err)
	}
	defer pg.Close()

	expenseRepository := repository.NewExpenseRepository(pg)
	createExpenseUsecase := usecase.NewCreateExpenseUsecase(expenseRepository)
	getExpenseUsecase := usecase.NewGetExpenseUsecase(expenseRepository)
	updateExpenseUsecase := usecase.NewUpdateExpenseUsecase(expenseRepository)
	getAllExpenseUsecase := usecase.NewGetAllExpenseUsecase(expenseRepository)

	// Httpserver
	httpServer := initHttpServer(createExpenseUsecase, getExpenseUsecase, updateExpenseUsecase, getAllExpenseUsecase)

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

// Init HTTP Server
func initHttpServer(createExpenseUsecase usecase.ICreateExpenseUsecase, getExpenseUsecase usecase.IGetExpenseUsecase,
	updateExpenseUsecase usecase.IUpdateExpenseUsecase, getAllExpenseUsecase usecase.IGetAllExpenseUsecase) *httpserver.Server {
	appPort := os.Getenv("PORT")
	e := echo.New()
	e.Validator = customvalidator.NewCustomValidator(validator.New())
	e.Use(middleware.Logger())
	e.Use(expense_middleware.AuthMiddleware)
	createExpenseHandler := handler.NewCreateExpenseHandler(createExpenseUsecase)
	getExpenseHandler := handler.NewGetExpenseHandler(getExpenseUsecase)
	updateExpenseHandler := handler.NewUpdateExpenseHandler(updateExpenseUsecase)
	getAllExpenseHandler := handler.NewGetAllExpenseHandler(getAllExpenseUsecase)
	router.New(e, createExpenseHandler, getExpenseHandler, updateExpenseHandler, getAllExpenseHandler)
	httpServer := httpserver.New(e, appPort)
	return httpServer
}

// Init Database
func initDatabase() (*postgresql.Postgres, error) {
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
