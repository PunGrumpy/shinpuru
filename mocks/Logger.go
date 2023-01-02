// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	guildlog "github.com/zekroTJA/shinpuru/internal/services/guildlog"
)

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// Debugf provides a mock function with given fields: guildID, message, data
func (_m *Logger) Debugf(guildID string, message string, data ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, guildID, message)
	_ca = append(_ca, data...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, ...interface{}) error); ok {
		r0 = rf(guildID, message, data...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Errorf provides a mock function with given fields: guildID, message, data
func (_m *Logger) Errorf(guildID string, message string, data ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, guildID, message)
	_ca = append(_ca, data...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, ...interface{}) error); ok {
		r0 = rf(guildID, message, data...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fatalf provides a mock function with given fields: guildID, message, data
func (_m *Logger) Fatalf(guildID string, message string, data ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, guildID, message)
	_ca = append(_ca, data...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, ...interface{}) error); ok {
		r0 = rf(guildID, message, data...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Infof provides a mock function with given fields: guildID, message, data
func (_m *Logger) Infof(guildID string, message string, data ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, guildID, message)
	_ca = append(_ca, data...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, ...interface{}) error); ok {
		r0 = rf(guildID, message, data...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Section provides a mock function with given fields: module
func (_m *Logger) Section(module string) guildlog.Logger {
	ret := _m.Called(module)

	var r0 guildlog.Logger
	if rf, ok := ret.Get(0).(func(string) guildlog.Logger); ok {
		r0 = rf(module)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(guildlog.Logger)
		}
	}

	return r0
}

// Warnf provides a mock function with given fields: guildID, message, data
func (_m *Logger) Warnf(guildID string, message string, data ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, guildID, message)
	_ca = append(_ca, data...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, ...interface{}) error); ok {
		r0 = rf(guildID, message, data...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewLogger interface {
	mock.TestingT
	Cleanup(func())
}

// NewLogger creates a new instance of Logger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLogger(t mockConstructorTestingTNewLogger) *Logger {
	mock := &Logger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
