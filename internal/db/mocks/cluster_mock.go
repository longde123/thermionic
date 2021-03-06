// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spoke-d/thermionic/internal/db (interfaces: ClusterTxProvider)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	db "github.com/spoke-d/thermionic/internal/db"
	database "github.com/spoke-d/thermionic/internal/db/database"
	reflect "reflect"
)

// MockClusterTxProvider is a mock of ClusterTxProvider interface
type MockClusterTxProvider struct {
	ctrl     *gomock.Controller
	recorder *MockClusterTxProviderMockRecorder
}

// MockClusterTxProviderMockRecorder is the mock recorder for MockClusterTxProvider
type MockClusterTxProviderMockRecorder struct {
	mock *MockClusterTxProvider
}

// NewMockClusterTxProvider creates a new mock instance
func NewMockClusterTxProvider(ctrl *gomock.Controller) *MockClusterTxProvider {
	mock := &MockClusterTxProvider{ctrl: ctrl}
	mock.recorder = &MockClusterTxProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterTxProvider) EXPECT() *MockClusterTxProviderMockRecorder {
	return m.recorder
}

// New mocks base method
func (m *MockClusterTxProvider) New(arg0 database.Tx, arg1 int64) *db.ClusterTx {
	ret := m.ctrl.Call(m, "New", arg0, arg1)
	ret0, _ := ret[0].(*db.ClusterTx)
	return ret0
}

// New indicates an expected call of New
func (mr *MockClusterTxProviderMockRecorder) New(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockClusterTxProvider)(nil).New), arg0, arg1)
}
