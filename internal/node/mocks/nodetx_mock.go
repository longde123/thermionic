// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spoke-d/thermionic/internal/node (interfaces: Tx)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	db "github.com/spoke-d/thermionic/internal/db"
	reflect "reflect"
)

// MockTx is a mock of Tx interface
type MockTx struct {
	ctrl     *gomock.Controller
	recorder *MockTxMockRecorder
}

// MockTxMockRecorder is the mock recorder for MockTx
type MockTxMockRecorder struct {
	mock *MockTx
}

// NewMockTx creates a new mock instance
func NewMockTx(ctrl *gomock.Controller) *MockTx {
	mock := &MockTx{ctrl: ctrl}
	mock.recorder = &MockTxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTx) EXPECT() *MockTxMockRecorder {
	return m.recorder
}

// Config mocks base method
func (m *MockTx) Config() (map[string]string, error) {
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Config indicates an expected call of Config
func (mr *MockTxMockRecorder) Config() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockTx)(nil).Config))
}

// RaftNodes mocks base method
func (m *MockTx) RaftNodes() ([]db.RaftNode, error) {
	ret := m.ctrl.Call(m, "RaftNodes")
	ret0, _ := ret[0].([]db.RaftNode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RaftNodes indicates an expected call of RaftNodes
func (mr *MockTxMockRecorder) RaftNodes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RaftNodes", reflect.TypeOf((*MockTx)(nil).RaftNodes))
}

// UpdateConfig mocks base method
func (m *MockTx) UpdateConfig(arg0 map[string]string) error {
	ret := m.ctrl.Call(m, "UpdateConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfig indicates an expected call of UpdateConfig
func (mr *MockTxMockRecorder) UpdateConfig(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfig", reflect.TypeOf((*MockTx)(nil).UpdateConfig), arg0)
}
