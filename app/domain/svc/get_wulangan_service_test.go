package svc_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	mocks "github.com/fikrirnurhidayat/kasusastran/mocks/domain/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type MockGetWulanganService struct {
	wulanganRepository *mocks.WulanganRepository
}

func TestGetWulanganService_Call(t *testing.T) {
	type input struct {
		ctx context.Context
		id  uuid.UUID
	}

	type output struct {
		w   *entity.Wulangan
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
				w: &entity.Wulangan{
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
				o.w.ID = i.id
				mgws.wulanganRepository.On("Get", i.ctx, i.id).Return(o.w, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockGetWulanganService{
				wulanganRepository: new(mocks.WulanganRepository),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := svc.NewGetWulanganService(m.wulanganRepository)
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
