package svc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/svc/svc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	api "github.com/fikrirnurhidayat/kasusastran/api"
	mockUseCase "github.com/fikrirnurhidayat/kasusastran/mocks/domain/usecase"
)

func TestSeratService_ListSerats(t *testing.T) {
	type input struct {
		ctx context.Context
		req *api.ListSeratsRequest
	}

	type output struct {
		res *api.ListSeratsResponse
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratService, *input, *output)
	}

	tests := []scenario{
		{
			name: "ListSeratsUseCase.Exec return error",
			in: &input{
				ctx: context.Background(),
				req: &api.ListSeratsRequest{},
			},
			out: &output{
				err: fmt.Errorf("ListSeratsUseCase.Exec: failed to execute usecase"),
			},
			on: func(m *MockSeratService, i *input, o *output) {
				svc := []entity.Serat{}
				pagination := &entity.Pagination{}
				m.ListSeratsUseCase.On("Exec", i.ctx, mock.AnythingOfType("*entity.Pagination")).Return(svc, pagination, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &api.ListSeratsRequest{},
			},
			out: &output{
				res: &api.ListSeratsResponse{
					Meta: &api.ListSeratsResponse_MetaListSerats{
						Pagination: &api.ResponsePagination{
							Page:      1,
							PageSize:  10,
							PageCount: 1,
						},
					},
					Serats: []*api.Serat{
						{
							Id:                "f6834cdc-93a3-4d60-b975-d42e7aa26b81",
							Title:             "Lorem ipsum dolor sit amet",
							Description:       "Lorem ipsum dolor sit amet",
							CoverImageUrl:     "https://placeimg.com/640/480/any",
							ThumbnailImageUrl: "https://placeimg.com/640/480/any",
						},
					},
				},
				err: nil,
			},
			on: func(m *MockSeratService, i *input, o *output) {
				svc := []entity.Serat{}

				pagination := &entity.Pagination{
					Page:      o.res.Meta.Pagination.Page,
					PageSize:  o.res.Meta.Pagination.PageSize,
					PageCount: o.res.Meta.Pagination.PageCount,
				}

				for _, serat := range o.res.Serats {
					pack := entity.Serat{
						ID:                uuid.MustParse(serat.GetId()),
						Title:             serat.GetTitle(),
						Description:       serat.GetDescription(),
						CoverImageUrl:     serat.GetCoverImageUrl(),
						ThumbnailImageUrl: serat.GetThumbnailImageUrl(),
					}

					svc = append(svc, pack)
				}

				m.ListSeratsUseCase.On("Exec", i.ctx, mock.AnythingOfType("*entity.Pagination")).Return(svc, pagination, o.err)
			},
		},
		{
			name: "OK With Pagination",
			in: &input{
				ctx: context.Background(),
				req: &api.ListSeratsRequest{
					Pagination: &api.RequestPagination{
						Page:     2,
						PageSize: 10,
					},
				},
			},
			out: &output{
				res: &api.ListSeratsResponse{
					Meta: &api.ListSeratsResponse_MetaListSerats{
						Pagination: &api.ResponsePagination{
							Page:      2,
							PageSize:  10,
							PageCount: 1,
						},
					},
					Serats: []*api.Serat{
						{
							Id:                "f6834cdc-93a3-4d60-b975-d42e7aa26b81",
							Title:             "Lorem ipsum dolor sit amet",
							Description:       "Lorem ipsum dolor sit amet",
							CoverImageUrl:     "https://placeimg.com/640/480/any",
							ThumbnailImageUrl: "https://placeimg.com/640/480/any",
						},
					},
				},
				err: nil,
			},
			on: func(m *MockSeratService, i *input, o *output) {
				svc := []entity.Serat{}

				pagination := &entity.Pagination{
					Page:      o.res.Meta.Pagination.Page,
					PageSize:  o.res.Meta.Pagination.PageSize,
					PageCount: o.res.Meta.Pagination.PageCount,
				}

				for _, serat := range o.res.Serats {
					pack := entity.Serat{
						ID:                uuid.MustParse(serat.GetId()),
						Title:             serat.GetTitle(),
						Description:       serat.GetDescription(),
						CoverImageUrl:     serat.GetCoverImageUrl(),
						ThumbnailImageUrl: serat.GetThumbnailImageUrl(),
					}

					svc = append(svc, pack)
				}

				m.ListSeratsUseCase.On("Exec", i.ctx, mock.AnythingOfType("*entity.Pagination")).Return(svc, pagination, o.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratService{
				ListSeratsUseCase: new(mockUseCase.ListSeratsUseCaser),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := svc.
				New().
				SetListSeratsUseCase(m.ListSeratsUseCase)

			res, err := subject.ListSerats(tt.in.ctx, tt.in.req)

			if tt.out.err != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, tt.out.res, res)
		})
	}
}
