package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	Db   *sql.DB
	name string
	uri  string
}

func NewConnection() *Connection {
	db, err := sql.Open("mysql", "guigui:guigui@tcp(localhost:3305)/bank_db")

	if err != nil {
		panic(err)
	}

	return &Connection{
		Db:   db,
		name: "mysql",
		uri:  "guigui:guigui@tcp(localhost:3305)/bank_db",
	}
}