package db

import (
	"database/sql"
	"log"

	// _ "github.com/mattn/go-sqlite3"
	"github.com/nicewook/authjwt/helper"
	_ "modernc.org/sqlite"
)

var (
	db *sql.DB
)

func init() {

	// connect
	var err error
	db, err = sql.Open("sqlite", "users.db")
	// db, err = sql.Open("sqlite3", "users.db")
	if err != nil {
		panic(err)
	}

	// create table
	sql := `
    CREATE TABLE IF NOT EXISTS users (
		username TEXT,
		password TEXT,
		UNIQUE(username));`

	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}

	// create default admin
	// 1. first make password hash
	passwordHash, err := helper.HashPassword("1234")
	if err != nil {
		log.Fatal(err)
	}

	// 2. then insert user
	sqlInsertAdmin := `INSERT INTO users(username, password) VALUES ("admin", ?);`

	_, err = db.Exec(sqlInsertAdmin, passwordHash)
	if err != nil {
		log.Println(err)
	}
}

func Connect() *sql.DB {
	return db
}
