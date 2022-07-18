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

type MockCreateSeratService struct {
	seratRepository   *mockRepo.SeratRepository
	seratEventEmitter *mockEvent.SeratEventEmitter
}

func TestCreateSeratService_Call(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *svc.CreateSeratParams
	}

	type output struct {
		serat *svc.CreateSeratResult
		err   error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockCreateSeratService, *input, *output)
	}

	tests := []scenario{
		{
			name: "s.seratRepository.Create return error",
			in: &input{
				ctx:    context.Background(),
				params: &svc.CreateSeratParams{},
			},
			out: &output{
				serat: nil,
				err:   trouble.INTERNAL_SERVER_ERROR,
			},
			on: func(m *MockCreateSeratService, in *input, out *output) {
				m.seratRepository.On("Create", in.ctx, mock.AnythingOfType("*entity.Serat")).Return(nil, out.err)
			},
		},
		{
			name: "s.seratEventEmitter.EmitCreatedEvent return error",
			in: &input{
				ctx:    context.Background(),
				params: &svc.CreateSeratParams{},
			},
			out: &output{
				serat: nil,
				err:   trouble.INTERNAL_SERVER_ERROR,
			},
			on: func(m *MockCreateSeratService, in *input, out *output) {
				serat := &entity.Serat{
					ID:                uuid.New(),
					Title:             "Lorem ipsum",
					Description:       "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "http://placekitten.com/192/168",
					ThumbnailImageUrl: "http://placekitten.com/192/168",
				}

				m.seratRepository.On("Create", in.ctx, mock.AnythingOfType("*entity.Serat")).Return(serat, nil)
				m.seratEventEmitter.On("EmitCreatedEvent", mock.AnythingOfType("*message.Serat")).Return(out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: &svc.CreateSeratParams{},
			},
			out: &output{
				serat: &svc.CreateSeratResult{
					ID:                uuid.New(),
					Title:             "Jayabaya",
					Description:       "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
				err: nil,
			},
			on: func(m *MockCreateSeratService, in *input, out *output) {
				serat := &entity.Serat{
					ID:                out.serat.ID,
					Title:             out.serat.Title,
					Description:       out.serat.Description,
					CoverImageUrl:     out.serat.CoverImageUrl,
					ThumbnailImageUrl: out.serat.ThumbnailImageUrl,
				}

				m.seratRepository.On("Create", in.ctx, mock.AnythingOfType("*entity.Serat")).Return(serat, out.err)
				m.seratEventEmitter.On("EmitCreatedEvent", mock.AnythingOfType("*message.Serat")).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockCreateSeratService{
				seratRepository:   new(mockRepo.SeratRepository),
				seratEventEmitter: new(mockEvent.SeratEventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := svc.NewCreateSeratService(m.seratRepository, m.seratEventEmitter)
			out, err := subject.Call(tt.in.ctx, tt.in.params)

			if err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}

			assert.Equal(t, tt.out.serat, out)
		})
	}
}
