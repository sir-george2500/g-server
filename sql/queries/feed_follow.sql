-- name: CreateFeedFollow :one
INSERT INTO feeds_follow(
  id, created_at, updated_at,user_id,feed_id
) VALUES ($1, $2, $3, $4,$5)
RETURNING *;
