package repository

import (
	"database/sql"
	"io/ioutil"
	"os"

	"forum/pkg/errors"
)

const (
	tableSchemas  = "up.sql"
	sessionsTable = "sessions"
	usersTable    = "users"
	postsTable    = "posts"
	actionsTable  = "actions"
)

func createTables(db *sql.DB) error {
	f, err := os.OpenFile(tableSchemas, os.O_RDONLY, 0o755)
	if err != nil {
		return errors.Fail(err, "Create tables")
	}
	defer f.Close()
	tables, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	tx, _ := db.Begin()
	_, err = tx.Exec(string(tables))
	if err != nil {
		return errors.Fail(err, "Create tables")
	}
	return tx.Commit()
}
