// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/ryanadiputraa/flows/flows-microservices/user/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *Repository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	ret := _m.Called(ctx, email)

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, id
func (_m *Repository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, _a1
func (_m *Repository) Save(ctx context.Context, _a1 *domain.User) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
