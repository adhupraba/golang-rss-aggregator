// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: posts.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (id, feed_id, title, url, description, published_at, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, feed_id, title, url, description, published_at, created_at, updated_at
`

type CreatePostParams struct {
	ID          uuid.UUID
	FeedID      uuid.UUID
	Title       string
	Url         string
	Description sql.NullString
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.FeedID,
		arg.Title,
		arg.Url,
		arg.Description,
		arg.PublishedAt,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.FeedID,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.PublishedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPostsForUser = `-- name: GetPostsForUser :many
SELECT po.id, po.feed_id, po.title, po.url, po.description, po.published_at, po.created_at, po.updated_at FROM posts AS po
INNER JOIN feed_follows AS ff ON ff.feed_id = po.feed_id
WHERE ff.user_id = $1
ORDER BY po.published_at DESC
LIMIT $2
`

type GetPostsForUserParams struct {
	UserID uuid.UUID
	Limit  int32
}

func (q *Queries) GetPostsForUser(ctx context.Context, arg GetPostsForUserParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsForUser, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.FeedID,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.PublishedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
