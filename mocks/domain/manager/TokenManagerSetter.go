// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	manager "github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	mock "github.com/stretchr/testify/mock"
)

// TokenManagerSetter is an autogenerated mock type for the TokenManagerSetter type
type TokenManagerSetter struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *TokenManagerSetter) Execute(_a0 *manager.TokenManagerImpl) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewTokenManagerSetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewTokenManagerSetter creates a new instance of TokenManagerSetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTokenManagerSetter(t mockConstructorTestingTNewTokenManagerSetter) *TokenManagerSetter {
	mock := &TokenManagerSetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}