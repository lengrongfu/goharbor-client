// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	repository "github.com/mittwald/goharbor-client/apiv2/internal/api/client/repository"
	mock "github.com/stretchr/testify/mock"
)

// MockRepositoryClientService is an autogenerated mock type for the ClientService type
type MockRepositoryClientService struct {
	mock.Mock
}

// DeleteRepository provides a mock function with given fields: params, authInfo
func (_m *MockRepositoryClientService) DeleteRepository(params *repository.DeleteRepositoryParams, authInfo runtime.ClientAuthInfoWriter) (*repository.DeleteRepositoryOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *repository.DeleteRepositoryOK
	if rf, ok := ret.Get(0).(func(*repository.DeleteRepositoryParams, runtime.ClientAuthInfoWriter) *repository.DeleteRepositoryOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.DeleteRepositoryOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*repository.DeleteRepositoryParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepository provides a mock function with given fields: params, authInfo
func (_m *MockRepositoryClientService) GetRepository(params *repository.GetRepositoryParams, authInfo runtime.ClientAuthInfoWriter) (*repository.GetRepositoryOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *repository.GetRepositoryOK
	if rf, ok := ret.Get(0).(func(*repository.GetRepositoryParams, runtime.ClientAuthInfoWriter) *repository.GetRepositoryOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.GetRepositoryOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*repository.GetRepositoryParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRepositories provides a mock function with given fields: params, authInfo
func (_m *MockRepositoryClientService) ListRepositories(params *repository.ListRepositoriesParams, authInfo runtime.ClientAuthInfoWriter) (*repository.ListRepositoriesOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *repository.ListRepositoriesOK
	if rf, ok := ret.Get(0).(func(*repository.ListRepositoriesParams, runtime.ClientAuthInfoWriter) *repository.ListRepositoriesOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.ListRepositoriesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*repository.ListRepositoriesParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockRepositoryClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateRepository provides a mock function with given fields: params, authInfo
func (_m *MockRepositoryClientService) UpdateRepository(params *repository.UpdateRepositoryParams, authInfo runtime.ClientAuthInfoWriter) (*repository.UpdateRepositoryOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *repository.UpdateRepositoryOK
	if rf, ok := ret.Get(0).(func(*repository.UpdateRepositoryParams, runtime.ClientAuthInfoWriter) *repository.UpdateRepositoryOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.UpdateRepositoryOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*repository.UpdateRepositoryParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
