// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/slok/alertgram/internal/model"
)

// Notifier is an autogenerated mock type for the Notifier type
type Notifier struct {
	mock.Mock
}

// Notify provides a mock function with given fields: ctx, alertGroup
func (_m *Notifier) Notify(ctx context.Context, alertGroup *model.AlertGroup) error {
	ret := _m.Called(ctx, alertGroup)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.AlertGroup) error); ok {
		r0 = rf(ctx, alertGroup)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *Notifier) Type() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}