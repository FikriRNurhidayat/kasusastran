package srv_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/srv"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	api "github.com/fikrirnurhidayat/kasusastran/api"
	mocks "github.com/fikrirnurhidayat/kasusastran/mocks/domain/svc"
)

func TestSeratService_GetSerat(t *testing.T) {
	type input struct {
		ctx context.Context
		req *api.GetSeratRequest
	}

	type output struct {
		res *api.Serat
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
			name: "uuid.Parse return error",
			in: &input{
				ctx: context.Background(),
				req: &api.GetSeratRequest{
					Id: "this-is-not-uuid",
				},
			},
			out: &output{
				res: nil,
				err: fmt.Errorf("uuid.Parse: failed to parse uuid: this-is-not-uuid"),
			},
			on: nil,
		},
		{
			name: "GetSeratUseCase.Call return error",
			in: &input{
				ctx: context.Background(),
				req: &api.GetSeratRequest{
					Id: "f6834cdc-93a3-4d60-b975-d42e7aa26b81",
				},
			},
			out: &output{
				res: nil,
				err: fmt.Errorf("GetSeratUseCase.Call: failed to run svc: bababoey"),
			},
			on: func(m *MockSeratsServer, in *input, out *output) {
				m.getSeratService.On("Call", in.ctx, mock.AnythingOfType("uuid.UUID")).Return(nil, out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &api.GetSeratRequest{
					Id: "f6834cdc-93a3-4d60-b975-d42e7aa26b81",
				},
			},
			out: &output{
				res: &api.Serat{
					Id:                uuid.New().String(),
					Title:             "Lorem ipsum",
					Description:       "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
				err: nil,
			},
			on: func(m *MockSeratsServer, in *input, out *output) {
				m.getSeratService.On("Call", in.ctx, mock.AnythingOfType("uuid.UUID")).Return(&entity.Serat{
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
				getSeratService: new(mocks.GetSeratService),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := srv.NewSeratsServer(srv.WithGetSeratService(m.getSeratService))
			out, err := subject.GetSerat(tt.in.ctx, tt.in.req)

			if tt.out.err != nil {
				assert.NotNil(t, tt.out.err.Error(), err)
			}

			assert.Equal(t, tt.out.res, out)
		})
	}
}
