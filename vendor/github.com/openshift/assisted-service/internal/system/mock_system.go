// Code generated by MockGen. DO NOT EDIT.
// Source: system.go

// Package system is a generated GoMock package.
package system

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSystemInfo is a mock of SystemInfo interface.
type MockSystemInfo struct {
	ctrl     *gomock.Controller
	recorder *MockSystemInfoMockRecorder
}

// MockSystemInfoMockRecorder is the mock recorder for MockSystemInfo.
type MockSystemInfoMockRecorder struct {
	mock *MockSystemInfo
}

// NewMockSystemInfo creates a new mock instance.
func NewMockSystemInfo(ctrl *gomock.Controller) *MockSystemInfo {
	mock := &MockSystemInfo{ctrl: ctrl}
	mock.recorder = &MockSystemInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSystemInfo) EXPECT() *MockSystemInfoMockRecorder {
	return m.recorder
}

// FIPSEnabled mocks base method.
func (m *MockSystemInfo) FIPSEnabled() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FIPSEnabled")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FIPSEnabled indicates an expected call of FIPSEnabled.
func (mr *MockSystemInfoMockRecorder) FIPSEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FIPSEnabled", reflect.TypeOf((*MockSystemInfo)(nil).FIPSEnabled))
}
