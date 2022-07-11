package svc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockEvent "github.com/fikrirnurhidayat/kasusastran/mocks/domain/event"
	mockRepo "github.com/fikrirnurhidayat/kasusastran/mocks/domain/repository"
)

type MockListSeratsService struct {
	seratEventEmitter *mockEvent.SeratEventEmitter
	seratRepository   *mockRepo.SeratRepository
}

func TestListSeratsService_Call(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *svc.ListSeratsParams
	}

	type output struct {
		result *svc.ListSeratsResult
		err    error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockListSeratsService, *input, *output)
	}

	tests := []scenario{
		{
			name: "SeratRepository.ListSerats return error",
			in: &input{
				ctx: context.Background(),
				params: &svc.ListSeratsParams{
					Pagination: &manager.Pagination{
						Page:     1,
						PageSize: 10,
					},
				},
			},
			out: &output{
				err: fmt.Errorf("seratRepository.List: failed to run query: bababoey"),
			},
			on: func(m *MockListSeratsService, i *input, o *output) {
				m.seratRepository.On("List", i.ctx, i.params.Pagination.ToListQuery()).Return(nil, uint32(0), o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				params: &svc.ListSeratsParams{
					Pagination: &manager.Pagination{
						Page:     1,
						PageSize: 10,
					},
				},
			},
			out: &output{
				result: &svc.ListSeratsResult{
					Pagination: &manager.Pagination{
						Page:      1,
						PageSize:  10,
						PageCount: 1,
						Total:     1,
					},
					Serats: []*entity.Serat{},
				},
				err: nil,
			},
			on: func(m *MockListSeratsService, i *input, o *output) {
				m.seratRepository.On("List", i.ctx, i.params.Pagination.ToListQuery()).Return(o.result.Serats, uint32(1), nil)
				m.seratEventEmitter.On("EmitListedEvent", mock.AnythingOfType("*event.Message")).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockListSeratsService{
				seratRepository:   new(mockRepo.SeratRepository),
				seratEventEmitter: new(mockEvent.SeratEventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := svc.NewListSeratsService(m.seratRepository, m.seratEventEmitter)
			got, err := subject.Call(tt.in.ctx, tt.in.params)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}

			if tt.out.result != nil {
				assert.Equal(t, tt.out.result.Serats, got.Serats)
				assert.Equal(t, tt.out.result.Pagination, got.Pagination)
			}
		})
	}
}
