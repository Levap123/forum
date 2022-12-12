package repository

import (
	"database/sql"

	"forum/internal/entities"
)

type Post interface {
	CreatePost(userId int, body, title string) (int, error)
	GetAllPosts() ([]entities.Post, error)
	GetAllUsersPosts(userId int) ([]entities.Post, error)
	GetPostByPostId(postId int) (entities.Post, error)
	// DeletePost(userId, postId int) (int, error)
	// UpdatePost(userId int, post entities.Post) (int, error)
}

type Action interface {
	VotePost(userId, postId int, vote string) (int, error)
	GetPostVotes(postId int) (int, int, error)
}

type User interface {
	// UpdateUser(email, pasword, username string) (entities.User, error)
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

type Repository struct {
	Post
	User
	Auth
	Action
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Post:   NewPostRepo(db),
		User:   NewUserRepo(db),
		Auth:   NewAuthRepo(db),
		Action: NewActionRepo(db),
	}
}
