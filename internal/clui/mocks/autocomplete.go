// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spoke-d/thermionic/internal/clui (interfaces: AutoCompleteInstaller)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAutoCompleteInstaller is a mock of AutoCompleteInstaller interface
type MockAutoCompleteInstaller struct {
	ctrl     *gomock.Controller
	recorder *MockAutoCompleteInstallerMockRecorder
}

// MockAutoCompleteInstallerMockRecorder is the mock recorder for MockAutoCompleteInstaller
type MockAutoCompleteInstallerMockRecorder struct {
	mock *MockAutoCompleteInstaller
}

// NewMockAutoCompleteInstaller creates a new mock instance
func NewMockAutoCompleteInstaller(ctrl *gomock.Controller) *MockAutoCompleteInstaller {
	mock := &MockAutoCompleteInstaller{ctrl: ctrl}
	mock.recorder = &MockAutoCompleteInstallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAutoCompleteInstaller) EXPECT() *MockAutoCompleteInstallerMockRecorder {
	return m.recorder
}

// Install mocks base method
func (m *MockAutoCompleteInstaller) Install(arg0 string) error {
	ret := m.ctrl.Call(m, "Install", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install
func (mr *MockAutoCompleteInstallerMockRecorder) Install(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockAutoCompleteInstaller)(nil).Install), arg0)
}

// Uninstall mocks base method
func (m *MockAutoCompleteInstaller) Uninstall(arg0 string) error {
	ret := m.ctrl.Call(m, "Uninstall", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Uninstall indicates an expected call of Uninstall
func (mr *MockAutoCompleteInstallerMockRecorder) Uninstall(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninstall", reflect.TypeOf((*MockAutoCompleteInstaller)(nil).Uninstall), arg0)
}
