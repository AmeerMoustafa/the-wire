package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() {
	db, err := sql.Open("sqlite3", "../internal/database/users.db")

	if err != nil {
		log.Fatal(err)
	}
	tables, err := db.Exec("SELECT * FROM users")

	fmt.Println(err)
	fmt.Println(tables)

}
