// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "jadwaldokter/features/dokter/domain"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddDokter provides a mock function with given fields: newItem
func (_m *Service) AddDokter(newItem domain.DokterCore) (domain.DokterCore, error) {
	ret := _m.Called(newItem)

	var r0 domain.DokterCore
	if rf, ok := ret.Get(0).(func(domain.DokterCore) domain.DokterCore); ok {
		r0 = rf(newItem)
	} else {
		r0 = ret.Get(0).(domain.DokterCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.DokterCore) error); ok {
		r1 = rf(newItem)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDokter provides a mock function with given fields: ID
func (_m *Service) DeleteDokter(ID uint) (domain.DokterCore, error) {
	ret := _m.Called(ID)

	var r0 domain.DokterCore
	if rf, ok := ret.Get(0).(func(uint) domain.DokterCore); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(domain.DokterCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDokter provides a mock function with given fields:
func (_m *Service) GetDokter() ([]domain.DokterCore, error) {
	ret := _m.Called()

	var r0 []domain.DokterCore
	if rf, ok := ret.Get(0).(func() []domain.DokterCore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.DokterCore)
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

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}