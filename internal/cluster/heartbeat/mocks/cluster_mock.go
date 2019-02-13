// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spoke-d/thermionic/internal/cluster/heartbeat (interfaces: Cluster)

// Package mocks is a generated GoMock package.
package mocks

import (
	db "github.com/spoke-d/thermionic/internal/db"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCluster is a mock of Cluster interface
type MockCluster struct {
	ctrl     *gomock.Controller
	recorder *MockClusterMockRecorder
}

// MockClusterMockRecorder is the mock recorder for MockCluster
type MockClusterMockRecorder struct {
	mock *MockCluster
}

// NewMockCluster creates a new mock instance
func NewMockCluster(ctrl *gomock.Controller) *MockCluster {
	mock := &MockCluster{ctrl: ctrl}
	mock.recorder = &MockClusterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCluster) EXPECT() *MockClusterMockRecorder {
	return m.recorder
}

// Transaction mocks base method
func (m *MockCluster) Transaction(arg0 func(*db.ClusterTx) error) error {
	ret := m.ctrl.Call(m, "Transaction", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transaction indicates an expected call of Transaction
func (mr *MockClusterMockRecorder) Transaction(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockCluster)(nil).Transaction), arg0)
}
