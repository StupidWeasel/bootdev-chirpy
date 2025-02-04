// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, hashed_password)
VALUES ($1, $2)
RETURNING id, created_at, updated_at, email, hashed_password, password_reset_required
`

type CreateUserParams struct {
	Email          string
	HashedPassword string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordResetRequired,
	)
	return i, err
}

const loginUser = `-- name: LoginUser :one
SELECT id, created_at, updated_at, email, hashed_password
FROM users
WHERE (email = $1)
LIMIT 1
`

type LoginUserRow struct {
	ID             uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Email          string
	HashedPassword string
}

func (q *Queries) LoginUser(ctx context.Context, email string) (LoginUserRow, error) {
	row := q.db.QueryRowContext(ctx, loginUser, email)
	var i LoginUserRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.HashedPassword,
	)
	return i, err
}

const usersClear = `-- name: UsersClear :exec
TRUNCATE users, messages, refresh_tokens
`

func (q *Queries) UsersClear(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, usersClear)
	return err
}
