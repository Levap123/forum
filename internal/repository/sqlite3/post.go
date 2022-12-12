package repository

import (
	"database/sql"
	"fmt"

	"forum/internal/entities"
	"forum/pkg/errors"
)

type PostRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (pr *PostRepo) CreatePost(userId int, title, body string) (int, error) {
	var postId int
	tx, err := pr.db.Begin()
	if err != nil {
		return 0, errors.Fail(err, "Create post")
	}
	defer tx.Rollback()
	query := fmt.Sprintf("INSERT INTO %s (title, body, user_id) VALUES ($1, $2, $3) RETURNING id", postsTable)
	row := tx.QueryRow(query, title, body, userId)
	if err := row.Scan(&postId); err != nil {
		return 0, errors.Fail(err, "Create post")
	}
	return postId, tx.Commit()
}

func (pr *PostRepo) GetAllUsersPosts(userId int) ([]entities.Post, error) {
	var posts []entities.Post
	tx, err := pr.db.Begin()
	if err != nil {
		return nil, errors.Fail(err, "Get all users posts")
	}
	defer tx.Rollback()
	query := fmt.Sprintf("SELECT id, title, body, actions  FROM %s WHERE user_id = $1", postsTable)
	rows, err := tx.Query(query, userId)
	if err != nil {
		return nil, errors.Fail(err, "Get all users posts")
	}
	for rows.Next() {
		var postBuffer entities.Post
		if err := rows.Scan(&postBuffer.Id, &postBuffer.Title, &postBuffer.Body, &postBuffer.Actions); err != nil {
			return nil, errors.Fail(err, "Get all users posts")
		}
		posts = append(posts, postBuffer)
	}
	return posts, nil
}
