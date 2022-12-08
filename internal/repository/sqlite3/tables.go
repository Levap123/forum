package repository

import "database/sql"

const userTableC = `CREATE TABLE IF NOT EXIST users (
	id INTEGER PRIMARY KEY,
	email TEXT UNIQUE,
	user_name TEXT,

)`

func createTables(db *sql.DB) {
	db.Exec(userTableC)
}
