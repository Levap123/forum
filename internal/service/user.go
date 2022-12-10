package service

import (
	"forum/internal/entities"
	repository "forum/internal/repository/sqlite3"
)

type UserService struct {
	repo repository.User
}

func NewUserService(user repository.User) *UserService {
	return &UserService{repo: user}
}

func (us *UserService) GetUserById(id int) (entities.User, error) {
	return us.repo.GetUserById(id)
}
