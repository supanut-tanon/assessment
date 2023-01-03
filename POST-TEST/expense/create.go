package expense

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type Err struct {
	Message string `json:"message"`
}

func CreateHandler(c echo.Context) error {
	cust := Expense{}
	err := c.Bind(&cust)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	row := db.QueryRow("INSERT INTO expenses(title, amount, note, tags) VALUES($1, $2, $3, $4) RETURNING id;", cust.Title, cust.Amount, cust.Note, pq.Array(&cust.Tags))
	err = row.Scan(&cust.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	fmt.Printf("id : % #v\n", cust)

	return c.JSON(http.StatusCreated, cust)
}
