package service

import (
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
