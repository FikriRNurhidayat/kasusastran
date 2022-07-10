// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// GetWulanganService is an autogenerated mock type for the GetWulanganService type
type GetWulanganService struct {
	mock.Mock
}

// Call provides a mock function with given fields: ctx, id
func (_m *GetWulanganService) Call(ctx context.Context, id uuid.UUID) (*entity.Wulangan, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Wulangan
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *entity.Wulangan); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Wulangan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
