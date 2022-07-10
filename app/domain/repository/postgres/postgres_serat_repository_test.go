package postgres_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/query"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/repository/postgres"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockQuery "github.com/fikrirnurhidayat/kasusastran/mocks/domain/query"
)

type MockSeratRepository struct {
	db *mockQuery.Querier
}

func TestSeratRepository_Create(t *testing.T) {
	type input struct {
		ctx    context.Context
		iserat *entity.Serat
	}

	type output struct {
		serat *entity.Serat
		err   error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratRepository, *input, *output)
	}

	tests := []scenario{
		{
			name: "db.CreateSerat return error",
			in: &input{
				ctx: context.Background(),
				iserat: &entity.Serat{
					Title:             "Lorem Ipsum",
					Description:       "Lorem Ipsum",
					Content:           "Lorem Ipsum",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
			},
			out: &output{
				serat: nil,
				err:   fmt.Errorf("CreateSerat: failed executing db query: Bababoey"),
			},
			on: func(m *MockSeratRepository, in *input, out *output) {
				m.db.On("CreateSerat", in.ctx, mock.AnythingOfType("*query.CreateSeratParams")).Return(query.CreateSeratRow{}, out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				iserat: &entity.Serat{
					Title:             "Lorem Ipsum",
					Description:       "Lorem Ipsum",
					Content:           "Lorem Ipsum",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
			},
			out: &output{
				serat: &entity.Serat{
					ID:                uuid.New(),
					Title:             "Jayabaya",
					Description:       "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
				err: nil,
			},
			on: func(m *MockSeratRepository, in *input, out *output) {
				row := query.CreateSeratRow{
					ID:                out.serat.ID,
					Title:             out.serat.Title,
					Description:       out.serat.Description,
					CoverImageUrl:     out.serat.CoverImageUrl,
					ThumbnailImageUrl: out.serat.ThumbnailImageUrl,
				}
				m.db.On("CreateSerat", in.ctx, mock.AnythingOfType("*query.CreateSeratParams")).Return(row, out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratRepository{
				db: new(mockQuery.Querier),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := postgres.NewPostgresSeratRepository(m.db)
			out, err := subject.Create(tt.in.ctx, tt.in.iserat)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), tt.out.err.Error())
			}

			assert.Equal(t, tt.out.serat, out)
		})
	}
}

func TestSeratRepository_GetSerat(t *testing.T) {
	type input struct {
		ctx context.Context
		id  uuid.UUID
	}

	type output struct {
		serat *entity.Serat
		err   error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratRepository, *input, *output)
	}

	tests := []scenario{
		{
			name: "db.GetSerat return error",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				serat: nil,
				err:   fmt.Errorf("GetSerat: failed executing db query: Bababoey"),
			},
			on: func(m *MockSeratRepository, in *input, out *output) {
				m.db.On("GetSerat", in.ctx, in.id).Return(query.GetSeratRow{}, out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				serat: &entity.Serat{
					ID:                uuid.New(),
					Title:             "Jayabaya",
					Description:       "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
				err: nil,
			},
			on: func(m *MockSeratRepository, in *input, out *output) {
				row := query.GetSeratRow{
					ID:                out.serat.ID,
					Title:             out.serat.Title,
					Description:       out.serat.Description,
					CoverImageUrl:     out.serat.CoverImageUrl,
					ThumbnailImageUrl: out.serat.ThumbnailImageUrl,
				}
				m.db.On("GetSerat", in.ctx, in.id).Return(row, out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratRepository{
				db: new(mockQuery.Querier),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := postgres.NewPostgresSeratRepository(m.db)
			out, err := subject.Get(tt.in.ctx, tt.in.id)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), tt.out.err.Error())
			}

			assert.Equal(t, tt.out.serat, out)
		})
	}
}

func TestSeratRepository_DeleteSerat(t *testing.T) {
	type input struct {
		ctx context.Context
		id  uuid.UUID
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratRepository, *input, *output)
	}

	tests := []scenario{
		{
			name: "db.DeleteSerat return error",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				err: fmt.Errorf("DeleteSerat: failed executing db query: Bababoey"),
			},
			on: func(m *MockSeratRepository, in *input, out *output) {
				m.db.On("DeleteSerat", in.ctx, in.id).Return(out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockSeratRepository, in *input, out *output) {
				m.db.On("DeleteSerat", in.ctx, in.id).Return(out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratRepository{
				db: new(mockQuery.Querier),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := postgres.NewPostgresSeratRepository(m.db)
			err := subject.Delete(tt.in.ctx, tt.in.id)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), tt.out.err.Error())
			}
		})
	}
}

func TestSeratRepository_ListSerats(t *testing.T) {
	type input struct {
		ctx        context.Context
		pagination *entity.Pagination
	}

	type output struct {
		serats []entity.Serat
		count  uint32
		err    error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratRepository, *input, *output)
	}

	tests := []scenario{
		{
			name: "db.ListSerats return error",
			in: &input{
				ctx: context.Background(),
				pagination: &entity.Pagination{
					Page:     1,
					PageSize: 10,
				},
			},
			out: &output{
				count: uint32(0),
				err:   fmt.Errorf("ListSerats: failed executing db query: Bababoey"),
			},
			on: func(m *MockSeratRepository, in *input, out *output) {
				var rows []query.ListSeratsRow

				params := &query.ListSeratsParams{
					Limit:  int32(in.pagination.GetLimit()),
					Offset: int32(in.pagination.GetOffset()),
				}

				m.db.On("ListSerats", in.ctx, params).Return(rows, out.err)
			},
		},
		{
			name: "db.CountSerats return error",
			in: &input{
				ctx: context.Background(),
				pagination: &entity.Pagination{
					Page:     1,
					PageSize: 10,
				},
			},
			out: &output{
				serats: []entity.Serat{
					{
						ID:                uuid.New(),
						Title:             "Jayabaya",
						Description:       "Lorem ipsum dolor sit amet",
						CoverImageUrl:     "https://placeimg.com/640/480/any",
						ThumbnailImageUrl: "https://placeimg.com/640/480/any",
					},
				},
				count: uint32(0),
				err:   fmt.Errorf("CountSerats: failed executing db query: Bababoey"),
			},
			on: func(m *MockSeratRepository, in *input, out *output) {
				rows := []query.ListSeratsRow{
					{
						ID:                out.serats[0].ID,
						Title:             out.serats[0].Title,
						Description:       out.serats[0].Description,
						CoverImageUrl:     out.serats[0].CoverImageUrl,
						ThumbnailImageUrl: out.serats[0].ThumbnailImageUrl,
					},
				}

				params := &query.ListSeratsParams{
					Limit:  int32(in.pagination.GetLimit()),
					Offset: int32(in.pagination.GetOffset()),
				}

				c := int64(out.count)

				m.db.On("ListSerats", in.ctx, params).Return(rows, nil)
				m.db.On("CountSerats", in.ctx).Return(c, out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				pagination: &entity.Pagination{
					Page:     1,
					PageSize: 10,
				},
			},
			out: &output{
				serats: []entity.Serat{
					{
						ID:                uuid.New(),
						Title:             "Jayabaya",
						Description:       "Lorem ipsum dolor sit amet",
						CoverImageUrl:     "https://placeimg.com/640/480/any",
						ThumbnailImageUrl: "https://placeimg.com/640/480/any",
					},
				},
				count: uint32(1),
				err:   nil,
			},
			on: func(m *MockSeratRepository, in *input, out *output) {
				rows := []query.ListSeratsRow{
					{
						ID:                out.serats[0].ID,
						Title:             out.serats[0].Title,
						Description:       out.serats[0].Description,
						CoverImageUrl:     out.serats[0].CoverImageUrl,
						ThumbnailImageUrl: out.serats[0].ThumbnailImageUrl,
					},
				}

				params := &query.ListSeratsParams{
					Limit:  int32(in.pagination.GetLimit()),
					Offset: int32(in.pagination.GetOffset()),
				}

				c := int64(out.count)

				m.db.On("ListSerats", in.ctx, params).Return(rows, out.err)
				m.db.On("CountSerats", in.ctx).Return(c, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratRepository{
				db: new(mockQuery.Querier),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := postgres.NewPostgresSeratRepository(m.db)
			out, count, err := subject.List(tt.in.ctx, tt.in.pagination)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), tt.out.err.Error())
			}

			assert.Equal(t, tt.out.serats, out)
			assert.Equal(t, tt.out.count, count)
		})
	}
}

func TestSeratRepository_UpdateSerat(t *testing.T) {
	type input struct {
		ctx    context.Context
		id     uuid.UUID
		userat *entity.Serat
	}

	type output struct {
		serat *entity.Serat
		err   error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratRepository, *input, *output)
	}

	tests := []scenario{
		{
			name: "db.UpdateSerat return error",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
				userat: &entity.Serat{
					Title:             "Lorem Ipsum",
					Description:       "Lorem Ipsum",
					Content:           "Lorem Ipsum",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
			},
			out: &output{
				serat: nil,
				err:   fmt.Errorf("UpdateSerat: failed executing db query: Bababoey"),
			},
			on: func(m *MockSeratRepository, in *input, out *output) {
				m.db.On("UpdateSerat", in.ctx, mock.AnythingOfType("*query.UpdateSeratParams")).Return(query.UpdateSeratRow{}, out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				userat: &entity.Serat{
					Title:             "Lorem Ipsum",
					Description:       "Lorem Ipsum",
					Content:           "Lorem Ipsum",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
			},
			out: &output{
				serat: &entity.Serat{
					ID:                uuid.New(),
					Title:             "Jayabaya",
					Description:       "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
				err: nil,
			},
			on: func(m *MockSeratRepository, in *input, out *output) {
				row := query.UpdateSeratRow{
					ID:                out.serat.ID,
					Title:             out.serat.Title,
					Description:       out.serat.Description,
					CoverImageUrl:     out.serat.CoverImageUrl,
					ThumbnailImageUrl: out.serat.ThumbnailImageUrl,
				}
				m.db.On("UpdateSerat", in.ctx, mock.AnythingOfType("*query.UpdateSeratParams")).Return(row, out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratRepository{
				db: new(mockQuery.Querier),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := postgres.NewPostgresSeratRepository(m.db)
			out, err := subject.Update(tt.in.ctx, tt.in.id, tt.in.userat)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), tt.out.err.Error())
			}

			assert.Equal(t, tt.out.serat, out)
		})
	}
}
