// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TestingT is an autogenerated mock type for the TestingT type
type TestingT struct {
	mock.Mock
}

// Errorf provides a mock function with given fields: format, args
func (_m *TestingT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type mockConstructorTestingTNewTestingT interface {
	mock.TestingT
	Cleanup(func())
}

// NewTestingT creates a new instance of TestingT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTestingT(t mockConstructorTestingTNewTestingT) *TestingT {
	mock := &TestingT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
