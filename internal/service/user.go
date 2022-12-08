package service

import repository "forum/internal/repository/sqlite3"

type UserService struct {
	repository.User
}

func NewUserService(user repository.User) *UserService {
	return &UserService{User: user}
}
