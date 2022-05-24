package usecase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/usecase"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	mockPostgres "github.com/fikrirnurhidayat/kasusastran/mocks/domain/repository/postgres"
)

type MockGetSeratUseCase struct {
	SeratRepository *mockPostgres.SeratRepositorier
}

func TestGetSeratUseCase_Exec(t *testing.T) {
	type input struct {
		ctx context.Context
		id  uuid.UUID
	}

	type output struct {
		serat *entity.Serat
		err   error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockGetSeratUseCase, *input, *output)
	}

	tests := []scenario{
		{
			name: "SeratRepository.GetSerat return error",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				serat: nil,
				err:   fmt.Errorf("GetSerat: failed to retrieve: baboey"),
			},
			on: func(m *MockGetSeratUseCase, in *input, out *output) {
				m.SeratRepository.On("GetSerat", in.ctx, in.id).Return(nil, out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				serat: &entity.Serat{
					ID:                uuid.New(),
					Title:             "Jayabaya",
					Description:       "Lorem ipsum dolor sit amet",
					CoverImageUrl:     "https://placeimg.com/640/480/any",
					ThumbnailImageUrl: "https://placeimg.com/640/480/any",
				},
				err: nil,
			},
			on: func(m *MockGetSeratUseCase, in *input, out *output) {
				m.SeratRepository.On("GetSerat", in.ctx, in.id).Return(out.serat, out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockGetSeratUseCase{
				SeratRepository: new(mockPostgres.SeratRepositorier),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := usecase.NewGetSeratUseCase(m.SeratRepository)
			out, err := subject.Exec(tt.in.ctx, tt.in.id)

			if err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}

			assert.Equal(t, tt.out.serat, out)
		})
	}
}
