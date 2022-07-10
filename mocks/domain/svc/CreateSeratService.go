// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	svc "github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	mock "github.com/stretchr/testify/mock"
)

// CreateSeratService is an autogenerated mock type for the CreateSeratService type
type CreateSeratService struct {
	mock.Mock
}

// Call provides a mock function with given fields: ctx, params
func (_m *CreateSeratService) Call(ctx context.Context, params *svc.CreateSeratParams) (*svc.CreateSeratResult, error) {
	ret := _m.Called(ctx, params)

	var r0 *svc.CreateSeratResult
	if rf, ok := ret.Get(0).(func(context.Context, *svc.CreateSeratParams) *svc.CreateSeratResult); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*svc.CreateSeratResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *svc.CreateSeratParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
