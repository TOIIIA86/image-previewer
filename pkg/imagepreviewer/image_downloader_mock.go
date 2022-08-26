// Code generated by MockGen. DO NOT EDIT.
// Source: image_downloader.go

// Package imagepreviewer is a generated GoMock package.
package imagepreviewer

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockImageDownloader is a mock of ImageDownloader interface.
type MockImageDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockImageDownloaderMockRecorder
}

// MockImageDownloaderMockRecorder is the mock recorder for MockImageDownloader.
type MockImageDownloaderMockRecorder struct {
	mock *MockImageDownloader
}

// NewMockImageDownloader creates a new mock instance.
func NewMockImageDownloader(ctrl *gomock.Controller) *MockImageDownloader {
	mock := &MockImageDownloader{ctrl: ctrl}
	mock.recorder = &MockImageDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageDownloader) EXPECT() *MockImageDownloaderMockRecorder {
	return m.recorder
}

// Download mocks base method.
func (m *MockImageDownloader) Download(ctx context.Context, imageURL string, headers map[string][]string) (*DownloadResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", ctx, imageURL, headers)
	ret0, _ := ret[0].(*DownloadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download.
func (mr *MockImageDownloaderMockRecorder) Download(ctx, imageURL, headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockImageDownloader)(nil).Download), ctx, imageURL, headers)
}
