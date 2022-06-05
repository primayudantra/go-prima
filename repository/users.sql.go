// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: users.sql

package repository

import (
	"context"
	"database/sql"
)

const getUser = `-- name: GetUser :one
SELECT id, created_at, updated_at, email, password_hash FROM users
WHERE email = $1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.PasswordHash,
	)
	return i, err
}

const registration = `-- name: Registration :exec
INSERT INTO users (email, password_hash) VALUES ($1, $2)
`

type RegistrationParams struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func (q *Queries) Registration(ctx context.Context, arg RegistrationParams) error {
	_, err := q.db.ExecContext(ctx, registration, arg.Email, arg.PasswordHash)
	return err
}

const updateProfile = `-- name: UpdateProfile :exec
UPDATE users set email = $1, password_hash = $2 WHERE id = $3
`

type UpdateProfileParams struct {
	Email        string        `json:"email"`
	PasswordHash string        `json:"password_hash"`
	ID           sql.NullInt64 `json:"id"`
}

func (q *Queries) UpdateProfile(ctx context.Context, arg UpdateProfileParams) error {
	_, err := q.db.ExecContext(ctx, updateProfile, arg.Email, arg.PasswordHash, arg.ID)
	return err
}
