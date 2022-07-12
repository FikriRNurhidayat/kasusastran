package svc_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockEvent "github.com/fikrirnurhidayat/kasusastran/mocks/domain/event"
	mockRepo "github.com/fikrirnurhidayat/kasusastran/mocks/domain/repository"
)

type MockGetWulanganService struct {
	wulanganRepository   *mockRepo.WulanganRepository
	wulanganEventEmitter *mockEvent.WulanganEventEmitter
}

func TestGetWulanganService_Call(t *testing.T) {
	type input struct {
		ctx context.Context
		id  uuid.UUID
	}

	type output struct {
		w   *svc.GetWulanganResult
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockGetWulanganService, *input, *output)
	}

	tests := []scenario{
		{
			name: "s.wulanganRepository.Get return error",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				w:   nil,
				err: errors.New("s.wulanganRepository.Get: failed to retrieve a wulangan"),
			},
			on: func(mgws *MockGetWulanganService, i *input, o *output) {
				mgws.wulanganRepository.On("Get", i.ctx, i.id).Return(nil, o.err)
			},
		},
		{
			name: "Not Found",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				w:   nil,
				err: errors.New("wulangan not found"),
			},
			on: func(mgws *MockGetWulanganService, i *input, o *output) {
				mgws.wulanganRepository.On("Get", i.ctx, i.id).Return(nil, nil)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				w: &svc.GetWulanganResult{
					Title:             "Industrial Society And Its Future",
					Description:       "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "https://placekitten.com/192/108",
					ThumbnailImageUrl: "https://placekitten.com/192/108",
					CreatedAt:         time.Now(),
					UpdatedAt:         time.Now(),
				},
				err: nil,
			},
			on: func(mgws *MockGetWulanganService, i *input, o *output) {
				wulangan := &entity.Wulangan{
					ID:                i.id,
					Title:             o.w.Title,
					Description:       o.w.Description,
					CoverImageUrl:     o.w.CoverImageUrl,
					ThumbnailImageUrl: o.w.ThumbnailImageUrl,
					CreatedAt:         o.w.CreatedAt,
					UpdatedAt:         o.w.UpdatedAt,
				}

				o.w.ID = wulangan.ID

				mgws.wulanganRepository.On("Get", i.ctx, i.id).Return(wulangan, nil)
				mgws.wulanganEventEmitter.On("EmitRetrievedEvent", mock.AnythingOfType("*message.Wulangan")).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockGetWulanganService{
				wulanganRepository:   new(mockRepo.WulanganRepository),
				wulanganEventEmitter: new(mockEvent.WulanganEventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := svc.NewGetWulanganService(m.wulanganRepository, m.wulanganEventEmitter)
			got, err := subject.Call(tt.in.ctx, tt.in.id)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Error(t, tt.out.err, err)
			}

			if tt.out.w != nil {
				assert.NotNil(t, got)
				assert.Equal(t, tt.out.w, got)
			}
		})
	}
}
