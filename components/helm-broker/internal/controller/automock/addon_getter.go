// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import addon "github.com/kyma-project/kyma/components/helm-broker/internal/addon"

import mock "github.com/stretchr/testify/mock"

// AddonGetter is an autogenerated mock type for the AddonGetter type
type AddonGetter struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields:
func (_m *AddonGetter) Cleanup() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCompleteAddon provides a mock function with given fields: _a0
func (_m *AddonGetter) GetCompleteAddon(_a0 addon.EntryDTO) (addon.CompleteAddon, error) {
	ret := _m.Called(_a0)

	var r0 addon.CompleteAddon
	if rf, ok := ret.Get(0).(func(addon.EntryDTO) addon.CompleteAddon); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(addon.CompleteAddon)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(addon.EntryDTO) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIndex provides a mock function with given fields:
func (_m *AddonGetter) GetIndex() (*addon.IndexDTO, error) {
	ret := _m.Called()

	var r0 *addon.IndexDTO
	if rf, ok := ret.Get(0).(func() *addon.IndexDTO); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*addon.IndexDTO)
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