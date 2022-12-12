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

const uniqueError = "UNIQUE constraint failed"

func (ar *ActionRepo) GetPostVotes(postId int) (int, int, error) {
	tx, err := ar.db.Begin()
	likes := 0
	dislikes := 0
	if err != nil {
		return 0, 0, errors.Fail(err, "Get Post Likes")
	}
	defer tx.Rollback()
	query := fmt.Sprintf("SELECT vote FROM %s WHERE post_id = $1", actionsTable)
	rows, err := tx.Query(query, postId)
	for rows.Next() {
		var buffer int
		if err := rows.Scan(&buffer); err != nil {
			return 0, 0, errors.Fail(err, "Get post votes")
		}
		if buffer == 1 {
			likes++
		} else {
			dislikes++
		}
	}
	return likes, dislikes, nil
}

func (ar *ActionRepo) VotePost(userId, postId int, vote string) error {
	tx, err := ar.db.Begin()
	if err != nil {
		return errors.Fail(err, "Get Post Likes")
	}
	defer tx.Rollback()
	action := 0
	fmt.Println(vote)
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
