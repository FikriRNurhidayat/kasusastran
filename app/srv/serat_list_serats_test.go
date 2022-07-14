package srv_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/app/srv"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockManager "github.com/fikrirnurhidayat/kasusastran/mocks/domain/manager"
	mockService "github.com/fikrirnurhidayat/kasusastran/mocks/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/proto"
)

func TestSeratService_ListSerats(t *testing.T) {
	type input struct {
		ctx context.Context
		req *proto.ListSeratsRequest
	}

	type output struct {
		res *proto.ListSeratsResponse
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratsServer, *input, *output)
	}

	tests := []scenario{
		{
			name: "ListSeratsUseCase.Call return error",
			in: &input{
				ctx: context.Background(),
				req: &proto.ListSeratsRequest{},
			},
			out: &output{
				err: fmt.Errorf("ListSeratsUseCase.Call: failed to execute svc"),
			},
			on: func(m *MockSeratsServer, i *input, o *output) {
				m.paginationManager.On("FromIncomingRequest", i.req.GetPagination()).Return(&manager.Pagination{
					Page:     1,
					PageSize: 10,
				})
				m.listSeratsService.On("Call", i.ctx, mock.AnythingOfType("*svc.ListSeratsParams")).Return(nil, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &proto.ListSeratsRequest{},
			},
			out: &output{
				res: &proto.ListSeratsResponse{
					Meta: &proto.ListSeratsResponse_MetaListSerats{
						Pagination: &proto.ResponsePagination{
							Page:      1,
							PageSize:  10,
							PageCount: 1,
						},
					},
					Serats: []*proto.Serat{
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
			on: func(m *MockSeratsServer, i *input, o *output) {
				m.paginationManager.On("FromIncomingRequest", i.req.GetPagination()).Return(&manager.Pagination{
					Page:     1,
					PageSize: 10,
				})

				serats := []*entity.Serat{}

				for _, serat := range o.res.Serats {
					pack := &entity.Serat{
						ID:                uuid.MustParse(serat.GetId()),
						Title:             serat.GetTitle(),
						Description:       serat.GetDescription(),
						CoverImageUrl:     serat.GetCoverImageUrl(),
						ThumbnailImageUrl: serat.GetThumbnailImageUrl(),
					}

					serats = append(serats, pack)
				}

				result := &svc.ListSeratsResult{
					Serats: serats,
					Pagination: &manager.Pagination{
						Page:      o.res.Meta.Pagination.Page,
						PageSize:  o.res.Meta.Pagination.PageSize,
						PageCount: o.res.Meta.Pagination.PageCount,
						Total:     o.res.Meta.Pagination.Total,
					},
				}

				m.listSeratsService.On("Call", i.ctx, mock.AnythingOfType("*svc.ListSeratsParams")).Return(result, o.err)
				m.paginationManager.On("NewOutgoingResponse", mock.AnythingOfType("*manager.Pagination")).Return(o.res.GetMeta().GetPagination())
			},
		},
		{
			name: "OK With Pagination",
			in: &input{
				ctx: context.Background(),
				req: &proto.ListSeratsRequest{
					Pagination: &proto.RequestPagination{
						Page:     2,
						PageSize: 10,
					},
				},
			},
			out: &output{
				res: &proto.ListSeratsResponse{
					Meta: &proto.ListSeratsResponse_MetaListSerats{
						Pagination: &proto.ResponsePagination{
							Page:      2,
							PageSize:  10,
							PageCount: 1,
						},
					},
					Serats: []*proto.Serat{
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
			on: func(m *MockSeratsServer, i *input, o *output) {
				m.paginationManager.On("FromIncomingRequest", i.req.GetPagination()).Return(&manager.Pagination{
					Page:     o.res.Meta.Pagination.Page,
					PageSize: o.res.Meta.Pagination.PageSize,
				})

				serats := []*entity.Serat{}

				for _, serat := range o.res.Serats {
					pack := &entity.Serat{
						ID:                uuid.MustParse(serat.GetId()),
						Title:             serat.GetTitle(),
						Description:       serat.GetDescription(),
						CoverImageUrl:     serat.GetCoverImageUrl(),
						ThumbnailImageUrl: serat.GetThumbnailImageUrl(),
					}

					serats = append(serats, pack)
				}

				result := &svc.ListSeratsResult{
					Serats: serats,
					Pagination: &manager.Pagination{
						Page:      o.res.Meta.Pagination.Page,
						PageSize:  o.res.Meta.Pagination.PageSize,
						PageCount: o.res.Meta.Pagination.PageCount,
						Total:     o.res.Meta.Pagination.Total,
					},
				}

				m.listSeratsService.On("Call", i.ctx, mock.AnythingOfType("*svc.ListSeratsParams")).Return(result, o.err)
				m.paginationManager.On("NewOutgoingResponse", mock.AnythingOfType("*manager.Pagination")).Return(o.res.Meta.GetPagination())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratsServer{
				listSeratsService: new(mockService.ListSeratsService),
				paginationManager: new(mockManager.PaginationManager),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := srv.NewSeratsServer(
				srv.WithListSeratsService(m.listSeratsService),
				srv.WithPaginationManager(m.paginationManager),
			)
			res, err := subject.ListSerats(tt.in.ctx, tt.in.req)

			if tt.out.err != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, tt.out.res, res)
		})
	}
}
