package service

import (
	"forum/internal/entities"
	repository "forum/internal/repository/sqlite3"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(post repository.Post) *PostService {
	return &PostService{
		repo: post,
	}
}

func (ps *PostService) CreatePost(userId int, body, title string) (int, error) {
	return ps.repo.CreatePost(userId, body, title)
}

func (ps *PostService) GetAllUsersPosts(userId int) ([]entities.Post, error) {
	return ps.repo.GetAllUsersPosts(userId)
}

func (ps *PostService) GetPostByPostId(postId int) (entities.Post, error) {
	return ps.repo.GetPostByPostId(postId)
}
