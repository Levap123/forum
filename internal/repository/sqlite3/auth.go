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

func (ar *AuthRepo) CreateUser(user entities.User) (int, error) {
	var id int
	tx, err := ar.db.Begin()
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

func (ar *AuthRepo) CreateSession(email, password string) (string, error) {
	id := 0
	tx, err := ar.db.Begin()
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
		fmt.Println(id)
		query = fmt.Sprintf("UPDATE %s SET uuid = $1 WHERE user_id = $2", sessionsTable)
		_, err = tx.Exec(query, uuid, id)
		if err != nil {
			return "", errors.Fail(err, "Create session")
		}
	}
	return uuid, tx.Commit()
}

func (ar *AuthRepo) GetIdFromSession(uuid string) (int, error) {
	id := 0
	tx, err := ar.db.Begin()
	if err != nil {
		return id, errors.Fail(err, "Get id from session")
	}
	defer tx.Rollback()
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE uuid = $1", sessionsTable)
	row := tx.QueryRow(query, uuid)
	if err := row.Scan(&id); err != nil {
		return 0, errors.Fail(err, "Get id from session")
	}
	return id, tx.Commit()
}

func (ar *AuthRepo) DeleteSession(id int) error {
	tx, err := ar.db.Begin()
	if err != nil {
		return errors.Fail(err, "Delete session")
	}
	defer tx.Rollback()
	query := fmt.Sprintf("DELETE FROM %s where user_id = $1", sessionsTable)
	if _, err := tx.Exec(query, id); err != nil {
		return errors.Fail(err, "Delete session")
	}
	return tx.Commit()
}
