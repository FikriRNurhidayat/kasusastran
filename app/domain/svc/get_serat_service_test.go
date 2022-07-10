package svc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	mocks "github.com/fikrirnurhidayat/kasusastran/mocks/domain/repository"
)

type MockGetSeratService struct {
	seratRepository *mocks.SeratRepository
}

func TestGetSeratUseCase_Call(t *testing.T) {
	type input struct {
		ctx context.Context
		id  uuid.UUID
	}

	type output struct {
		serat *svc.GetSeratResult
		err   error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockGetSeratService, *input, *output)
	}

	tests := []scenario{
		{
			name: "SeratRepository.GetSerat return error",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				serat: nil,
				err:   fmt.Errorf("seratRepository.Get: failed to retrieve: baboey"),
			},
			on: func(m *MockGetSeratService, in *input, out *output) {
				m.seratRepository.On("Get", in.ctx, in.id).Return(nil, out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				serat: &svc.GetSeratResult{
					ID:                uuid.New(),
					Title:             "Jayabaya",
					Description:       "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
				err: nil,
			},
			on: func(m *MockGetSeratService, in *input, out *output) {
				serat := &entity.Serat{
					ID:                out.serat.ID,
					Title:             out.serat.Title,
					Description:       out.serat.Description,
					CoverImageUrl:     out.serat.CoverImageUrl,
					ThumbnailImageUrl: out.serat.ThumbnailImageUrl,
				}

				m.seratRepository.On("Get", in.ctx, in.id).Return(serat, out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockGetSeratService{
				seratRepository: new(mocks.SeratRepository),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := svc.NewGetSeratService(m.seratRepository)
			out, err := subject.Call(tt.in.ctx, tt.in.id)

			if err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}

			assert.Equal(t, tt.out.serat, out)
		})
	}
}
