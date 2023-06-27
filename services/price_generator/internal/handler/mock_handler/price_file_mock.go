// Code generated by MockGen. DO NOT EDIT.
// Source: price_file.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockServicePriceFile is a mock of ServicePriceFile interface.
type MockServicePriceFile struct {
	ctrl     *gomock.Controller
	recorder *MockServicePriceFileMockRecorder
}

// MockServicePriceFileMockRecorder is the mock recorder for MockServicePriceFile.
type MockServicePriceFileMockRecorder struct {
	mock *MockServicePriceFile
}

// NewMockServicePriceFile creates a new mock instance.
func NewMockServicePriceFile(ctrl *gomock.Controller) *MockServicePriceFile {
	mock := &MockServicePriceFile{ctrl: ctrl}
	mock.recorder = &MockServicePriceFileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicePriceFile) EXPECT() *MockServicePriceFileMockRecorder {
	return m.recorder
}

// DeleteFileByName mocks base method.
func (m *MockServicePriceFile) DeleteFileByName(fileName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFileByName", fileName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFileByName indicates an expected call of DeleteFileByName.
func (mr *MockServicePriceFileMockRecorder) DeleteFileByName(fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFileByName", reflect.TypeOf((*MockServicePriceFile)(nil).DeleteFileByName), fileName)
}

// GetFileByName mocks base method.
func (m *MockServicePriceFile) GetFileByName(fileName string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileByName", fileName)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileByName indicates an expected call of GetFileByName.
func (mr *MockServicePriceFileMockRecorder) GetFileByName(fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileByName", reflect.TypeOf((*MockServicePriceFile)(nil).GetFileByName), fileName)
}

// SetFile mocks base method.
func (m *MockServicePriceFile) SetFile(fileName, barcode, title string, cost int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFile", fileName, barcode, title, cost)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFile indicates an expected call of SetFile.
func (mr *MockServicePriceFileMockRecorder) SetFile(fileName, barcode, title, cost interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFile", reflect.TypeOf((*MockServicePriceFile)(nil).SetFile), fileName, barcode, title, cost)
}
