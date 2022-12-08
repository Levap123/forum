package service

import (
	"forum/internal/entities"
	repository "forum/internal/repository/sqlite3"
)

type Post interface {
	CreatePost(userId int, post entities.Post) (int, error)
	GetAllPosts() ([]entities.Post, error)
	GetAllUsersPosts(userId int) ([]entities.Post, error)
	GetUserPost(userId, postId int) (entities.Post, error)
	DeletePost(userId, postId int) (int, error)
	UpdatePost(userId int, post entities.Post) (int, error)
	PostAction(userId, postId int) (int, error) // like or dislike post / remove like or dislike
}

type User interface {
	CreateUser(user entities.User) (int, error)
	GetUser(email, password string) (entities.User, error)
	// DeleteUser(userId int) (int, error)
	// GetAllUsers() ([]entities.User, error)
}

type Service struct {
	User
	Post
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo.User),
	}
}
