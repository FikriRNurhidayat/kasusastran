package srv_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/app/srv"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mocks "github.com/fikrirnurhidayat/kasusastran/mocks/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/proto"
)

func TestSeratService_CreateSerat(t *testing.T) {
	type input struct {
		ctx context.Context
		req *proto.CreateSeratRequest
	}

	type output struct {
		res *proto.Serat
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
			name: "CreateSeratUseCase.Call return error",
			in: &input{
				ctx: context.Background(),
				req: &proto.CreateSeratRequest{},
			},
			out: &output{
				res: nil,
				err: fmt.Errorf("CreateSeratUseCase.Call: failed to run svc: bababoey"),
			},
			on: func(m *MockSeratsServer, in *input, out *output) {
				m.createSeratService.On("Call", in.ctx, mock.AnythingOfType("*svc.CreateSeratParams")).Return(nil, out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &proto.CreateSeratRequest{
					Title:             "Lorem ipsum",
					Description:       "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
			},
			out: &output{
				res: &proto.Serat{
					Id:                uuid.New().String(),
					Title:             "Lorem ipsum",
					Description:       "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
				err: nil,
			},
			on: func(m *MockSeratsServer, in *input, out *output) {
				m.createSeratService.On("Call", in.ctx, mock.AnythingOfType("*svc.CreateSeratParams")).Return(&svc.CreateSeratResult{
					ID:                uuid.MustParse(out.res.GetId()),
					Title:             out.res.GetTitle(),
					Description:       out.res.GetDescription(),
					CoverImageUrl:     out.res.GetCoverImageUrl(),
					ThumbnailImageUrl: out.res.GetThumbnailImageUrl(),
				}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratsServer{
				createSeratService: new(mocks.CreateSeratService),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := srv.NewSeratsServer(srv.WithCreateSeratService(m.createSeratService))
			out, err := subject.CreateSerat(tt.in.ctx, tt.in.req)

			if tt.out.err != nil {
				assert.NotNil(t, tt.out.err.Error(), err)
			}

			assert.Equal(t, tt.out.res, out)
		})
	}
}
