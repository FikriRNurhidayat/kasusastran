// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PasswordManager is an autogenerated mock type for the PasswordManager type
type PasswordManager struct {
	mock.Mock
}

// Compare provides a mock function with given fields: encryptedPassword, password
func (_m *PasswordManager) Compare(encryptedPassword string, password string) error {
	ret := _m.Called(encryptedPassword, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(encryptedPassword, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Encrypt provides a mock function with given fields: password
func (_m *PasswordManager) Encrypt(password string) (string, error) {
	ret := _m.Called(password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPasswordManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewPasswordManager creates a new instance of PasswordManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPasswordManager(t mockConstructorTestingTNewPasswordManager) *PasswordManager {
	mock := &PasswordManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
