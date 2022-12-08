package service

import (
	"crypto/sha1"
	"fmt"

	"forum/internal/entities"
	repository "forum/internal/repository/sqlite3"
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

func (as *AuthService) GetUser(email, password string) (entities.User, error) {
	panic("")
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
