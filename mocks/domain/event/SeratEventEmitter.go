// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	message "github.com/fikrirnurhidayat/kasusastran/app/domain/message"
	mock "github.com/stretchr/testify/mock"
)

// SeratEventEmitter is an autogenerated mock type for the SeratEventEmitter type
type SeratEventEmitter struct {
	mock.Mock
}

// EmitCreatedEvent provides a mock function with given fields: _a0
func (_m *SeratEventEmitter) EmitCreatedEvent(_a0 *message.Serat) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*message.Serat) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EmitDeletedEvent provides a mock function with given fields: _a0
func (_m *SeratEventEmitter) EmitDeletedEvent(_a0 *message.Serat) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*message.Serat) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EmitListedEvent provides a mock function with given fields: _a0
func (_m *SeratEventEmitter) EmitListedEvent(_a0 *message.Serats) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*message.Serats) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EmitRetrievedEvent provides a mock function with given fields: _a0
func (_m *SeratEventEmitter) EmitRetrievedEvent(_a0 *message.Serat) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*message.Serat) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EmitUpdatedEvent provides a mock function with given fields: _a0
func (_m *SeratEventEmitter) EmitUpdatedEvent(_a0 *message.Serat) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*message.Serat) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewSeratEventEmitter interface {
	mock.TestingT
	Cleanup(func())
}

// NewSeratEventEmitter creates a new instance of SeratEventEmitter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSeratEventEmitter(t mockConstructorTestingTNewSeratEventEmitter) *SeratEventEmitter {
	mock := &SeratEventEmitter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
