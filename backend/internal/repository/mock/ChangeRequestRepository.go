// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	domain "game-server/internal/core/domain"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// ChangeRequestRepository is an autogenerated mock type for the ChangeRequestRepository type
type ChangeRequestRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0
func (_m *ChangeRequestRepository) Add(_a0 domain.ChangeRequest) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(domain.ChangeRequest) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.ChangeRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: request_id
func (_m *ChangeRequestRepository) Delete(request_id string) error {
	ret := _m.Called(request_id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(request_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: request_id
func (_m *ChangeRequestRepository) Get(request_id string) (domain.ChangeRequest, error) {
	ret := _m.Called(request_id)

	var r0 domain.ChangeRequest
	if rf, ok := ret.Get(0).(func(string) domain.ChangeRequest); ok {
		r0 = rf(request_id)
	} else {
		r0 = ret.Get(0).(domain.ChangeRequest)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(request_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewChangeRequestRepository creates a new instance of ChangeRequestRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewChangeRequestRepository(t testing.TB) *ChangeRequestRepository {
	mock := &ChangeRequestRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
