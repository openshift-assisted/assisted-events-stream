// Code generated by MockGen. DO NOT EDIT.
// Source: downloader.go

// Package onprem is a generated GoMock package.
package onprem

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIFileDownloader is a mock of IFileDownloader interface.
type MockIFileDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockIFileDownloaderMockRecorder
}

// MockIFileDownloaderMockRecorder is the mock recorder for MockIFileDownloader.
type MockIFileDownloaderMockRecorder struct {
	mock *MockIFileDownloader
}

// NewMockIFileDownloader creates a new mock instance.
func NewMockIFileDownloader(ctrl *gomock.Controller) *MockIFileDownloader {
	mock := &MockIFileDownloader{ctrl: ctrl}
	mock.recorder = &MockIFileDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFileDownloader) EXPECT() *MockIFileDownloaderMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIFileDownloader) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockIFileDownloaderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIFileDownloader)(nil).Close))
}

// DownloadFile mocks base method.
func (m *MockIFileDownloader) DownloadFile(rawFileURL string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadFile", rawFileURL)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadFile indicates an expected call of DownloadFile.
func (mr *MockIFileDownloaderMockRecorder) DownloadFile(rawFileURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFile", reflect.TypeOf((*MockIFileDownloader)(nil).DownloadFile), rawFileURL)
}
