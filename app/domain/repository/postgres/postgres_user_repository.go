package postgres

import (
	"context"
	"database/sql"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/query"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository"
	"github.com/google/uuid"
)

type PostgresUserRepository struct {
	db query.Querier
}

// GetByEmail implements repository.UserRepository
func (r *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	row, err := r.db.GetUserByEmail(ctx, sql.NullString{
		String: email,
		Valid:  true,
	})
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		ID:                row.ID,
		Name:              row.Name,
		Email:             row.Email.String,
		EncryptedPassword: row.EncryptedPassword.String,
		CreatedAt:         row.CreatedAt,
		UpdatedAt:         row.UpdatedAt,
	}

	return user, nil
}

func NewPostgresUserRepository(db query.Querier) repository.UserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) EmailExist(ctx context.Context, email string) (bool, error) {
	return r.db.DoesUserEmailExist(ctx, sql.NullString{
		String: email,
		Valid:  true,
	})
}

func (r *PostgresUserRepository) Create(ctx context.Context, i *entity.User) (o *entity.User, err error) {
	row, err := r.db.CreateUser(ctx, &query.CreateUserParams{
		Name: i.Name,
		Email: sql.NullString{
			String: i.Email,
			Valid:  true,
		},
		EncryptedPassword: sql.NullString{
			String: i.EncryptedPassword,
			Valid:  true,
		},
	})

	if err != nil {
		return
	}

	o = &entity.User{
		ID:                row.ID,
		Name:              row.Name,
		Email:             row.Email.String,
		EncryptedPassword: row.EncryptedPassword.String,
		CreatedAt:         row.CreatedAt,
		UpdatedAt:         row.UpdatedAt,
	}

	return o, nil
}

func (r *PostgresUserRepository) Get(ctx context.Context, id uuid.UUID) (User *entity.User, err error) {
	panic("not implemented")
}
