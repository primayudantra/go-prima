-- name: CreatePosts :exec
INSERT INTO posts (content) VALUES ($1);