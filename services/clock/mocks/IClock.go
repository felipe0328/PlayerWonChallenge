// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// IClock is an autogenerated mock type for the IClock type
type IClock struct {
	mock.Mock
}

// Now provides a mock function with given fields:
func (_m *IClock) Now() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// Parse provides a mock function with given fields: _a0, _a1
func (_m *IClock) Parse(_a0 string, _a1 string) (time.Time, error) {
	ret := _m.Called(_a0, _a1)

	var r0 time.Time
	if rf, ok := ret.Get(0).(func(string, string) time.Time); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIClock interface {
	mock.TestingT
	Cleanup(func())
}

// NewIClock creates a new instance of IClock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIClock(t mockConstructorTestingTNewIClock) *IClock {
	mock := &IClock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
