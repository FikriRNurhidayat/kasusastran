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
		serat *entity.Wulangan
		err   error
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
				serat: nil,
				err:   fmt.Errorf("CreateWulangan: failed executing db query: Bababoey"),
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
				serat: &entity.Wulangan{
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
					ID:                out.serat.ID,
					Title:             out.serat.Title,
					Description:       out.serat.Description,
					CoverImageUrl:     out.serat.CoverImageUrl,
					ThumbnailImageUrl: out.serat.ThumbnailImageUrl,
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				}

				out.serat.CreatedAt = row.CreatedAt
				out.serat.UpdatedAt = row.UpdatedAt

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

			assert.Equal(t, tt.out.serat, out)
		})
	}
}
