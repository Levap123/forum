package service

import repository "forum/internal/repository/sqlite3"

type ActionService struct {
	repo repository.Action
}

func NewActionService(repo repository.Action) *ActionService {
	return &ActionService{
		repo: repo,
	}
}
