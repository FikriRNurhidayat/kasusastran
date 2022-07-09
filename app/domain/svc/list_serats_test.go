package svc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fikrirnurhidayat/api.kasusastran.io/app/domain/entity"
	"github.com/fikrirnurhidayat/api.kasusastran.io/app/domain/svc"
	"github.com/stretchr/testify/assert"

	mocks "github.com/fikrirnurhidayat/api.kasusastran.io/mocks/domain/repository"
)

type MockListSeratsService struct {
	seratRepository *mocks.SeratRepository
}

func TestListSeratsService_Exec(t *testing.T) {
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
		on   func(*MockListSeratsService, *input, *output)
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
				err: fmt.Errorf("seratRepository.List: failed to run query: bababoey"),
			},
			on: func(m *MockListSeratsService, i *input, o *output) {
				m.seratRepository.On("List", i.ctx, i.pagination).Return(o.serats, uint32(0), o.err)
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
			on: func(m *MockListSeratsService, i *input, o *output) {
				m.seratRepository.On("List", i.ctx, i.pagination).Return(o.serats, uint32(1), nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockListSeratsService{
				seratRepository: new(mocks.SeratRepository),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := svc.NewListSeratsService(m.seratRepository)
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
