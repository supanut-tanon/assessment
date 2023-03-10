package expense

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func GetExpenseByIdHandler(c echo.Context) error {
	id := c.Param("id")

	rowID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "id should be int " + err.Error()})
	}

	row := db.QueryRow("SELECT id, title, amount, note, tags FROM expenses WHERE id=$1", rowID)

	cst := Expense{}
	err = row.Scan(&cst.ID, &cst.Title, &cst.Amount, &cst.Note, pq.Array(&cst.Tags))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	fmt.Printf("cst % #v\n", cst)

	return c.JSON(http.StatusOK, cst)
}

func GetExpenseHandler(c echo.Context) error {
	custs := []Expense{}

	rows, err := db.Query("SELECT id, title, amount, note, tags FROM expenses")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	for rows.Next() {
		cst := Expense{}
		err := rows.Scan(&cst.ID, &cst.Title, &cst.Amount, &cst.Note, pq.Array(&cst.Tags))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})

		}
		custs = append(custs, cst)
	}

	return c.JSON(http.StatusOK, custs)
}
