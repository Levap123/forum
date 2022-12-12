package service

import (
	"forum/internal/entities"
	repository "forum/internal/repository/sqlite3"
)

type Post interface {
	CreatePost(userId int, title, body string) (int, error)
	GetAllPosts() ([]entities.Post, error)
	GetAllUsersPosts(userId int) ([]entities.Post, error)
	GetPostByPostId(postId int) (entities.Post, error)
	// DeletePost(userId, postId int) (int, error)
	// UpdatePost(userId int, post entities.Post) (int, error)
}

type User interface {
	// DeleteUser(userId int) (int, error)
	GetAllUsers() ([]entities.User, error)
	GetUserById(id int) (entities.User, error)
}

type Auth interface {
	CreateUser(user entities.User) (int, error)
	CreateSession(email, password string) (string, error)
	DeleteSession(id int) error
	GetIdFromSession(uuid string) (int, error)
}

type Action interface {
	VotePost(userId, postId int, vote string) error
	GetPostVotes(postId int) (int, int, error)
}
type Service struct {
	User
	Post
	Auth
	Action
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Post:   NewPostService(repo.Post),
		User:   NewUserService(repo.User),
		Auth:   NewAuthService(repo.Auth),
		Action: NewActionService(repo.Action),
	}
}
