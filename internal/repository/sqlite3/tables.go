package repository

import (
	"database/sql"
	"fmt"
)

const userTableC = `CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT UNIQUE,
	user_name TEXT,
	password TEXT
)`

func createTables(db *sql.DB) {
	_, err := db.Exec(userTableC)
	if err != nil {
		fmt.Println(err)
	}
}
