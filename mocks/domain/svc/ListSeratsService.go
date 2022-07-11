// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	svc "github.com/fikrirnurhidayat/kasusastran/app/domain/svc"
	mock "github.com/stretchr/testify/mock"
)

// ListSeratsService is an autogenerated mock type for the ListSeratsService type
type ListSeratsService struct {
	mock.Mock
}

// Call provides a mock function with given fields: _a0, _a1
func (_m *ListSeratsService) Call(_a0 context.Context, _a1 *svc.ListSeratsParams) (*svc.ListSeratsResult, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *svc.ListSeratsResult
	if rf, ok := ret.Get(0).(func(context.Context, *svc.ListSeratsParams) *svc.ListSeratsResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*svc.ListSeratsResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *svc.ListSeratsParams) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
