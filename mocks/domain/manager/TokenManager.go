// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	manager "github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	mock "github.com/stretchr/testify/mock"
)

// TokenManager is an autogenerated mock type for the TokenManager type
type TokenManager struct {
	mock.Mock
}

// NewSession provides a mock function with given fields: userID
func (_m *TokenManager) NewSession(userID string) (*manager.Session, error) {
	ret := _m.Called(userID)

	var r0 *manager.Session
	if rf, ok := ret.Get(0).(func(string) *manager.Session); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*manager.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTokenManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewTokenManager creates a new instance of TokenManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTokenManager(t mockConstructorTestingTNewTokenManager) *TokenManager {
	mock := &TokenManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
