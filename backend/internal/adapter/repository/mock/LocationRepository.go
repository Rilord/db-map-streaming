// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	domain "game-server/internal/core/domain"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// LocationRepository is an autogenerated mock type for the LocationRepository type
type LocationRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0
func (_m *LocationRepository) Add(_a0 domain.Location) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Location) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: location_id
func (_m *LocationRepository) Get(location_id string) (domain.Location, error) {
	ret := _m.Called(location_id)

	var r0 domain.Location
	if rf, ok := ret.Get(0).(func(string) domain.Location); ok {
		r0 = rf(location_id)
	} else {
		r0 = ret.Get(0).(domain.Location)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(location_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *LocationRepository) Update(_a0 domain.Location) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Location) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewLocationRepository creates a new instance of LocationRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewLocationRepository(t testing.TB) *LocationRepository {
	mock := &LocationRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
