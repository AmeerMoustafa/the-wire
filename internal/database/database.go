package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "../internal/database/users.db")

	if err != nil {
		log.Fatal("DB connection error: ", err)
	}

	db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(25) NOT NULL UNIQUE,
		password_hash VARCHAR(255) NOT NULL,
    	salt VARCHAR(50) NOT NULL,
	);`)

	return db
}
