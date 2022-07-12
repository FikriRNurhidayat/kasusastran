package postgres_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/query"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository/postgres"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockQuery "github.com/fikrirnurhidayat/kasusastran/mocks/domain/query"
)

type MockPostgresWulanganRepository struct {
	db *mockQuery.Querier
}

func TestPostgresWulanganRepository_Create(t *testing.T) {
	type input struct {
		ctx       context.Context
		iwulangan *entity.Wulangan
	}

	type output struct {
		wulangan *entity.Wulangan
		err      error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockPostgresWulanganRepository, *input, *output)
	}

	tests := []scenario{
		{
			name: "db.CreateWulangan return error",
			in: &input{
				ctx: context.Background(),
				iwulangan: &entity.Wulangan{
					Title:             "Lorem Ipsum",
					Description:       "Lorem Ipsum",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
			},
			out: &output{
				wulangan: nil,
				err:      fmt.Errorf("CreateWulangan: failed executing db query: Bababoey"),
			},
			on: func(m *MockPostgresWulanganRepository, in *input, out *output) {
				m.db.On("CreateWulangan", in.ctx, mock.AnythingOfType("*query.CreateWulanganParams")).Return(query.CreateWulanganRow{}, out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				iwulangan: &entity.Wulangan{
					Title:             "Lorem Ipsum",
					Description:       "Lorem Ipsum",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
			},
			out: &output{
				wulangan: &entity.Wulangan{
					ID:                uuid.New(),
					Title:             "Jayabaya",
					Description:       "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
				err: nil,
			},
			on: func(m *MockPostgresWulanganRepository, in *input, out *output) {
				row := query.CreateWulanganRow{
					ID:                out.wulangan.ID,
					Title:             out.wulangan.Title,
					Description:       out.wulangan.Description,
					CoverImageUrl:     out.wulangan.CoverImageUrl,
					ThumbnailImageUrl: out.wulangan.ThumbnailImageUrl,
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				}

				out.wulangan.CreatedAt = row.CreatedAt
				out.wulangan.UpdatedAt = row.UpdatedAt

				m.db.On("CreateWulangan", in.ctx, mock.AnythingOfType("*query.CreateWulanganParams")).Return(row, out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockPostgresWulanganRepository{
				db: new(mockQuery.Querier),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := postgres.NewPostgresWulanganRepository(m.db)
			out, err := subject.Create(tt.in.ctx, tt.in.iwulangan)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), tt.out.err.Error())
			}

			assert.Equal(t, tt.out.wulangan, out)
		})
	}
}

func TestWulanganRepository_GetWulangan(t *testing.T) {
	type input struct {
		ctx context.Context
		id  uuid.UUID
	}

	type output struct {
		wulangan *entity.Wulangan
		err      error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockPostgresWulanganRepository, *input, *output)
	}

	tests := []scenario{
		{
			name: "db.GetWulangan return error",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				wulangan: nil,
				err:      fmt.Errorf("GetWulangan: failed executing db query: Bababoey"),
			},
			on: func(m *MockPostgresWulanganRepository, in *input, out *output) {
				m.db.On("GetWulangan", in.ctx, in.id).Return(query.GetWulanganRow{}, out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				wulangan: &entity.Wulangan{
					ID:                uuid.New(),
					Title:             "Jayabaya",
					Description:       "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
				err: nil,
			},
			on: func(m *MockPostgresWulanganRepository, in *input, out *output) {
				row := query.GetWulanganRow{
					ID:                out.wulangan.ID,
					Title:             out.wulangan.Title,
					Description:       out.wulangan.Description,
					CoverImageUrl:     out.wulangan.CoverImageUrl,
					ThumbnailImageUrl: out.wulangan.ThumbnailImageUrl,
					CreatedAt:         out.wulangan.CreatedAt,
					UpdatedAt:         out.wulangan.UpdatedAt,
				}
				m.db.On("GetWulangan", in.ctx, in.id).Return(row, out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockPostgresWulanganRepository{
				db: new(mockQuery.Querier),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := postgres.NewPostgresWulanganRepository(m.db)
			out, err := subject.Get(tt.in.ctx, tt.in.id)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), tt.out.err.Error())
			}

			assert.Equal(t, tt.out.wulangan, out)
		})
	}
}
