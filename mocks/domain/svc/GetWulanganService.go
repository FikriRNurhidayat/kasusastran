// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	svc "github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// GetWulanganService is an autogenerated mock type for the GetWulanganService type
type GetWulanganService struct {
	mock.Mock
}

// Call provides a mock function with given fields: ctx, id
func (_m *GetWulanganService) Call(ctx context.Context, id uuid.UUID) (*svc.GetWulanganResult, error) {
	ret := _m.Called(ctx, id)

	var r0 *svc.GetWulanganResult
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *svc.GetWulanganResult); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*svc.GetWulanganResult)
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

type mockConstructorTestingTNewGetWulanganService interface {
	mock.TestingT
	Cleanup(func())
}

// NewGetWulanganService creates a new instance of GetWulanganService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGetWulanganService(t mockConstructorTestingTNewGetWulanganService) *GetWulanganService {
	mock := &GetWulanganService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
