package srv_test

import (
	"context"
	"errors"
	"testing"
	"time"

	mocks "github.com/fikrirnurhidayat/kasusastran/mocks/domain/svc"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/app/srv"
	"github.com/fikrirnurhidayat/kasusastran/proto"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestWulangansServer_CreateWulangan(t *testing.T) {
	type input struct {
		ctx context.Context
		req *proto.CreateWulanganRequest
	}

	type output struct {
		res *proto.Wulangan
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
			name: "req.ValidateAll() return error",
			in: &input{
				ctx: context.Background(),
				req: &proto.CreateWulanganRequest{
					Title:             "",
					Description:       "",
					CoverImageUrl:     "nonuri",
					ThumbnailImageUrl: "nonuri",
				},
			},
			out: &output{
				res: nil,
				err: status.Errorf(codes.OutOfRange, "validation error"),
			},
		},
		{
			name: "s.createWulanganService.Call() return error",
			in: &input{
				ctx: context.Background(),
				req: &proto.CreateWulanganRequest{
					Title:             "Technological Society",
					Description:       "Our society is degrading.",
					CoverImageUrl:     "https://source.unsplash.com/617x598",
					ThumbnailImageUrl: "https://source.unsplash.com/617x598",
				},
			},
			out: &output{
				res: nil,
				err: errors.New("s.createWulanganService.Call: failed to call createWulanganService"),
			},
			on: func(mws *MockWulangansServer, i *input, o *output) {
				params := &svc.CreateWulanganParams{
					Title:             i.req.GetTitle(),
					Description:       i.req.GetDescription(),
					CoverImageUrl:     i.req.GetCoverImageUrl(),
					ThumbnailImageUrl: i.req.GetThumbnailImageUrl(),
				}
				mws.createWulanganService.On("Call", i.ctx, params).Return(nil, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &proto.CreateWulanganRequest{
					Title:             "Technological Society",
					Description:       "Our society is degrading.",
					CoverImageUrl:     "https://source.unsplash.com/617x598",
					ThumbnailImageUrl: "https://source.unsplash.com/617x598",
				},
			},
			out: &output{
				err: nil,
			},
			on: func(mws *MockWulangansServer, i *input, o *output) {
				params := &svc.CreateWulanganParams{
					Title:             i.req.GetTitle(),
					Description:       i.req.GetDescription(),
					CoverImageUrl:     i.req.GetCoverImageUrl(),
					ThumbnailImageUrl: i.req.GetThumbnailImageUrl(),
				}

				w := &svc.CreateWulanganResult{
					ID:                uuid.New(),
					Title:             params.Title,
					Description:       params.Description,
					CoverImageUrl:     params.CoverImageUrl,
					ThumbnailImageUrl: params.ThumbnailImageUrl,
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				}

				o.res = &proto.Wulangan{
					Id:                w.ID.String(),
					Title:             w.Title,
					Description:       w.Description,
					CoverImageUrl:     w.CoverImageUrl,
					ThumbnailImageUrl: w.ThumbnailImageUrl,
					CreatedAt:         timestamppb.New(w.CreatedAt),
					UpdatedAt:         timestamppb.New(w.UpdatedAt),
				}

				mws.createWulanganService.On("Call", i.ctx, params).Return(w, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockWulangansServer{
				createWulanganService: new(mocks.CreateWulanganService),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := srv.NewWulangansServer(srv.WithCreateWulanganService(m.createWulanganService))
			res, err := subject.CreateWulangan(tt.in.ctx, tt.in.req)

			if tt.out.err != nil {
				assert.NotNil(t, err)
			}

			if tt.out.res != nil {
				assert.NotNil(t, res)
				assert.Equal(t, tt.out.res.GetId(), res.GetId())
				assert.Equal(t, tt.out.res.GetTitle(), res.GetTitle())
				assert.Equal(t, tt.out.res.GetDescription(), res.GetDescription())
				assert.Equal(t, tt.out.res.GetCoverImageUrl(), res.GetCoverImageUrl())
				assert.Equal(t, tt.out.res.GetThumbnailImageUrl(), res.GetThumbnailImageUrl())
			}
		})
	}
}
