// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	discordgo "github.com/bwmarrin/discordgo"
	mock "github.com/stretchr/testify/mock"
)

// VerificationProvider is an autogenerated mock type for the Provider type
type VerificationProvider struct {
	mock.Mock
}

// EnqueueVerification provides a mock function with given fields: member
func (_m *VerificationProvider) EnqueueVerification(member discordgo.Member) error {
	ret := _m.Called(member)

	var r0 error
	if rf, ok := ret.Get(0).(func(discordgo.Member) error); ok {
		r0 = rf(member)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetEnabled provides a mock function with given fields: guildID
func (_m *VerificationProvider) GetEnabled(guildID string) (bool, error) {
	ret := _m.Called(guildID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(guildID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(guildID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsVerified provides a mock function with given fields: userID
func (_m *VerificationProvider) IsVerified(userID string) (bool, error) {
	ret := _m.Called(userID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KickRoutine provides a mock function with given fields:
func (_m *VerificationProvider) KickRoutine() {
	_m.Called()
}

// SetEnabled provides a mock function with given fields: guildID, enabled
func (_m *VerificationProvider) SetEnabled(guildID string, enabled bool) error {
	ret := _m.Called(guildID, enabled)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool) error); ok {
		r0 = rf(guildID, enabled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Verify provides a mock function with given fields: userID
func (_m *VerificationProvider) Verify(userID string) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewVerificationProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewVerificationProvider creates a new instance of VerificationProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewVerificationProvider(t mockConstructorTestingTNewVerificationProvider) *VerificationProvider {
	mock := &VerificationProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
