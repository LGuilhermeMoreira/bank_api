package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	Db *sql.DB
}

func NewConnection(name, uri string) *Connection {
	db, err := sql.Open(name, uri)

	if err != nil {
		panic(err)
	}

	return &Connection{
		Db: db,
	}
}
