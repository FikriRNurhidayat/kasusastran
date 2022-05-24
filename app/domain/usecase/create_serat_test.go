package usecase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/usecase"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockPostgres "github.com/fikrirnurhidayat/kasusastran/mocks/domain/repository/postgres"
)

type MockCreateSeratUseCase struct {
	SeratRepository *mockPostgres.SeratRepositorier
}

func TestCreateSeratUseCase_Exec(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *usecase.CreateSeratParams
	}

	type output struct {
		serat *entity.Serat
		err   error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockCreateSeratUseCase, *input, *output)
	}

	tests := []scenario{
		{
			name: "SeratRepository.CreateSerat return error",
			in: &input{
				ctx:    context.Background(),
				params: &usecase.CreateSeratParams{},
			},
			out: &output{
				serat: nil,
				err:   fmt.Errorf("CreateSerat: failed to retrieve: baboey"),
			},
			on: func(m *MockCreateSeratUseCase, in *input, out *output) {
				m.SeratRepository.On("CreateSerat", in.ctx, mock.AnythingOfType("*entity.Serat")).Return(nil, out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: &usecase.CreateSeratParams{},
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
			on: func(m *MockCreateSeratUseCase, in *input, out *output) {
				m.SeratRepository.On("CreateSerat", in.ctx, mock.AnythingOfType("*entity.Serat")).Return(out.serat, out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockCreateSeratUseCase{
				SeratRepository: new(mockPostgres.SeratRepositorier),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := usecase.NewCreateSeratUseCase(m.SeratRepository)
			out, err := subject.Exec(tt.in.ctx, tt.in.params)

			if err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}

			assert.Equal(t, tt.out.serat, out)
		})
	}
}
