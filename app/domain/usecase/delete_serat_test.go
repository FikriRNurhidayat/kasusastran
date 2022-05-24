package usecase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/usecase"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	mockPostgres "github.com/fikrirnurhidayat/kasusastran/mocks/domain/repository/postgres"
)

type MockDeleteSeratUseCase struct {
	SeratRepository *mockPostgres.SeratRepositorier
}

func TestDeleteSeratUseCase_Exec(t *testing.T) {
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
		on   func(*MockDeleteSeratUseCase, *input, *output)
	}

	tests := []scenario{
		{
			name: "SeratRepository.DeleteSerat return error",
			in: &input{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			out: &output{
				err: fmt.Errorf("DeleteSerat: failed to retrieve: baboey"),
			},
			on: func(m *MockDeleteSeratUseCase, in *input, out *output) {
				m.SeratRepository.On("DeleteSerat", in.ctx, in.id).Return(out.err)
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
			on: func(m *MockDeleteSeratUseCase, in *input, out *output) {
				m.SeratRepository.On("DeleteSerat", in.ctx, in.id).Return(out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDeleteSeratUseCase{
				SeratRepository: new(mockPostgres.SeratRepositorier),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := usecase.NewDeleteSeratUseCase(m.SeratRepository)
			err := subject.Exec(tt.in.ctx, tt.in.id)

			if err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}
		})
	}
}
