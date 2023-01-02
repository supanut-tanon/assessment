package customer

import (
	"database/sql"
	"log"
	

	_ "github.com/lib/pq"
	_ "github.com/proullon/ramsql/driver"
)

var db *sql.DB

func init() {
	conn, err := sql.Open("postgres","postgres://wddhwbsh:dYftALDi3cTkaIk-ONyvjjUh9Z_jxMH3@tiny.db.elephantsql.com/wddhwbsh")
	if err != nil {
		log.Fatal("can't connect to database", err)
	}
	db = conn
	createCustomerTable()
}

func createCustomerTable() {
	createCustomerTable := `
	CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);
	`

	if _, err := db.Exec(createCustomerTable); err != nil {
		log.Fatal("can't create table ", err)
	}
}
