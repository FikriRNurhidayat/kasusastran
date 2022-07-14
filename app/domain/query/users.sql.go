// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: users.sql

package query

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
       name,
       email,
       encrypted_password,
       created_at,
       updated_at
)
VALUES (
  $1,
  $2,
  $3,
  NOW(),
  NOW()
)
RETURNING
    id,
    name,
    email,
    encrypted_password,
    created_at,
    updated_at
`

type CreateUserParams struct {
	Name              string         `json:"name"`
	Email             sql.NullString `json:"email"`
	EncryptedPassword sql.NullString `json:"encrypted_password"`
}

type CreateUserRow struct {
	ID                uuid.UUID      `json:"id"`
	Name              string         `json:"name"`
	Email             sql.NullString `json:"email"`
	EncryptedPassword sql.NullString `json:"encrypted_password"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg *CreateUserParams) (CreateUserRow, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, arg.Name, arg.Email, arg.EncryptedPassword)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.EncryptedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const doesUserEmailExist = `-- name: DoesUserEmailExist :one
SELECT EXISTS (
       SELECT email FROM users
       WHERE users.deleted_at IS NULL
       AND users.email = $1
)
`

func (q *Queries) DoesUserEmailExist(ctx context.Context, email sql.NullString) (bool, error) {
	row := q.queryRow(ctx, q.doesUserEmailExistStmt, doesUserEmailExist, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, encrypted_password, created_at, updated_at
FROM users
WHERE email = $1
AND deleted_at IS NULL
`

type GetUserByEmailRow struct {
	ID                uuid.UUID      `json:"id"`
	Name              string         `json:"name"`
	Email             sql.NullString `json:"email"`
	EncryptedPassword sql.NullString `json:"encrypted_password"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, email sql.NullString) (GetUserByEmailRow, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.EncryptedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
