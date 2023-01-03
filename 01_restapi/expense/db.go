package expense

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
	createUserTable()
}

func createUserTable() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`

	if _, err := db.Exec(createUserTable); err != nil {
		log.Fatal("can't create table ", err)
	}
}