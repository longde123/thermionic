// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spoke-d/thermionic/internal/discovery (interfaces: EventBus,Members,MemberList,Member)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	discovery "github.com/spoke-d/thermionic/internal/discovery"
	members "github.com/spoke-d/thermionic/internal/discovery/members"
	reflect "reflect"
)

// MockEventBus is a mock of EventBus interface
type MockEventBus struct {
	ctrl     *gomock.Controller
	recorder *MockEventBusMockRecorder
}

// MockEventBusMockRecorder is the mock recorder for MockEventBus
type MockEventBusMockRecorder struct {
	mock *MockEventBus
}

// NewMockEventBus creates a new mock instance
func NewMockEventBus(ctrl *gomock.Controller) *MockEventBus {
	mock := &MockEventBus{ctrl: ctrl}
	mock.recorder = &MockEventBusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEventBus) EXPECT() *MockEventBusMockRecorder {
	return m.recorder
}

// AddHandler mocks base method
func (m *MockEventBus) AddHandler(arg0 members.Handler) {
	m.ctrl.Call(m, "AddHandler", arg0)
}

// AddHandler indicates an expected call of AddHandler
func (mr *MockEventBusMockRecorder) AddHandler(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHandler", reflect.TypeOf((*MockEventBus)(nil).AddHandler), arg0)
}

// DispatchEvent mocks base method
func (m *MockEventBus) DispatchEvent(arg0 members.Event) error {
	ret := m.ctrl.Call(m, "DispatchEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DispatchEvent indicates an expected call of DispatchEvent
func (mr *MockEventBusMockRecorder) DispatchEvent(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchEvent", reflect.TypeOf((*MockEventBus)(nil).DispatchEvent), arg0)
}

// RemoveHandler mocks base method
func (m *MockEventBus) RemoveHandler(arg0 members.Handler) {
	m.ctrl.Call(m, "RemoveHandler", arg0)
}

// RemoveHandler indicates an expected call of RemoveHandler
func (mr *MockEventBusMockRecorder) RemoveHandler(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveHandler", reflect.TypeOf((*MockEventBus)(nil).RemoveHandler), arg0)
}

// MockMembers is a mock of Members interface
type MockMembers struct {
	ctrl     *gomock.Controller
	recorder *MockMembersMockRecorder
}

// MockMembersMockRecorder is the mock recorder for MockMembers
type MockMembersMockRecorder struct {
	mock *MockMembers
}

// NewMockMembers creates a new mock instance
func NewMockMembers(ctrl *gomock.Controller) *MockMembers {
	mock := &MockMembers{ctrl: ctrl}
	mock.recorder = &MockMembersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMembers) EXPECT() *MockMembersMockRecorder {
	return m.recorder
}

// AddHandler mocks base method
func (m *MockMembers) AddHandler(arg0 members.Handler) {
	m.ctrl.Call(m, "AddHandler", arg0)
}

// AddHandler indicates an expected call of AddHandler
func (mr *MockMembersMockRecorder) AddHandler(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHandler", reflect.TypeOf((*MockMembers)(nil).AddHandler), arg0)
}

// DispatchEvent mocks base method
func (m *MockMembers) DispatchEvent(arg0 members.Event) error {
	ret := m.ctrl.Call(m, "DispatchEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DispatchEvent indicates an expected call of DispatchEvent
func (mr *MockMembersMockRecorder) DispatchEvent(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchEvent", reflect.TypeOf((*MockMembers)(nil).DispatchEvent), arg0)
}

// Join mocks base method
func (m *MockMembers) Join() (int, error) {
	ret := m.ctrl.Call(m, "Join")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Join indicates an expected call of Join
func (mr *MockMembersMockRecorder) Join() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*MockMembers)(nil).Join))
}

// Leave mocks base method
func (m *MockMembers) Leave() error {
	ret := m.ctrl.Call(m, "Leave")
	ret0, _ := ret[0].(error)
	return ret0
}

// Leave indicates an expected call of Leave
func (mr *MockMembersMockRecorder) Leave() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leave", reflect.TypeOf((*MockMembers)(nil).Leave))
}

// MemberList mocks base method
func (m *MockMembers) MemberList() discovery.MemberList {
	ret := m.ctrl.Call(m, "MemberList")
	ret0, _ := ret[0].(discovery.MemberList)
	return ret0
}

// MemberList indicates an expected call of MemberList
func (mr *MockMembersMockRecorder) MemberList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberList", reflect.TypeOf((*MockMembers)(nil).MemberList))
}

// RemoveHandler mocks base method
func (m *MockMembers) RemoveHandler(arg0 members.Handler) {
	m.ctrl.Call(m, "RemoveHandler", arg0)
}

// RemoveHandler indicates an expected call of RemoveHandler
func (mr *MockMembersMockRecorder) RemoveHandler(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveHandler", reflect.TypeOf((*MockMembers)(nil).RemoveHandler), arg0)
}

// Shutdown mocks base method
func (m *MockMembers) Shutdown() error {
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockMembersMockRecorder) Shutdown() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockMembers)(nil).Shutdown))
}

// Walk mocks base method
func (m *MockMembers) Walk(arg0 func(members.PeerInfo) error) error {
	ret := m.ctrl.Call(m, "Walk", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Walk indicates an expected call of Walk
func (mr *MockMembersMockRecorder) Walk(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Walk", reflect.TypeOf((*MockMembers)(nil).Walk), arg0)
}

// MockMemberList is a mock of MemberList interface
type MockMemberList struct {
	ctrl     *gomock.Controller
	recorder *MockMemberListMockRecorder
}

// MockMemberListMockRecorder is the mock recorder for MockMemberList
type MockMemberListMockRecorder struct {
	mock *MockMemberList
}

// NewMockMemberList creates a new mock instance
func NewMockMemberList(ctrl *gomock.Controller) *MockMemberList {
	mock := &MockMemberList{ctrl: ctrl}
	mock.recorder = &MockMemberListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMemberList) EXPECT() *MockMemberListMockRecorder {
	return m.recorder
}

// LocalNode mocks base method
func (m *MockMemberList) LocalNode() discovery.Member {
	ret := m.ctrl.Call(m, "LocalNode")
	ret0, _ := ret[0].(discovery.Member)
	return ret0
}

// LocalNode indicates an expected call of LocalNode
func (mr *MockMemberListMockRecorder) LocalNode() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalNode", reflect.TypeOf((*MockMemberList)(nil).LocalNode))
}

// Members mocks base method
func (m *MockMemberList) Members() []discovery.Member {
	ret := m.ctrl.Call(m, "Members")
	ret0, _ := ret[0].([]discovery.Member)
	return ret0
}

// Members indicates an expected call of Members
func (mr *MockMemberListMockRecorder) Members() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Members", reflect.TypeOf((*MockMemberList)(nil).Members))
}

// NumMembers mocks base method
func (m *MockMemberList) NumMembers() int {
	ret := m.ctrl.Call(m, "NumMembers")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumMembers indicates an expected call of NumMembers
func (mr *MockMemberListMockRecorder) NumMembers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumMembers", reflect.TypeOf((*MockMemberList)(nil).NumMembers))
}

// MockMember is a mock of Member interface
type MockMember struct {
	ctrl     *gomock.Controller
	recorder *MockMemberMockRecorder
}

// MockMemberMockRecorder is the mock recorder for MockMember
type MockMemberMockRecorder struct {
	mock *MockMember
}

// NewMockMember creates a new mock instance
func NewMockMember(ctrl *gomock.Controller) *MockMember {
	mock := &MockMember{ctrl: ctrl}
	mock.recorder = &MockMemberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMember) EXPECT() *MockMemberMockRecorder {
	return m.recorder
}

// Address mocks base method
func (m *MockMember) Address() string {
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(string)
	return ret0
}

// Address indicates an expected call of Address
func (mr *MockMemberMockRecorder) Address() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockMember)(nil).Address))
}

// Name mocks base method
func (m *MockMember) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockMemberMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockMember)(nil).Name))
}
