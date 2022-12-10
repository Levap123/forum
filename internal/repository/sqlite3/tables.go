package repository

import (
	"database/sql"
	"forum/pkg/errors"
	"io/ioutil"
	"os"
)

const tableSchemas = "pkg/schemas/up.sql"

func createTables(db *sql.DB) error {
	f, err := os.OpenFile(tableSchemas, os.O_RDONLY, 0755)
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
