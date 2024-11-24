package main

import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
)


func initDB() {
	db, err = sql.open("sqlite", "./users.db")
	fmt.println(db)

}