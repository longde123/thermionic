// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spoke-d/thermionic/internal/db/cluster (interfaces: ServerStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	cluster "github.com/spoke-d/thermionic/internal/db/cluster"
	reflect "reflect"
)

// MockServerStore is a mock of ServerStore interface
type MockServerStore struct {
	ctrl     *gomock.Controller
	recorder *MockServerStoreMockRecorder
}

// MockServerStoreMockRecorder is the mock recorder for MockServerStore
type MockServerStoreMockRecorder struct {
	mock *MockServerStore
}

// NewMockServerStore creates a new mock instance
func NewMockServerStore(ctrl *gomock.Controller) *MockServerStore {
	mock := &MockServerStore{ctrl: ctrl}
	mock.recorder = &MockServerStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerStore) EXPECT() *MockServerStoreMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockServerStore) Get(arg0 context.Context) ([]cluster.ServerInfo, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]cluster.ServerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockServerStoreMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServerStore)(nil).Get), arg0)
}

// Set mocks base method
func (m *MockServerStore) Set(arg0 context.Context, arg1 []cluster.ServerInfo) error {
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockServerStoreMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockServerStore)(nil).Set), arg0, arg1)
}
