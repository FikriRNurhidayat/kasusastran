package svc_test

import (
	"context"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/fikrirnurhidayat/kasusastran/app/trouble"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockEvent "github.com/fikrirnurhidayat/kasusastran/mocks/domain/event"
	mockRepo "github.com/fikrirnurhidayat/kasusastran/mocks/domain/repository"
	mockPackage "github.com/fikrirnurhidayat/kasusastran/mocks/package"
)

type MockListSeratsService struct {
	logger            *mockPackage.LoggerV2
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
			name: "s.seratRepository.List return error",
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
				err: trouble.INTERNAL_SERVER_ERROR,
			},
			on: func(m *MockListSeratsService, i *input, o *output) {
				m.seratRepository.On("List", i.ctx, i.params.Pagination.ToListQuery()).Return(nil, uint32(0), o.err)
			},
		},
		{
			name: "s.seratEventEmitter.EmitListedEvent return error",
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
				result: nil,
				err:    trouble.INTERNAL_SERVER_ERROR,
			},
			on: func(m *MockListSeratsService, i *input, o *output) {
				result := &svc.ListSeratsResult{
					Serats: []*entity.Serat{
						{
							ID:                uuid.New(),
							Title:             "Lorem ipsum",
							Description:       "Lorem ipsum",
							CoverImageUrl:     "http://placekitten.com/192/108",
							ThumbnailImageUrl: "http://placekitten.com/192/108",
						},
					},
					Pagination: &manager.Pagination{
						Page:      1,
						PageSize:  10,
						PageCount: 1,
						Total:     1,
					},
				}
				m.seratRepository.On("List", i.ctx, i.params.Pagination.ToListQuery()).Return(result.Serats, uint32(1), nil)
				m.seratEventEmitter.On("EmitListedEvent", mock.AnythingOfType("*message.Serats")).Return(o.err)
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
					Serats: []*entity.Serat{
						{
							ID:                uuid.New(),
							Title:             "Lorem ipsum",
							Description:       "Lorem ipsum",
							CoverImageUrl:     "http://placekitten.com/192/108",
							ThumbnailImageUrl: "http://placekitten.com/192/108",
						},
					},
				},
				err: nil,
			},
			on: func(m *MockListSeratsService, i *input, o *output) {
				m.seratRepository.On("List", i.ctx, i.params.Pagination.ToListQuery()).Return(o.result.Serats, uint32(1), nil)
				m.seratEventEmitter.On("EmitListedEvent", mock.AnythingOfType("*message.Serats")).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockListSeratsService{
				logger:            new(mockPackage.LoggerV2),
				seratRepository:   new(mockRepo.SeratRepository),
				seratEventEmitter: new(mockEvent.SeratEventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			m.logger.On("Error", mock.AnythingOfType("*status.Error"))

			subject := svc.NewListSeratsService(m.logger, m.seratRepository, m.seratEventEmitter)
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
