package service

import (
	"crypto/sha1"
	"fmt"

	"forum/internal/entities"
	repository "forum/internal/repository/sqlite3"
	"forum/pkg/errors"
)

const salt = "I)_#GQ@*&&*DSAFweqwAFytasgf(*DS"

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(auth repository.Auth) *AuthService {
	return &AuthService{repo: auth}
}

func (as *AuthService) CreateUser(user entities.User) (int, error) {
	if err := ValidateEmail(user.Email); err != nil {
		return 0, errors.Fail(err, "Create user")
	}
	user.Password = generatePasswordHash(user.Password)
	return as.repo.CreateUser(user)
}

func (as *AuthService) CreateSession(email, password string) (string, error) {
	return as.repo.CreateSession(email, generatePasswordHash(password))
}

func (as *AuthService) GetIdFromSession(uuid string) (int, error) {
	return as.repo.GetIdFromSession(uuid)
}

func (as *AuthService) DeleteSession(id int) error {
	return as.repo.DeleteSession(id)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
