package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func init() {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		panic(err)
	}
	sql := `
	  CREATE TABLE IF NOT EXISTS users (
		userid INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		password TEXT);`

	_, err = db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func Connect() *sql.DB {
	return db
}
