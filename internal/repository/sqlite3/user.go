package repository

import (
	"database/sql"
	"fmt"
	"forum/internal/entities"
	"forum/pkg/errors"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (ur *UserRepo) GetUserById(id int) (entities.User, error) {
	var user entities.User
	tx, err := ur.db.Begin()
	if err != nil {
		return user, errors.Fail(err, "Get user")
	}
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	row := tx.QueryRow(query, id)
	if err := row.Scan(&user.Id, &user.Email, &user.Username); err != nil {
		return user, errors.Fail(err, "Get user")
	}
	return user, nil
}
