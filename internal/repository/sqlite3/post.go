package repository

import "database/sql"

type PostRepo struct {
	db *sql.DB
}
