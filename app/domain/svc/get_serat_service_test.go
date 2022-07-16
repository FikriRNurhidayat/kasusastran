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

type MockGetSeratService struct {
	seratRepository   *mockRepo.SeratRepository
	seratEventEmitter *mockEvent.SeratEventEmitter
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
			name: "s.seratRepository.Get return error",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				serat: nil,
				err:   trouble.SERAT_NOT_FOUND,
			},
			on: func(m *MockGetSeratService, in *input, out *output) {
				m.seratRepository.On("Get", in.ctx, in.id).Return(nil, out.err)
			},
		},
		{
			name: "s.seratEventEmitter.EmitRetrievedEvent return error",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				serat: nil,
				err:   trouble.INTERNAL_SERVER_ERROR,
			},
			on: func(m *MockGetSeratService, in *input, out *output) {
				serat := &entity.Serat{
					ID:                uuid.New(),
					Title:             "Jayabaya",
					Description:       "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				}

				m.seratRepository.On("Get", in.ctx, in.id).Return(serat, nil)
				m.seratEventEmitter.On("EmitRetrievedEvent", mock.AnythingOfType("*message.Serat")).Return(out.err)
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
				m.seratEventEmitter.On("EmitRetrievedEvent", mock.AnythingOfType("*message.Serat")).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockGetSeratService{
				seratRepository:   new(mockRepo.SeratRepository),
				seratEventEmitter: new(mockEvent.SeratEventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := svc.NewGetSeratService(m.seratRepository, m.seratEventEmitter)
			out, err := subject.Call(tt.in.ctx, tt.in.id)

			if err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}

			assert.Equal(t, tt.out.serat, out)
		})
	}
}
