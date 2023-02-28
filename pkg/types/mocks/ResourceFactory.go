// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	types "github.com/k8snetworkplumbingwg/sriov-network-device-plugin/pkg/types"
	mock "github.com/stretchr/testify/mock"
)

// ResourceFactory is an autogenerated mock type for the ResourceFactory type
type ResourceFactory struct {
	mock.Mock
}

// GetDefaultInfoProvider provides a mock function with given fields: _a0, _a1
func (_m *ResourceFactory) GetDefaultInfoProvider(_a0 string, _a1 string) []types.DeviceInfoProvider {
	ret := _m.Called(_a0, _a1)

	var r0 []types.DeviceInfoProvider
	if rf, ok := ret.Get(0).(func(string, string) []types.DeviceInfoProvider); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.DeviceInfoProvider)
		}
	}

	return r0
}

// GetDeviceFilter provides a mock function with given fields: _a0
func (_m *ResourceFactory) GetDeviceFilter(_a0 *types.ResourceConfig) (interface{}, error) {
	ret := _m.Called(_a0)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*types.ResourceConfig) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.ResourceConfig) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceProvider provides a mock function with given fields: _a0
func (_m *ResourceFactory) GetDeviceProvider(_a0 types.DeviceType) types.DeviceProvider {
	ret := _m.Called(_a0)

	var r0 types.DeviceProvider
	if rf, ok := ret.Get(0).(func(types.DeviceType) types.DeviceProvider); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.DeviceProvider)
		}
	}

	return r0
}

// GetNadUtils provides a mock function with given fields:
func (_m *ResourceFactory) GetNadUtils() types.NadUtils {
	ret := _m.Called()

	var r0 types.NadUtils
	if rf, ok := ret.Get(0).(func() types.NadUtils); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.NadUtils)
		}
	}

	return r0
}

// GetRdmaSpec provides a mock function with given fields: _a0, _a1
func (_m *ResourceFactory) GetRdmaSpec(_a0 types.DeviceType, _a1 string) types.RdmaSpec {
	ret := _m.Called(_a0, _a1)

	var r0 types.RdmaSpec
	if rf, ok := ret.Get(0).(func(types.DeviceType, string) types.RdmaSpec); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.RdmaSpec)
		}
	}

	return r0
}

// GetResourcePool provides a mock function with given fields: rc, deviceList
func (_m *ResourceFactory) GetResourcePool(rc *types.ResourceConfig, deviceList []types.HostDevice) (types.ResourcePool, error) {
	ret := _m.Called(rc, deviceList)

	var r0 types.ResourcePool
	if rf, ok := ret.Get(0).(func(*types.ResourceConfig, []types.HostDevice) types.ResourcePool); ok {
		r0 = rf(rc, deviceList)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ResourcePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.ResourceConfig, []types.HostDevice) error); ok {
		r1 = rf(rc, deviceList)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourceServer provides a mock function with given fields: _a0
func (_m *ResourceFactory) GetResourceServer(_a0 types.ResourcePool) (types.ResourceServer, error) {
	ret := _m.Called(_a0)

	var r0 types.ResourceServer
	if rf, ok := ret.Get(0).(func(types.ResourcePool) types.ResourceServer); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ResourceServer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.ResourcePool) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSelector provides a mock function with given fields: _a0, _a1
func (_m *ResourceFactory) GetSelector(_a0 string, _a1 []string) (types.DeviceSelector, error) {
	ret := _m.Called(_a0, _a1)

	var r0 types.DeviceSelector
	if rf, ok := ret.Get(0).(func(string, []string) types.DeviceSelector); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.DeviceSelector)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVdpaDevice provides a mock function with given fields: _a0
func (_m *ResourceFactory) GetVdpaDevice(_a0 string) types.VdpaDevice {
	ret := _m.Called(_a0)

	var r0 types.VdpaDevice
	if rf, ok := ret.Get(0).(func(string) types.VdpaDevice); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.VdpaDevice)
		}
	}

	return r0
}

type mockConstructorTestingTNewResourceFactory interface {
	mock.TestingT
	Cleanup(func())
}

// NewResourceFactory creates a new instance of ResourceFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewResourceFactory(t mockConstructorTestingTNewResourceFactory) *ResourceFactory {
	mock := &ResourceFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
