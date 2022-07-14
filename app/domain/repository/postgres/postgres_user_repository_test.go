package postgres_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/query"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository/postgres"
	mockQuery "github.com/fikrirnurhidayat/kasusastran/mocks/domain/query"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type MockPostgresUserRepository struct {
	db *mockQuery.Querier
}

func TestPostgresUserRepository_EmailExist(t *testing.T) {
	type input struct {
		ctx   context.Context
		email string
	}

	type output struct {
		exist bool
		err   error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockPostgresUserRepository, *input, *output)
	}

	tests := []scenario{
		{
			name: "r.db.DoesUserEmailExist return error",
			in: &input{
				ctx:   context.Background(),
				email: "fikrirnurhidayat@gmail.com",
			},
			out: &output{
				exist: false,
				err:   errors.New("r.db.DoesUserEmailExist: failed to check user email existence"),
			},
			on: func(mpur *MockPostgresUserRepository, i *input, o *output) {
				mpur.db.On("DoesUserEmailExist", i.ctx, sql.NullString{
					String: i.email,
					Valid:  true,
				}).Return(o.exist, o.err)
			},
		},
		{
			name: "OK: false",
			in: &input{
				ctx:   context.Background(),
				email: "fikrirnurhidayat@gmail.com",
			},
			out: &output{
				exist: false,
				err:   nil,
			},
			on: func(mpur *MockPostgresUserRepository, i *input, o *output) {
				mpur.db.On("DoesUserEmailExist", i.ctx, sql.NullString{String: i.email, Valid: true}).Return(o.exist, o.err)
			},
		},
		{
			name: "OK: true",
			in: &input{
				ctx:   context.Background(),
				email: "fikrirnurhidayat@gmail.com",
			},
			out: &output{
				exist: true,
				err:   nil,
			},
			on: func(mpur *MockPostgresUserRepository, i *input, o *output) {
				mpur.db.On("DoesUserEmailExist", i.ctx, sql.NullString{String: i.email, Valid: true}).Return(o.exist, o.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockPostgresUserRepository{
				db: new(mockQuery.Querier),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := postgres.NewPostgresUserRepository(m.db)
			exist, err := subject.EmailExist(tt.in.ctx, tt.in.email)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}

			assert.Equal(t, tt.out.exist, exist)
		})
	}
}

func TestPostgresUserRepository_Create(t *testing.T) {
	type input struct {
		ctx context.Context
		i   *entity.User
	}

	type output struct {
		o   *entity.User
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockPostgresUserRepository, *input, *output)
	}

	tests := []scenario{
		{
			name: "r.db.CreateUser return error",
			in: &input{
				ctx: context.Background(),
				i: &entity.User{
					Name:              "Fikri",
					Email:             "fikrirnurhidayat@gmail.com",
					EncryptedPassword: "$2a$10$O2pp2NvX/Y3OgBM66NUrjOtASWhg3rMhft4X0Ii4U8gX3AOJqcItK",
				},
			},
			out: &output{
				o:   nil,
				err: errors.New("r.db.Create: failed to insert user"),
			},
			on: func(mpur *MockPostgresUserRepository, i *input, o *output) {
				params := &query.CreateUserParams{
					Name: i.i.Name,
					Email: sql.NullString{
						String: i.i.Email,
						Valid:  true,
					},
					EncryptedPassword: sql.NullString{
						String: i.i.EncryptedPassword,
						Valid:  true,
					},
				}

				mpur.db.On("CreateUser", i.ctx, params).Return(query.CreateUserRow{}, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				i: &entity.User{
					Name:              "Fikri",
					Email:             "fikrirnurhidayat@gmail.com",
					EncryptedPassword: "$2a$10$O2pp2NvX/Y3OgBM66NUrjOtASWhg3rMhft4X0Ii4U8gX3AOJqcItK",
				},
			},
			out: &output{
				o: &entity.User{
					ID:                uuid.New(),
					Name:              "Fikri",
					Email:             "fikrirnurhidayat@gmail.com",
					EncryptedPassword: "$2a$10$O2pp2NvX/Y3OgBM66NUrjOtASWhg3rMhft4X0Ii4U8gX3AOJqcItK",
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				},
				err: nil,
			},
			on: func(mpur *MockPostgresUserRepository, i *input, o *output) {
				params := &query.CreateUserParams{
					Name: i.i.Name,
					Email: sql.NullString{
						String: i.i.Email,
						Valid:  true,
					},
					EncryptedPassword: sql.NullString{
						String: i.i.EncryptedPassword,
						Valid:  true,
					},
				}

				mpur.db.On("CreateUser", i.ctx, params).Return(query.CreateUserRow{
					ID:   o.o.ID,
					Name: o.o.Name,
					Email: sql.NullString{
						Valid:  true,
						String: o.o.Email,
					},
					EncryptedPassword: sql.NullString{
						String: o.o.EncryptedPassword,
						Valid:  true,
					},
					CreatedAt: o.o.CreatedAt,
					UpdatedAt: o.o.UpdatedAt,
				}, o.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockPostgresUserRepository{
				db: new(mockQuery.Querier),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := postgres.NewPostgresUserRepository(m.db)
			o, err := subject.Create(tt.in.ctx, tt.in.i)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}

			if tt.out.o != nil {
				assert.Equal(t, tt.out.o, o)
			}
		})
	}
}

func TestPostgresUserRepository_GetByEmail(t *testing.T) {
	type input struct {
		ctx   context.Context
		email string
	}

	type output struct {
		user *entity.User
		err  error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockPostgresUserRepository, *input, *output)
	}

	tests := []scenario{
		{
			name: "r.db.GetByEmailUser return error",
			in: &input{
				ctx:   context.Background(),
				email: "fikrirnurhidayat@gmail.com",
			},
			out: &output{
				user: nil,
				err:  errors.New("r.db.GetUserByEmail: failed to retrieve user"),
			},
			on: func(mpur *MockPostgresUserRepository, i *input, o *output) {
				mpur.db.On("GetUserByEmail", i.ctx, sql.NullString{
					String: i.email,
					Valid:  true,
				}).Return(query.GetUserByEmailRow{}, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:   context.Background(),
				email: "fikrirnurhidayat@gmail.com",
			},
			out: &output{
				user: &entity.User{
					ID:                uuid.New(),
					Name:              "Fikri",
					Email:             "fikrirnurhidayat@gmail.com",
					EncryptedPassword: "$2a$10$O2pp2NvX/Y3OgBM66NUrjOtASWhg3rMhft4X0Ii4U8gX3AOJqcItK",
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				},
				err: nil,
			},
			on: func(mpur *MockPostgresUserRepository, i *input, o *output) {
				row := query.GetUserByEmailRow{
					ID:   o.user.ID,
					Name: o.user.Name,
					Email: sql.NullString{
						String: o.user.Email,
						Valid:  true,
					},
					EncryptedPassword: sql.NullString{
						String: o.user.EncryptedPassword,
						Valid:  true,
					},
					CreatedAt: o.user.CreatedAt,
					UpdatedAt: o.user.UpdatedAt,
				}
				mpur.db.On("GetUserByEmail", i.ctx, sql.NullString{String: i.email, Valid: true}).Return(row, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockPostgresUserRepository{
				db: new(mockQuery.Querier),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := postgres.NewPostgresUserRepository(m.db)
			o, err := subject.GetByEmail(tt.in.ctx, tt.in.email)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}

			if tt.out.user != nil {
				assert.Equal(t, tt.out.user, o)
			}
		})
	}
}
