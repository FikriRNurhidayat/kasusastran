package svc_test

import (
	"context"
	"errors"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	mocks "github.com/fikrirnurhidayat/kasusastran/mocks/domain/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type MockCreateWulanganService struct {
	wulanganRepository *mocks.WulanganRepository
}

func TestCreateWulanganService_Call(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *svc.CreateWulanganParams
	}

	type output struct {
		w   *entity.Wulangan
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockCreateWulanganService, *input, *output)
	}

	tests := []scenario{
		{
			name: "s.WulanganRepository.Create return error",
			in: &input{
				ctx: context.Background(),
				params: &svc.CreateWulanganParams{
					Title:             "Technological Society",
					Description:       "Our society is degrading.",
					CoverImageUrl:     "https://source.unsplash.com/617x598",
					ThumbnailImageUrl: "https://source.unsplash.com/617x598",
				},
			},
			out: &output{
				w:   nil,
				err: errors.New("s.WulanganRepository.Create: failed to insert data"),
			},
			on: func(mcws *MockCreateWulanganService, i *input, o *output) {
				w := &entity.Wulangan{
					Title:             i.params.Title,
					Description:       i.params.Description,
					CoverImageUrl:     i.params.CoverImageUrl,
					ThumbnailImageUrl: i.params.ThumbnailImageUrl,
				}

				mcws.wulanganRepository.On("Create", i.ctx, w).Return(o.w, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				params: &svc.CreateWulanganParams{
					Title:             "Technological Society",
					Description:       "Our society is degrading.",
					CoverImageUrl:     "https://source.unsplash.com/617x598",
					ThumbnailImageUrl: "https://source.unsplash.com/617x598",
				},
			},
			out: &output{
				w: &entity.Wulangan{
					Title:             "Technological Society",
					Description:       "Our society is degrading.",
					CoverImageUrl:     "https://source.unsplash.com/617x598",
					ThumbnailImageUrl: "https://source.unsplash.com/617x598",
				},
				err: nil,
			},
			on: func(mcws *MockCreateWulanganService, i *input, o *output) {
				w := &entity.Wulangan{
					Title:             i.params.Title,
					Description:       i.params.Description,
					CoverImageUrl:     i.params.CoverImageUrl,
					ThumbnailImageUrl: i.params.ThumbnailImageUrl,
				}

				o.w.ID = uuid.New()

				mcws.wulanganRepository.On("Create", i.ctx, w).Return(o.w, o.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockCreateWulanganService{
				wulanganRepository: new(mocks.WulanganRepository),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := svc.NewCreateWulanganService(m.wulanganRepository)
			got, err := subject.Call(tt.in.ctx, tt.in.params)

			if tt.out.err != nil {
				assert.NotNil(t, err)
			}

			if tt.out.w != nil {
				assert.NotNil(t, got)
				assert.Equal(t, tt.out.w.ID, got.ID)
				assert.Equal(t, tt.out.w.Title, got.Title)
				assert.Equal(t, tt.out.w.Description, got.Description)
				assert.Equal(t, tt.out.w.CoverImageUrl, got.CoverImageUrl)
				assert.Equal(t, tt.out.w.ThumbnailImageUrl, got.ThumbnailImageUrl)
			}
		})
	}
}
