// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  email, 
  full_name,
  hashed_password
) VALUES (
  $1, $2, $3
) RETURNING id, email, hashed_password, full_name, created_at
`

type CreateUserParams struct {
	Email          string `json:"email"`
	FullName       string `json:"full_name"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.FullName, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.FullName,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, hashed_password, full_name, created_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.FullName,
		&i.CreatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, email, hashed_password, full_name, created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.FullName,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET
  full_name = $1
WHERE id = $2
RETURNING id, email, hashed_password, full_name, created_at
`

type UpdateUserParams struct {
	FullName string `json:"full_name"`
	ID       int64  `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.FullName, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.FullName,
		&i.CreatedAt,
	)
	return i, err
}
