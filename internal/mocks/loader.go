// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import config "github.com/andrealbinop/go-yac/pkg/config"
import mock "github.com/stretchr/testify/mock"

// Loader is an autogenerated mock type for the Loader type
type Loader struct {
	mock.Mock
}

// Load provides a mock function with given fields:
func (_m *Loader) Load() (config.Provider, error) {
	ret := _m.Called()

	var r0 config.Provider
	if rf, ok := ret.Get(0).(func() config.Provider); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Provider)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Source provides a mock function with given fields:
func (_m *Loader) Source() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}