// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package query

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	CountSerats(ctx context.Context) (int64, error)
	CountWulangans(ctx context.Context) (int64, error)
	CreateSerat(ctx context.Context, arg *CreateSeratParams) (CreateSeratRow, error)
	CreateUser(ctx context.Context, arg *CreateUserParams) (CreateUserRow, error)
	CreateWulangan(ctx context.Context, arg *CreateWulanganParams) (CreateWulanganRow, error)
	DeleteSerat(ctx context.Context, id uuid.UUID) error
	DeleteWulangan(ctx context.Context, id uuid.UUID) error
	DoesUserEmailExist(ctx context.Context, email sql.NullString) (bool, error)
	GetSerat(ctx context.Context, id uuid.UUID) (GetSeratRow, error)
	GetUserByEmail(ctx context.Context, email sql.NullString) (GetUserByEmailRow, error)
	GetWulangan(ctx context.Context, id uuid.UUID) (GetWulanganRow, error)
	ListSerats(ctx context.Context, arg *ListSeratsParams) ([]ListSeratsRow, error)
	ListWulangans(ctx context.Context, arg *ListWulangansParams) ([]ListWulangansRow, error)
	UpdateSerat(ctx context.Context, arg *UpdateSeratParams) (UpdateSeratRow, error)
	UpdateWulangan(ctx context.Context, arg *UpdateWulanganParams) (UpdateWulanganRow, error)
}

var _ Querier = (*Queries)(nil)
