package repository

import (
	"database/sql"

	"fmt"

	"forum/internal/entities"
	"forum/pkg/errors"
	"forum/pkg/sessions"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (ur *AuthRepo) CreateUser(user entities.User) (int, error) {
	var id int
	tx, err := ur.db.Begin()
	if err != nil {
		return 0, errors.Fail(err, "Create user")
	}

	defer tx.Rollback()

	query := fmt.Sprintf("INSERT INTO %s (email, user_name, password) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := tx.QueryRow(query, user.Email, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, errors.Fail(err, "Create user")
	}
	return id, tx.Commit()
}

func (ur *AuthRepo) CreateSession(email, password string) (string, error) {
	id := 0
	tx, err := ur.db.Begin()
	if err != nil {
		return "", errors.Fail(err, "Create session")
	}

	defer tx.Rollback()

	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 and password=$2", usersTable)
	row := tx.QueryRow(query, email, password)
	if err := row.Scan(&id); err != nil {
		return "", errors.Fail(err, "Create session")
	}
	uuid, err := sessions.GenerateUuid()
	query = fmt.Sprintf("INSERT INTO %s (user_id, uuid) VALUES ($1, $2)", sessionsTable)
	_, err = tx.Exec(query, id, uuid)
	if err != nil {
		return "", errors.Fail(err, "Create session")
	}
	return uuid, tx.Commit()
}
