package srv_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/api"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/app/srv"
	"github.com/fikrirnurhidayat/kasusastran/mocks/domain/svc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestWulangansServer_GetWulangan(t *testing.T) {
	type input struct {
		ctx context.Context
		req *api.GetWulanganRequest
	}

	type output struct {
		res *api.Wulangan
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockWulangansServer, *input, *output)
	}

	tests := []scenario{
		{
			name: "Invalid Request",
			in: &input{
				ctx: context.Background(),
				req: &api.GetWulanganRequest{
					Id: "I'm not an uuid",
				},
			},
			out: &output{
				res: nil,
				err: status.Errorf(codes.InvalidArgument, "validation error"),
			},
		},
		{
			name: "s.getWulanganService.Call() return error",
			in: &input{
				ctx: context.Background(),
				req: &api.GetWulanganRequest{
					Id: uuid.New().String(),
				},
			},
			out: &output{
				res: nil,
				err: status.Errorf(codes.NotFound, "wulangan not found"),
			},
			on: func(mws *MockWulangansServer, i *input, o *output) {
				id := uuid.MustParse(i.req.GetId())
				mws.getWulanganService.On("Call", i.ctx, id).Return(nil, errors.New("wulangan not found"))
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &api.GetWulanganRequest{
					Id: uuid.New().String(),
				},
			},
			out: &output{
				err: nil,
			},
			on: func(mws *MockWulangansServer, i *input, o *output) {
				id := uuid.MustParse(i.req.GetId())

				result := &svc.GetWulanganResult{
					ID:                id,
					Title:             "Industrial Society And Its Future",
					Description:       "Had been a disaster of human race.",
					CoverImageUrl:     "https://placekitten.com/192/108",
					ThumbnailImageUrl: "https://placekitten.com/192/108",
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				}

				o.res = &api.Wulangan{
					Id:                result.ID.String(),
					Title:             result.Title,
					Description:       result.Description,
					CoverImageUrl:     result.CoverImageUrl,
					ThumbnailImageUrl: result.ThumbnailImageUrl,
					CreatedAt:         timestamppb.New(result.CreatedAt),
					UpdatedAt:         timestamppb.New(result.UpdatedAt),
				}

				mws.getWulanganService.On("Call", i.ctx, id).Return(result, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockWulangansServer{
				getWulanganService: new(mocks.GetWulanganService),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := srv.NewWulangansServer(srv.WithGetWulanganService(m.getWulanganService))
			got, err := subject.GetWulangan(tt.in.ctx, tt.in.req)

			if tt.out.err != nil {
				assert.NotNil(t, err)
			}

			if tt.out.res != nil {
				assert.NotNil(t, got)
				assert.Equal(t, tt.out.res, got)
			}
		})
	}
}
