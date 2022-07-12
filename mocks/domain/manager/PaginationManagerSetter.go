// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	manager "github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	mock "github.com/stretchr/testify/mock"
)

// PaginationManagerSetter is an autogenerated mock type for the PaginationManagerSetter type
type PaginationManagerSetter struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *PaginationManagerSetter) Execute(_a0 *manager.PaginationManagerImpl) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewPaginationManagerSetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewPaginationManagerSetter creates a new instance of PaginationManagerSetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPaginationManagerSetter(t mockConstructorTestingTNewPaginationManagerSetter) *PaginationManagerSetter {
	mock := &PaginationManagerSetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
