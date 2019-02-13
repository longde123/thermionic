// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spoke-d/thermionic/internal/cluster (interfaces: ServerProvider)

// Package mocks is a generated GoMock package.
package mocks

import (
	cluster "github.com/spoke-d/thermionic/internal/cluster"
	raft "github.com/spoke-d/thermionic/internal/cluster/raft"
	log "github.com/go-kit/kit/log"
	gomock "github.com/golang/mock/gomock"
	net "net"
	reflect "reflect"
)

// MockServerProvider is a mock of ServerProvider interface
type MockServerProvider struct {
	ctrl     *gomock.Controller
	recorder *MockServerProviderMockRecorder
}

// MockServerProviderMockRecorder is the mock recorder for MockServerProvider
type MockServerProviderMockRecorder struct {
	mock *MockServerProvider
}

// NewMockServerProvider creates a new mock instance
func NewMockServerProvider(ctrl *gomock.Controller) *MockServerProvider {
	mock := &MockServerProvider{ctrl: ctrl}
	mock.recorder = &MockServerProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerProvider) EXPECT() *MockServerProviderMockRecorder {
	return m.recorder
}

// New mocks base method
func (m *MockServerProvider) New(arg0 cluster.RaftInstance, arg1 net.Listener, arg2 *raft.AddressProvider, arg3 log.Logger) (cluster.Server, error) {
	ret := m.ctrl.Call(m, "New", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cluster.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New
func (mr *MockServerProviderMockRecorder) New(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockServerProvider)(nil).New), arg0, arg1, arg2, arg3)
}
