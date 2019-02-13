// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spoke-d/thermionic/internal/cluster/raft (interfaces: SnapshotStoreProvider)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	raft "github.com/hashicorp/raft"
	reflect "reflect"
)

// MockSnapshotStoreProvider is a mock of SnapshotStoreProvider interface
type MockSnapshotStoreProvider struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotStoreProviderMockRecorder
}

// MockSnapshotStoreProviderMockRecorder is the mock recorder for MockSnapshotStoreProvider
type MockSnapshotStoreProviderMockRecorder struct {
	mock *MockSnapshotStoreProvider
}

// NewMockSnapshotStoreProvider creates a new mock instance
func NewMockSnapshotStoreProvider(ctrl *gomock.Controller) *MockSnapshotStoreProvider {
	mock := &MockSnapshotStoreProvider{ctrl: ctrl}
	mock.recorder = &MockSnapshotStoreProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSnapshotStoreProvider) EXPECT() *MockSnapshotStoreProviderMockRecorder {
	return m.recorder
}

// Store mocks base method
func (m *MockSnapshotStoreProvider) Store(arg0 string) (raft.SnapshotStore, error) {
	ret := m.ctrl.Call(m, "Store", arg0)
	ret0, _ := ret[0].(raft.SnapshotStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store
func (mr *MockSnapshotStoreProviderMockRecorder) Store(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockSnapshotStoreProvider)(nil).Store), arg0)
}
