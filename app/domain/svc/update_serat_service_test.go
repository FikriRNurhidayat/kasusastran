package svc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mocks "github.com/fikrirnurhidayat/kasusastran/mocks/domain/repository"
)

type MockUpdateSeratService struct {
	seratRepository *mocks.SeratRepository
}

func TestUpdateSeratService_Call(t *testing.T) {
	type input struct {
		ctx    context.Context
		id     uuid.UUID
		params *svc.UpdateSeratParams
	}

	type output struct {
		serat *entity.Serat
		err   error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockUpdateSeratService, *input, *output)
	}

	tests := []scenario{
		{
			name: "seratRepository.Update: return error",
			in: &input{
				ctx:    context.Background(),
				id:     uuid.New(),
				params: &svc.UpdateSeratParams{},
			},
			out: &output{
				serat: nil,
				err:   fmt.Errorf("seratRepository.Update: failed to retrieve: baboey"),
			},
			on: func(m *MockUpdateSeratService, in *input, out *output) {
				m.seratRepository.On("Update", in.ctx, in.id, mock.AnythingOfType("*entity.Serat")).Return(nil, out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				id:     uuid.New(),
				params: &svc.UpdateSeratParams{},
			},
			out: &output{
				serat: &entity.Serat{
					ID:                uuid.New(),
					Title:             "Jayabaya",
					Description:       "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
				err: nil,
			},
			on: func(m *MockUpdateSeratService, in *input, out *output) {
				m.seratRepository.On("Update", in.ctx, in.id, mock.AnythingOfType("*entity.Serat")).Return(out.serat, out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockUpdateSeratService{
				seratRepository: new(mocks.SeratRepository),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := svc.NewUpdateSeratService(m.seratRepository)
			out, err := subject.Call(tt.in.ctx, tt.in.id, tt.in.params)

			if err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}

			assert.Equal(t, tt.out.serat, out)
		})
	}
}
