// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spoke-d/thermionic/internal/cluster/heartbeat (interfaces: Node)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	db "github.com/spoke-d/thermionic/internal/db"
	reflect "reflect"
)

// MockNode is a mock of Node interface
type MockNode struct {
	ctrl     *gomock.Controller
	recorder *MockNodeMockRecorder
}

// MockNodeMockRecorder is the mock recorder for MockNode
type MockNodeMockRecorder struct {
	mock *MockNode
}

// NewMockNode creates a new mock instance
func NewMockNode(ctrl *gomock.Controller) *MockNode {
	mock := &MockNode{ctrl: ctrl}
	mock.recorder = &MockNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNode) EXPECT() *MockNodeMockRecorder {
	return m.recorder
}

// Transaction mocks base method
func (m *MockNode) Transaction(arg0 func(*db.NodeTx) error) error {
	ret := m.ctrl.Call(m, "Transaction", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transaction indicates an expected call of Transaction
func (mr *MockNodeMockRecorder) Transaction(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockNode)(nil).Transaction), arg0)
}
