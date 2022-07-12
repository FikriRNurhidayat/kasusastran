// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// EventEmitter is an autogenerated mock type for the EventEmitter type
type EventEmitter struct {
	mock.Mock
}

// DeferredPublish provides a mock function with given fields: topic, at, message
func (_m *EventEmitter) DeferredPublish(topic string, at time.Time, message interface{}) error {
	ret := _m.Called(topic, at, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, time.Time, interface{}) error); ok {
		r0 = rf(topic, at, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Publish provides a mock function with given fields: topic, message
func (_m *EventEmitter) Publish(topic string, message interface{}) error {
	ret := _m.Called(topic, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(topic, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewEventEmitter interface {
	mock.TestingT
	Cleanup(func())
}

// NewEventEmitter creates a new instance of EventEmitter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventEmitter(t mockConstructorTestingTNewEventEmitter) *EventEmitter {
	mock := &EventEmitter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
