// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// WulanganRepository is an autogenerated mock type for the WulanganRepository type
type WulanganRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, w
func (_m *WulanganRepository) Create(ctx context.Context, w *entity.Wulangan) (*entity.Wulangan, error) {
	ret := _m.Called(ctx, w)

	var r0 *entity.Wulangan
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Wulangan) *entity.Wulangan); ok {
		r0 = rf(ctx, w)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Wulangan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Wulangan) error); ok {
		r1 = rf(ctx, w)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *WulanganRepository) Delete(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *WulanganRepository) Get(ctx context.Context, id uuid.UUID) (*entity.Wulangan, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Wulangan
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *entity.Wulangan); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Wulangan)
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

// List provides a mock function with given fields: ctx
func (_m *WulanganRepository) List(ctx context.Context) ([]*entity.Wulangan, error) {
	ret := _m.Called(ctx)

	var r0 []*entity.Wulangan
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.Wulangan); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Wulangan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, w
func (_m *WulanganRepository) Update(ctx context.Context, id uuid.UUID, w *entity.Wulangan) (*entity.Wulangan, error) {
	ret := _m.Called(ctx, id, w)

	var r0 *entity.Wulangan
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, *entity.Wulangan) *entity.Wulangan); ok {
		r0 = rf(ctx, id, w)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Wulangan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, *entity.Wulangan) error); ok {
		r1 = rf(ctx, id, w)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewWulanganRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewWulanganRepository creates a new instance of WulanganRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWulanganRepository(t mockConstructorTestingTNewWulanganRepository) *WulanganRepository {
	mock := &WulanganRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
