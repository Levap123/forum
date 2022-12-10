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
	user.Password = generatePasswordHash(user.Password)
	return as.repo.CreateUser(user)
}

func (as *AuthService) CreateSession(email, password string) (string, error) {
	uuid, err := as.repo.CreateSession(email, generatePasswordHash(password))
	if err != nil {
		return "", errors.Fail(err, "Create Session")
	}
	return uuid, nil
}

func (as *AuthService) GetIdFromSession(uuid string) (int, error) {
	return as.repo.GetIdFromSession(uuid)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
