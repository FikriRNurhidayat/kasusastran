package svc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
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
			name: "seratRepository.Delete return error",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				err: fmt.Errorf("seratRepository.Delete failed to retrieve: baboey"),
			},
			on: func(m *MockDeleteSeratService, in *input, out *output) {
				m.seratRepository.On("Delete", in.ctx, in.id).Return(out.err)
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
				m.seratRepository.On("Delete", in.ctx, in.id).Return(out.err)
				m.seratEventEmitter.On("EmitDeletedEvent", mock.AnythingOfType("*event.Message")).Return(nil)
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
