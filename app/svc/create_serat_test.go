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

func TestSeratService_CreateSerat(t *testing.T) {
	type input struct {
		ctx context.Context
		req *api.CreateSeratRequest
	}

	type output struct {
		res *api.Serat
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
			name: "CreateSeratUseCase.Exec return error",
			in: &input{
				ctx: context.Background(),
				req: &api.CreateSeratRequest{},
			},
			out: &output{
				res: nil,
				err: fmt.Errorf("CreateSeratUseCase.Exec: failed to run usecase: bababoey"),
			},
			on: func(m *MockSeratService, in *input, out *output) {
				m.CreateSeratUseCase.On("Exec", in.ctx, mock.AnythingOfType("*usecase.CreateSeratParams")).Return(nil, out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &api.CreateSeratRequest{},
			},
			out: &output{
				res: &api.Serat{
					Id:                uuid.New().String(),
					Title:             "Lorem ipsum",
					Description:       "Lorem ipsum dolor sit amet",
					Content:           "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
				err: nil,
			},
			on: func(m *MockSeratService, in *input, out *output) {
				m.CreateSeratUseCase.On("Exec", in.ctx, mock.AnythingOfType("*usecase.CreateSeratParams")).Return(&entity.Serat{
					ID:                uuid.MustParse(out.res.GetId()),
					Title:             out.res.GetTitle(),
					Description:       out.res.GetDescription(),
					Content:           out.res.GetContent(),
					CoverImageUrl:     out.res.GetCoverImageUrl(),
					ThumbnailImageUrl: out.res.GetThumbnailImageUrl(),
				}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratService{
				CreateSeratUseCase: new(mockUseCase.CreateSeratUseCaser),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := svc.NewSeratsService().SetCreateSeratUseCase(m.CreateSeratUseCase)
			out, err := subject.CreateSerat(tt.in.ctx, tt.in.req)

			if tt.out.err != nil {
				assert.NotNil(t, tt.out.err.Error(), err)
			}

			assert.Equal(t, tt.out.res, out)
		})
	}
}
