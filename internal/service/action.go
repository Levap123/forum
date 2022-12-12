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

func (as *ActionService) VotePost(userId, postId int, vote string) error {
	return as.repo.VotePost(userId, postId, vote)
}

func (as *ActionService) GetPostVotes(postId int) (int, int, error) {
	return as.repo.GetPostVotes(postId)
}
