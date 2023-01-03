package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/supanut-tanon/assessment/expense"
)

func setupRoute() *echo.Echo {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
			return next(c)
		}
	})

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if token != os.Getenv("AUTH_TOKEN") {
				return c.JSON(http.StatusUnauthorized, "Unauthorized")
			}

			return next(c)
		}
	})

	e.POST("/expense", expense.CreateHandler)
	e.GET("/expense/:id", expense.GetExpenseByIdHandler)
	e.PUT("/expense/:id", expense.UpdateExpenseHandler)
	e.GET("/expense", expense.GetExpenseHandler)
	

	return e
}

func main() {
	r := setupRoute()

	go func() {
		if err := r.Start(":2565"); err != nil && err != http.ErrServerClosed { 
			r.Logger.Fatal("shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.Shutdown(ctx); err != nil {
		r.Logger.Fatal(err)
	}
}

