-- name: CreatePost :one
INSERT INTO posts (id, feed_id, title, url, description, published_at, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostsForUser :many
SELECT po.* FROM posts AS po
INNER JOIN feed_follows AS ff ON ff.feed_id = po.feed_id
WHERE ff.user_id = $1
ORDER BY po.published_at DESC
LIMIT $2;