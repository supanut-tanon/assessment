package main

import (
	"net/http"
	"os"

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
	r.Start(":2565")
}
