package service

import (
	"strings"

	repository "forum/internal/repository/sqlite3"
	"forum/pkg/errors"
)

type ActionService struct {
	repo repository.Action
}

const uniqueError = "UNIQUE constraint failed"

func NewActionService(repo repository.Action) *ActionService {
	return &ActionService{
		repo: repo,
	}
}

func (as *ActionService) VotePost(userId, postId int, vote string) error {
	err := as.repo.VotePost(userId, postId, vote)
	if err != nil {
		if strings.HasPrefix(err.Error(), uniqueError) {
			action, err := as.repo.DeleteVote(userId, postId)
			if err != nil {
				return errors.Fail(err, "Vote post")
			}

			if (vote == "like" && action == -1) || (vote == "dislike" && action == 1) {
				err = as.repo.VotePost(userId, postId, vote)
			}
		}
	}

	if err != nil && !strings.HasPrefix(err.Error(), uniqueError) {
		return errors.Fail(err, "Vote post")
	}

	return nil
}
