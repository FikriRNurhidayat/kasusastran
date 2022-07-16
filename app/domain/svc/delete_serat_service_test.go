package svc_test

import (
	"context"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/app/trouble"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockEvent "github.com/fikrirnurhidayat/kasusastran/mocks/domain/event"
	mockRepo "github.com/fikrirnurhidayat/kasusastran/mocks/domain/repository"
)

type MockDeleteSeratService struct {
	seratEventEmitter *mockEvent.SeratEventEmitter
	seratRepository   *mockRepo.SeratRepository
}

func TestDeleteSeratService_Call(t *testing.T) {
	type input struct {
		ctx context.Context
		id  uuid.UUID
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockDeleteSeratService, *input, *output)
	}

	tests := []scenario{
		{
			name: "s.seratRepository.Get return error",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				err: trouble.SERAT_NOT_FOUND,
			},
			on: func(m *MockDeleteSeratService, in *input, out *output) {
				m.seratRepository.On("Get", in.ctx, in.id).Return(nil, out.err)
			},
		},
		{
			name: "s.seratRepository.Delete return error",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				err: trouble.INTERNAL_SERVER_ERROR,
			},
			on: func(m *MockDeleteSeratService, in *input, out *output) {
				serat := &entity.Serat{
					ID:                uuid.New(),
					Title:             "Industrial Society And Its Future",
					Description:       "Lorem ipsum dolor sit amet",
					Content:           "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "http://placekitten.com/192/108",
					ThumbnailImageUrl: "http://placekitten.com/192/108",
				}
				m.seratRepository.On("Get", in.ctx, in.id).Return(serat, nil)
				m.seratRepository.On("Delete", in.ctx, in.id).Return(out.err)
			},
		},
		{
			name: "s.seratEventEmitter.EmitDeletedEvent return error",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				err: trouble.INTERNAL_SERVER_ERROR,
			},
			on: func(m *MockDeleteSeratService, in *input, out *output) {
				serat := &entity.Serat{
					ID:                uuid.New(),
					Title:             "Industrial Society And Its Future",
					Description:       "Lorem ipsum dolor sit amet",
					Content:           "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "http://placekitten.com/192/108",
					ThumbnailImageUrl: "http://placekitten.com/192/108",
				}
				m.seratRepository.On("Get", in.ctx, in.id).Return(serat, nil)
				m.seratRepository.On("Delete", in.ctx, in.id).Return(nil)
				m.seratEventEmitter.On("EmitDeletedEvent", mock.AnythingOfType("*message.Serat")).Return(out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockDeleteSeratService, in *input, out *output) {
				serat := &entity.Serat{
					ID:                uuid.New(),
					Title:             "Industrial Society And Its Future",
					Description:       "Lorem ipsum dolor sit amet",
					Content:           "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "http://placekitten.com/192/108",
					ThumbnailImageUrl: "http://placekitten.com/192/108",
				}
				m.seratRepository.On("Get", in.ctx, in.id).Return(serat, nil)
				m.seratRepository.On("Delete", in.ctx, in.id).Return(out.err)
				m.seratEventEmitter.On("EmitDeletedEvent", mock.AnythingOfType("*message.Serat")).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDeleteSeratService{
				seratRepository:   new(mockRepo.SeratRepository),
				seratEventEmitter: new(mockEvent.SeratEventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := svc.NewDeleteSeratService(m.seratRepository, m.seratEventEmitter)
			err := subject.Call(tt.in.ctx, tt.in.id)

			if err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}
		})
	}
}
