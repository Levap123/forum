package repository

import (
	"database/sql"
	"fmt"

	"forum/internal/entities"
	"forum/pkg/errors"
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
		return 0, errors.Fail(err, "CreateUser")
	}

	defer tx.Rollback()

	query := fmt.Sprint("INSERT INTO %s (email, name, password) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := tx.QueryRow(query, user.Email, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, errors.Fail(err, "CreateUser")
	}
	return id, nil
}

func (ur *AuthRepo) GetUser(email, password string) (entities.User, error) {
	var user entities.User
	tx, err := ur.db.Begin()
	if err != nil {
		return user, errors.Fail(err, "GetUser")
	}

	defer tx.Rollback()

	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1 and password=$2", usersTable)
	row := tx.QueryRow(query, email, password)
	if err := row.Scan(&user.Id, &user.Email, &user.Username); err != nil {
		return user, errors.Fail(err, "GetUser")
	}
	return user, nil
}
