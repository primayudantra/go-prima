-- name: Registration :exec
INSERT INTO users (email, password_hash) VALUES ($1, $2);

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateProfile :exec
UPDATE users set email = $1, password_hash = $2 WHERE id = $3;