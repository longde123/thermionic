// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spoke-d/thermionic/internal/db (interfaces: Query,QueryCluster,QueryNode,Transaction)

// Package mocks is a generated GoMock package.
package mocks

import (
	go_dqlite "github.com/CanonicalLtd/go-dqlite"
	gomock "github.com/golang/mock/gomock"
	cluster "github.com/spoke-d/thermionic/internal/db/cluster"
	database "github.com/spoke-d/thermionic/internal/db/database"
	query "github.com/spoke-d/thermionic/internal/db/query"
	schema "github.com/spoke-d/thermionic/internal/db/schema"
	reflect "reflect"
)

// MockQuery is a mock of Query interface
type MockQuery struct {
	ctrl     *gomock.Controller
	recorder *MockQueryMockRecorder
}

// MockQueryMockRecorder is the mock recorder for MockQuery
type MockQueryMockRecorder struct {
	mock *MockQuery
}

// NewMockQuery creates a new mock instance
func NewMockQuery(ctrl *gomock.Controller) *MockQuery {
	mock := &MockQuery{ctrl: ctrl}
	mock.recorder = &MockQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuery) EXPECT() *MockQueryMockRecorder {
	return m.recorder
}

// Count mocks base method
func (m *MockQuery) Count(arg0 database.Tx, arg1, arg2 string, arg3 ...interface{}) (int, error) {
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Count", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockQueryMockRecorder) Count(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockQuery)(nil).Count), varargs...)
}

// DeleteObject mocks base method
func (m *MockQuery) DeleteObject(arg0 database.Tx, arg1 string, arg2 int64) (bool, error) {
	ret := m.ctrl.Call(m, "DeleteObject", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObject indicates an expected call of DeleteObject
func (mr *MockQueryMockRecorder) DeleteObject(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockQuery)(nil).DeleteObject), arg0, arg1, arg2)
}

// SelectConfig mocks base method
func (m *MockQuery) SelectConfig(arg0 database.Tx, arg1, arg2 string, arg3 ...interface{}) (map[string]string, error) {
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SelectConfig", varargs...)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectConfig indicates an expected call of SelectConfig
func (mr *MockQueryMockRecorder) SelectConfig(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectConfig", reflect.TypeOf((*MockQuery)(nil).SelectConfig), varargs...)
}

// SelectObjects mocks base method
func (m *MockQuery) SelectObjects(arg0 database.Tx, arg1 query.Dest, arg2 string, arg3 ...interface{}) error {
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SelectObjects", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectObjects indicates an expected call of SelectObjects
func (mr *MockQueryMockRecorder) SelectObjects(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectObjects", reflect.TypeOf((*MockQuery)(nil).SelectObjects), varargs...)
}

// SelectStrings mocks base method
func (m *MockQuery) SelectStrings(arg0 database.Tx, arg1 string, arg2 ...interface{}) ([]string, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SelectStrings", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectStrings indicates an expected call of SelectStrings
func (mr *MockQueryMockRecorder) SelectStrings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectStrings", reflect.TypeOf((*MockQuery)(nil).SelectStrings), varargs...)
}

// UpdateConfig mocks base method
func (m *MockQuery) UpdateConfig(arg0 database.Tx, arg1 string, arg2 map[string]string) error {
	ret := m.ctrl.Call(m, "UpdateConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfig indicates an expected call of UpdateConfig
func (mr *MockQueryMockRecorder) UpdateConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfig", reflect.TypeOf((*MockQuery)(nil).UpdateConfig), arg0, arg1, arg2)
}

// UpsertObject mocks base method
func (m *MockQuery) UpsertObject(arg0 database.Tx, arg1 string, arg2 []string, arg3 []interface{}) (int64, error) {
	ret := m.ctrl.Call(m, "UpsertObject", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertObject indicates an expected call of UpsertObject
func (mr *MockQueryMockRecorder) UpsertObject(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertObject", reflect.TypeOf((*MockQuery)(nil).UpsertObject), arg0, arg1, arg2, arg3)
}

// MockQueryCluster is a mock of QueryCluster interface
type MockQueryCluster struct {
	ctrl     *gomock.Controller
	recorder *MockQueryClusterMockRecorder
}

// MockQueryClusterMockRecorder is the mock recorder for MockQueryCluster
type MockQueryClusterMockRecorder struct {
	mock *MockQueryCluster
}

// NewMockQueryCluster creates a new mock instance
func NewMockQueryCluster(ctrl *gomock.Controller) *MockQueryCluster {
	mock := &MockQueryCluster{ctrl: ctrl}
	mock.recorder = &MockQueryClusterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQueryCluster) EXPECT() *MockQueryClusterMockRecorder {
	return m.recorder
}

// DB mocks base method
func (m *MockQueryCluster) DB() database.DB {
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(database.DB)
	return ret0
}

// DB indicates an expected call of DB
func (mr *MockQueryClusterMockRecorder) DB() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockQueryCluster)(nil).DB))
}

// EnsureSchema mocks base method
func (m *MockQueryCluster) EnsureSchema(arg0, arg1 string) (bool, error) {
	ret := m.ctrl.Call(m, "EnsureSchema", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureSchema indicates an expected call of EnsureSchema
func (mr *MockQueryClusterMockRecorder) EnsureSchema(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureSchema", reflect.TypeOf((*MockQueryCluster)(nil).EnsureSchema), arg0, arg1)
}

// Open mocks base method
func (m *MockQueryCluster) Open(arg0 cluster.ServerStore, arg1 ...go_dqlite.DriverOption) error {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Open", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open
func (mr *MockQueryClusterMockRecorder) Open(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockQueryCluster)(nil).Open), varargs...)
}

// SchemaVersion mocks base method
func (m *MockQueryCluster) SchemaVersion() int {
	ret := m.ctrl.Call(m, "SchemaVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// SchemaVersion indicates an expected call of SchemaVersion
func (mr *MockQueryClusterMockRecorder) SchemaVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaVersion", reflect.TypeOf((*MockQueryCluster)(nil).SchemaVersion))
}

// MockQueryNode is a mock of QueryNode interface
type MockQueryNode struct {
	ctrl     *gomock.Controller
	recorder *MockQueryNodeMockRecorder
}

// MockQueryNodeMockRecorder is the mock recorder for MockQueryNode
type MockQueryNodeMockRecorder struct {
	mock *MockQueryNode
}

// NewMockQueryNode creates a new mock instance
func NewMockQueryNode(ctrl *gomock.Controller) *MockQueryNode {
	mock := &MockQueryNode{ctrl: ctrl}
	mock.recorder = &MockQueryNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQueryNode) EXPECT() *MockQueryNodeMockRecorder {
	return m.recorder
}

// DB mocks base method
func (m *MockQueryNode) DB() database.DB {
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(database.DB)
	return ret0
}

// DB indicates an expected call of DB
func (mr *MockQueryNodeMockRecorder) DB() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockQueryNode)(nil).DB))
}

// EnsureSchema mocks base method
func (m *MockQueryNode) EnsureSchema(arg0 schema.Hook) (int, error) {
	ret := m.ctrl.Call(m, "EnsureSchema", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureSchema indicates an expected call of EnsureSchema
func (mr *MockQueryNodeMockRecorder) EnsureSchema(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureSchema", reflect.TypeOf((*MockQueryNode)(nil).EnsureSchema), arg0)
}

// Open mocks base method
func (m *MockQueryNode) Open(arg0 string) error {
	ret := m.ctrl.Call(m, "Open", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open
func (mr *MockQueryNodeMockRecorder) Open(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockQueryNode)(nil).Open), arg0)
}

// MockTransaction is a mock of Transaction interface
type MockTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionMockRecorder
}

// MockTransactionMockRecorder is the mock recorder for MockTransaction
type MockTransactionMockRecorder struct {
	mock *MockTransaction
}

// NewMockTransaction creates a new mock instance
func NewMockTransaction(ctrl *gomock.Controller) *MockTransaction {
	mock := &MockTransaction{ctrl: ctrl}
	mock.recorder = &MockTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransaction) EXPECT() *MockTransactionMockRecorder {
	return m.recorder
}

// Transaction mocks base method
func (m *MockTransaction) Transaction(arg0 database.DB, arg1 func(database.Tx) error) error {
	ret := m.ctrl.Call(m, "Transaction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transaction indicates an expected call of Transaction
func (mr *MockTransactionMockRecorder) Transaction(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockTransaction)(nil).Transaction), arg0, arg1)
}
