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
		return entities.User{}, errors.Fail(err, "Get user")
	}
	defer tx.Rollback()
	query := fmt.Sprintf("SELECT id, email, user_name FROM %s WHERE id = $1", usersTable)
	row := tx.QueryRow(query, id)
	if err := row.Scan(&user.Id, &user.Email, &user.Username); err != nil {
		return entities.User{}, errors.Fail(err, "Get user")
	}
	return user, nil
}

func (ur *UserRepo) GetAllUsers() ([]entities.User, error) {
	users := make([]entities.User, 0)
	tx, err := ur.db.Begin()
	if err != nil {
		return nil, errors.Fail(err, "Get all users")
	}
	defer tx.Rollback()
	query := fmt.Sprintf("SELECT id, email, user_name from %s", usersTable)
	row, err := tx.Query(query)
	for row.Next() {
		var buffer entities.User
		if err := row.Scan(&buffer.Id, &buffer.Email, &buffer.Username); err != nil {
			return nil, errors.Fail(err, "Get all users")
		}
		users = append(users, buffer)
	}
	return users, nil
}
