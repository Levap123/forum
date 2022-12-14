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

func (pr *PostRepo) GetPostVotes(postId int) (int, int, error) {
	tx, err := pr.db.Begin()
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
	query := fmt.Sprintf("SELECT id, title, body FROM %s WHERE user_id = $1", postsTable)
	rows, err := tx.Query(query, userId)
	if err != nil {
		return nil, errors.Fail(err, "Get all users posts")
	}
	for rows.Next() {
		var buffer entities.Post
		if err := rows.Scan(&buffer.Id, &buffer.Title, &buffer.Body); err != nil {
			return nil, errors.Fail(err, "Get all users posts")
		}
		buffer.Likes, buffer.Dislikes, err = pr.GetPostVotes(buffer.Id)
		if err != nil {
			return nil, errors.Fail(err, "Get all users posts")
		}
		posts = append(posts, buffer)
	}
	return posts, tx.Commit()
}

func (pr *PostRepo) GetPostByPostId(postId int) (entities.Post, error) {
	var post entities.Post
	tx, err := pr.db.Begin()
	if err != nil {
		return entities.Post{}, errors.Fail(err, "Get post by post id")
	}
	defer tx.Rollback()
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postsTable)
	row := tx.QueryRow(query, postId)
	if err := row.Scan(&post.Id, &post.Title, &post.Body, &post.UserId); err != nil {
		return entities.Post{}, err
	}
	post.Likes, post.Dislikes, err = pr.GetPostVotes(post.Id)
	return post, err
}

func (pr *PostRepo) GetAllPosts() ([]entities.Post, error) {
	var posts []entities.Post
	tx, err := pr.db.Begin()
	if err != nil {
		return nil, errors.Fail(err, "Get all posts")
	}
	query := fmt.Sprintf("SELECT * FROM %s", postsTable)
	rows, err := tx.Query(query)
	if err != nil {
		return nil, errors.Fail(err, "Get all users posts")
	}
	for rows.Next() {
		var buffer entities.Post
		if err := rows.Scan(&buffer.Id, &buffer.Title, &buffer.Body, &buffer.UserId); err != nil {
			return nil, errors.Fail(err, "Get all users posts")
		}
		buffer.Likes, buffer.Dislikes, err = pr.GetPostVotes(buffer.Id)
		posts = append(posts, buffer)
	}
	return posts, tx.Commit()
}
