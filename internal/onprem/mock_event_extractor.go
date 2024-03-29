// Code generated by MockGen. DO NOT EDIT.
// Source: event_extractor.go

// Package onprem is a generated GoMock package.
package onprem

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/openshift-assisted/assisted-events-streams/internal/types"
)

// MockIEventExtractor is a mock of IEventExtractor interface.
type MockIEventExtractor struct {
	ctrl     *gomock.Controller
	recorder *MockIEventExtractorMockRecorder
}

// MockIEventExtractorMockRecorder is the mock recorder for MockIEventExtractor.
type MockIEventExtractorMockRecorder struct {
	mock *MockIEventExtractor
}

// NewMockIEventExtractor creates a new mock instance.
func NewMockIEventExtractor(ctrl *gomock.Controller) *MockIEventExtractor {
	mock := &MockIEventExtractor{ctrl: ctrl}
	mock.recorder = &MockIEventExtractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIEventExtractor) EXPECT() *MockIEventExtractorMockRecorder {
	return m.recorder
}

// ExtractEvents mocks base method.
func (m *MockIEventExtractor) ExtractEvents(tarFilename string) (chan types.EventEnvelope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractEvents", tarFilename)
	ret0, _ := ret[0].(chan types.EventEnvelope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtractEvents indicates an expected call of ExtractEvents.
func (mr *MockIEventExtractorMockRecorder) ExtractEvents(tarFilename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractEvents", reflect.TypeOf((*MockIEventExtractor)(nil).ExtractEvents), tarFilename)
}
