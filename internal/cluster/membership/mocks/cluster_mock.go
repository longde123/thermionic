// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spoke-d/thermionic/internal/cluster/membership (interfaces: Cluster)

// Package mocks is a generated GoMock package.
package mocks

import (
	go_dqlite "github.com/CanonicalLtd/go-dqlite"
	db "github.com/spoke-d/thermionic/internal/db"
	cluster "github.com/spoke-d/thermionic/internal/db/cluster"
	database "github.com/spoke-d/thermionic/internal/db/database"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
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

// Close mocks base method
func (m *MockCluster) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockClusterMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCluster)(nil).Close))
}

// DB mocks base method
func (m *MockCluster) DB() database.DB {
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(database.DB)
	return ret0
}

// DB indicates an expected call of DB
func (mr *MockClusterMockRecorder) DB() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockCluster)(nil).DB))
}

// EnterExclusive mocks base method
func (m *MockCluster) EnterExclusive() error {
	ret := m.ctrl.Call(m, "EnterExclusive")
	ret0, _ := ret[0].(error)
	return ret0
}

// EnterExclusive indicates an expected call of EnterExclusive
func (mr *MockClusterMockRecorder) EnterExclusive() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnterExclusive", reflect.TypeOf((*MockCluster)(nil).EnterExclusive))
}

// ExitExclusive mocks base method
func (m *MockCluster) ExitExclusive(arg0 func(*db.ClusterTx) error) error {
	ret := m.ctrl.Call(m, "ExitExclusive", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExitExclusive indicates an expected call of ExitExclusive
func (mr *MockClusterMockRecorder) ExitExclusive(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExitExclusive", reflect.TypeOf((*MockCluster)(nil).ExitExclusive), arg0)
}

// NodeID mocks base method
func (m *MockCluster) NodeID(arg0 int64) {
	m.ctrl.Call(m, "NodeID", arg0)
}

// NodeID indicates an expected call of NodeID
func (mr *MockClusterMockRecorder) NodeID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeID", reflect.TypeOf((*MockCluster)(nil).NodeID), arg0)
}

// Open mocks base method
func (m *MockCluster) Open(arg0 cluster.ServerStore, arg1, arg2 string, arg3 time.Duration, arg4 ...go_dqlite.DriverOption) error {
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Open", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open
func (mr *MockClusterMockRecorder) Open(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockCluster)(nil).Open), varargs...)
}

// SchemaVersion mocks base method
func (m *MockCluster) SchemaVersion() int {
	ret := m.ctrl.Call(m, "SchemaVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// SchemaVersion indicates an expected call of SchemaVersion
func (mr *MockClusterMockRecorder) SchemaVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaVersion", reflect.TypeOf((*MockCluster)(nil).SchemaVersion))
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
