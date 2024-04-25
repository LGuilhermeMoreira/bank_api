package database

import "database/sql"

type connection struct {
	Db   *sql.DB
	name string
	uri  string
}

func NewConnection() *connection {
	db, err := sql.Open("mysql", "guigui:guigui@tcp(localhost:3305)/bank_db")

	if err != nil {
		panic(err)
	}

	return &connection{
		Db:   db,
		name: "mysql",
		uri:  "guigui:guigui@tcp(localhost:3305)/bank_db",
	}
}
