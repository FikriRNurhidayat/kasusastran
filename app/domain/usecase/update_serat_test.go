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

type MockUpdateSeratUseCase struct {
	SeratRepository *mockPostgres.SeratRepositorier
}

func TestUpdateSeratUseCase_Exec(t *testing.T) {
	type input struct {
		ctx    context.Context
		id     uuid.UUID
		params *usecase.UpdateSeratParams
	}

	type output struct {
		serat *entity.Serat
		err   error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockUpdateSeratUseCase, *input, *output)
	}

	tests := []scenario{
		{
			name: "SeratRepository.UpdateSerat return error",
			in: &input{
				ctx:    context.Background(),
				id:     uuid.New(),
				params: &usecase.UpdateSeratParams{},
			},
			out: &output{
				serat: nil,
				err:   fmt.Errorf("UpdateSerat: failed to retrieve: baboey"),
			},
			on: func(m *MockUpdateSeratUseCase, in *input, out *output) {
				m.SeratRepository.On("UpdateSerat", in.ctx, in.id, mock.AnythingOfType("*entity.Serat")).Return(nil, out.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				id:     uuid.New(),
				params: &usecase.UpdateSeratParams{},
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
			on: func(m *MockUpdateSeratUseCase, in *input, out *output) {
				m.SeratRepository.On("UpdateSerat", in.ctx, in.id, mock.AnythingOfType("*entity.Serat")).Return(out.serat, out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockUpdateSeratUseCase{
				SeratRepository: new(mockPostgres.SeratRepositorier),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := usecase.NewUpdateSeratUseCase(m.SeratRepository)
			out, err := subject.Exec(tt.in.ctx, tt.in.id, tt.in.params)

			if err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}

			assert.Equal(t, tt.out.serat, out)
		})
	}
}
