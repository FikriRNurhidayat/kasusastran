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

	mockEvent "github.com/fikrirnurhidayat/kasusastran/mocks/domain/event"
	mockRepo "github.com/fikrirnurhidayat/kasusastran/mocks/domain/repository"
)

type MockUpdateSeratService struct {
	seratEventEmitter *mockEvent.SeratEventEmitter
	seratRepository   *mockRepo.SeratRepository
}

func TestUpdateSeratService_Call(t *testing.T) {
	type input struct {
		ctx    context.Context
		id     uuid.UUID
		params *svc.UpdateSeratParams
	}

	type output struct {
		res *svc.UpdateSeratResult
		err error
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
				res: nil,
				err: fmt.Errorf("seratRepository.Update: failed to retrieve: baboey"),
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
				res: &svc.UpdateSeratResult{
					ID:                uuid.New(),
					Title:             "Jayabaya",
					Description:       "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
				err: nil,
			},
			on: func(m *MockUpdateSeratService, in *input, out *output) {
				serat := &entity.Serat{
					ID:                out.res.ID,
					Title:             out.res.Title,
					Description:       out.res.Description,
					CoverImageUrl:     out.res.CoverImageUrl,
					ThumbnailImageUrl: out.res.ThumbnailImageUrl,
				}

				m.seratRepository.On("Update", in.ctx, in.id, mock.AnythingOfType("*entity.Serat")).Return(serat, out.err)
				m.seratEventEmitter.On("EmitUpdatedEvent", mock.AnythingOfType("*message.Serat")).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockUpdateSeratService{
				seratRepository:   new(mockRepo.SeratRepository),
				seratEventEmitter: new(mockEvent.SeratEventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := svc.NewUpdateSeratService(m.seratRepository, m.seratEventEmitter)
			out, err := subject.Call(tt.in.ctx, tt.in.id, tt.in.params)

			if err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}

			assert.Equal(t, tt.out.res, out)
		})
	}
}
