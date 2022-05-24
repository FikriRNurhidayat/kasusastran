// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// ListSeratsUseCaser is an autogenerated mock type for the ListSeratsUseCaser type
type ListSeratsUseCaser struct {
	mock.Mock
}

// Exec provides a mock function with given fields: _a0, _a1
func (_m *ListSeratsUseCaser) Exec(_a0 context.Context, _a1 *entity.Pagination) ([]entity.Serat, *entity.Pagination, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []entity.Serat
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Pagination) []entity.Serat); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Serat)
		}
	}

	var r1 *entity.Pagination
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Pagination) *entity.Pagination); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*entity.Pagination)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *entity.Pagination) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
