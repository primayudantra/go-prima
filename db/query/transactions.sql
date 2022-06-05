-- name: CreateTransaction :exec
INSERT INTO transactions (user_id, amount, currency, status) VALUES ($1, $2, $3, $4);