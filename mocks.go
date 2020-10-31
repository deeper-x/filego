package main

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileManager is a mock of FileManager interface
type MockFileManager struct {
	ctrl     *gomock.Controller
	recorder *MockFileManagerMockRecorder
}

// MockFileManagerMockRecorder is the mock recorder for MockFileManager
type MockFileManagerMockRecorder struct {
	mock *MockFileManager
}

// NewMockFileManager creates a new mock instance
func NewMockFileManager(ctrl *gomock.Controller) *MockFileManager {
	mock := &MockFileManager{ctrl: ctrl}
	mock.recorder = &MockFileManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileManager) EXPECT() *MockFileManagerMockRecorder {
	return m.recorder
}

// open mocks base method
func (m *MockFileManager) open(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "open", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// open indicates an expected call of open
func (mr *MockFileManagerMockRecorder) open(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "open", reflect.TypeOf((*MockFileManager)(nil).open), arg0)
}

// close mocks base method
func (m *MockFileManager) close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "close")
	ret0, _ := ret[0].(error)
	return ret0
}

// close indicates an expected call of close
func (mr *MockFileManagerMockRecorder) close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "close", reflect.TypeOf((*MockFileManager)(nil).close))
}

// WriteLine mocks base method
func (m *MockFileManager) WriteLine(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteLine", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteLine indicates an expected call of WriteLine
func (mr *MockFileManagerMockRecorder) WriteLine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteLine", reflect.TypeOf((*MockFileManager)(nil).WriteLine), arg0)
}

// DeleteLine mocks base method
func (m *MockFileManager) DeleteLine(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLine", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLine indicates an expected call of DeleteLine
func (mr *MockFileManagerMockRecorder) DeleteLine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLine", reflect.TypeOf((*MockFileManager)(nil).DeleteLine), arg0)
}
