package repository

import (
	"database/sql"

	"forum/internal/entities"
)

type Post interface {
	CreatePost(userId int, body, title string) (int, error)
	// GetAllPosts() ([]entities.Post, error)
	GetAllUsersPosts(userId int) ([]entities.Post, error)
	// GetUserPost(userId, postId int) (entities.Post, error)
	// DeletePost(userId, postId int) (int, error)
	// UpdatePost(userId int, post entities.Post) (int, error)
	// PostAction(userId, postId int) (int, error) // like or dislike post / remove like or dislike
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
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Post: NewPostRepo(db),
		User: NewUserRepo(db),
		Auth: NewAuthRepo(db),
	}
}
