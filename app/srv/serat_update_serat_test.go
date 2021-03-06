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

func TestSeratsServer_UpdateSerat(t *testing.T) {
	type input struct {
		ctx context.Context
		req *proto.UpdateSeratRequest
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
			name: "uuid.Parse return error",
			in: &input{
				ctx: context.Background(),
				req: &proto.UpdateSeratRequest{
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
			name: "UpdateSeratUseCase.Call return error",
			in: &input{
				ctx: context.Background(),
				req: &proto.UpdateSeratRequest{},
			},
			out: &output{
				res: nil,
				err: fmt.Errorf("UpdateSeratUseCase.Call: failed to run svc: bababoey"),
			},
			on: func(m *MockSeratsServer, in *input, out *output) {
				m.updateSeratService.On("Call", in.ctx, mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("*svc.UpdateSeratParams")).Return(nil, out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &proto.UpdateSeratRequest{
					Id:                "f6834cdc-93a3-4d60-b975-d42e7aa26b81",
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
				m.updateSeratService.On("Call", in.ctx, mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("*svc.UpdateSeratParams")).Return(&svc.UpdateSeratResult{
					ID:                uuid.MustParse(out.res.GetId()),
					Title:             out.res.GetTitle(),
					Description:       out.res.GetDescription(),
					CoverImageUrl:     out.res.GetCoverImageUrl(),
					ThumbnailImageUrl: out.res.GetThumbnailImageUrl(),
				}, out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratsServer{
				updateSeratService: new(mocks.UpdateSeratService),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := srv.NewSeratsServer(srv.WithUpdateSeratService(m.updateSeratService))
			out, err := subject.UpdateSerat(tt.in.ctx, tt.in.req)

			if tt.out.err != nil {
				assert.NotNil(t, tt.out.err.Error(), err)
			}

			assert.Equal(t, tt.out.res, out)
		})
	}
}
