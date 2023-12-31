// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// JWT is an autogenerated mock type for the JWT type
type JWT struct {
	mock.Mock
}

// CreateToken provides a mock function with given fields:
func (_m *JWT) CreateToken() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewJWT creates a new instance of JWT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJWT(t interface {
	mock.TestingT
	Cleanup(func())
}) *JWT {
	mock := &JWT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
