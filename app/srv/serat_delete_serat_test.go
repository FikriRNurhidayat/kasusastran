package srv_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/srv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/emptypb"

	api "github.com/fikrirnurhidayat/kasusastran/api"
	mocks "github.com/fikrirnurhidayat/kasusastran/mocks/domain/svc"
)

func TestSeratService_DeleteSerat(t *testing.T) {
	type input struct {
		ctx context.Context
		req *api.DeleteSeratRequest
	}

	type output struct {
		res *emptypb.Empty
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
			name: "uuid.Parse return error",
			in: &input{
				ctx: context.Background(),
				req: &api.DeleteSeratRequest{
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
			name: "DeleteSeratUseCase.Exec return error",
			in: &input{
				ctx: context.Background(),
				req: &api.DeleteSeratRequest{
					Id: "f6834cdc-93a3-4d60-b975-d42e7aa26b81",
				},
			},
			out: &output{
				res: nil,
				err: fmt.Errorf("DeleteSeratUseCase.Exec: failed to run svc: bababoey"),
			},
			on: func(m *MockSeratService, in *input, out *output) {
				m.deleteSeratService.On("Exec", in.ctx, mock.AnythingOfType("uuid.UUID")).Return(out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &api.DeleteSeratRequest{
					Id: "f6834cdc-93a3-4d60-b975-d42e7aa26b81",
				},
			},
			out: &output{
				res: &emptypb.Empty{},
				err: nil,
			},
			on: func(m *MockSeratService, in *input, out *output) {
				m.deleteSeratService.On("Exec", in.ctx, mock.AnythingOfType("uuid.UUID")).Return(out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratService{
				deleteSeratService: new(mocks.DeleteSeratService),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := srv.NewSeratsServer().SetDeleteSeratUseCase(m.deleteSeratService)
			out, err := subject.DeleteSerat(tt.in.ctx, tt.in.req)

			if tt.out.err != nil {
				assert.NotNil(t, tt.out.err.Error(), err)
			}

			assert.Equal(t, tt.out.res, out)
		})
	}
}
