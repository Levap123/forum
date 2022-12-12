package repository

import (
	"database/sql"
	"fmt"

	"forum/pkg/errors"
)

type ActionRepo struct {
	db *sql.DB
}

func NewActionRepo(db *sql.DB) *ActionRepo {
	return &ActionRepo{db: db}
}

func (ar *ActionRepo) VotePost(userId, postId int, vote string) error {
	tx, err := ar.db.Begin()
	if err != nil {
		return errors.Fail(err, "Get Post Likes")
	}
	defer tx.Rollback()
	action := 0
	if vote == "like" {
		action = 1
	}
	if vote == "dislike" {
		action = -1
	}
	if action == 0 {
		return errors.Fail(fmt.Errorf("vote type is not like or dislike"), "Vote post")
	}
	query := fmt.Sprintf("INSERT INTO %s (vote, user_id, post_id) VALUES ($1, $2, $3)", actionsTable)

	if _, err := tx.Exec(query, action, userId, postId); err != nil {
		return err
	}
	return tx.Commit()
}

func (ar *ActionRepo) DeleteVote(userId, postId int) (int, error) {
	var vote int
	tx, err := ar.db.Begin()
	if err != nil {
		return 0, errors.Fail(err, "Delete vote")
	}
	defer tx.Rollback()
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 and post_id = $2 RETURNING vote", actionsTable)
	row := tx.QueryRow(query, userId, postId)
	if err := row.Scan(&vote); err != nil {
		return 0, errors.Fail(err, "Delete vote")
	}
	return vote, tx.Commit()
}
