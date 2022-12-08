package repository

import "forum/internal/entities"

type Post interface {
	CreatePost(userId int, post entities.Post) (int, error)
	GetAllPosts() ([]entities.Post, error)
	GetUsersPosts(userId int) ([]entities.Post, error)
	DeletePost(userId, postId int) (int, error)
	UpdatePost(userId int, post entities.Post) (int, error)
}

type Auth interface {
	CreateUser(entities.User) 
}

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}
