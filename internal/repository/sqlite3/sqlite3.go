package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		return nil, err
	}
	return db, createTables(db)
}
