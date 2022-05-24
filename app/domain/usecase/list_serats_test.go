package usecase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/usecase"
	"github.com/stretchr/testify/assert"

	// "github.com/fikrirnurhidayat/kasusastran/app/domain/usecase"

	mockPostgresRepository "github.com/fikrirnurhidayat/kasusastran/mocks/domain/repository/postgres"
)

type MockListSeratsUseCase struct {
	SeratRepository *mockPostgresRepository.SeratRepositorier
}

func TestListSeratsUseCase_Exec(t *testing.T) {
	type input struct {
		ctx        context.Context
		pagination *entity.Pagination
	}

	type output struct {
		serats     []entity.Serat
		pagination *entity.Pagination
		err        error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockListSeratsUseCase, *input, *output)
	}

	tests := []scenario{
		{
			name: "SeratRepository.ListSerats return error",
			in: &input{
				ctx: context.Background(),
				pagination: &entity.Pagination{
					Page:     1,
					PageSize: 10,
				},
			},
			out: &output{
				err: fmt.Errorf("SeratRepository.ListSerats: failed to run query: bababoey"),
			},
			on: func(m *MockListSeratsUseCase, i *input, o *output) {
				m.SeratRepository.On("ListSerats", i.ctx, i.pagination).Return(o.serats, uint32(0), o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				pagination: &entity.Pagination{
					Page:     1,
					PageSize: 10,
				},
			},
			out: &output{
				serats: []entity.Serat{},
				pagination: &entity.Pagination{
					Page:      1,
					PageSize:  10,
					PageCount: 1,
					Total:     1,
				},
				err: nil,
			},
			on: func(m *MockListSeratsUseCase, i *input, o *output) {
				m.SeratRepository.On("ListSerats", i.ctx, i.pagination).Return(o.serats, uint32(1), nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockListSeratsUseCase{
				SeratRepository: new(mockPostgresRepository.SeratRepositorier),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := usecase.NewListSeratsUseCase(m.SeratRepository)
			serats, pagination, err := subject.Exec(tt.in.ctx, tt.in.pagination)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}

			assert.Equal(t, tt.out.serats, serats)
			assert.Equal(t, tt.out.pagination, pagination)
		})
	}
}
