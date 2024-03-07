// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package auth

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, email, password, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?)
RETURNING id, username, email, password, created_at, updated_at
`

type CreateUserParams struct {
	Username  string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserFromID = `-- name: GetUserFromID :one
SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = ?
`

func (q *Queries) GetUserFromID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserFromID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}