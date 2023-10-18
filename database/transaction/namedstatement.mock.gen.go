// Code generated by MockGen. DO NOT EDIT.
// Source: namedstatement.go

// Package transaction is a generated GoMock package.
package transaction

import (
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNamedStatement is a mock of NamedStatement interface.
type MockNamedStatement struct {
	ctrl     *gomock.Controller
	recorder *MockNamedStatementMockRecorder
}

// MockNamedStatementMockRecorder is the mock recorder for MockNamedStatement.
type MockNamedStatementMockRecorder struct {
	mock *MockNamedStatement
}

// NewMockNamedStatement creates a new mock instance.
func NewMockNamedStatement(ctrl *gomock.Controller) *MockNamedStatement {
	mock := &MockNamedStatement{ctrl: ctrl}
	mock.recorder = &MockNamedStatementMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamedStatement) EXPECT() *MockNamedStatementMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockNamedStatement) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockNamedStatementMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockNamedStatement)(nil).Close))
}

// Exec mocks base method.
func (m *MockNamedStatement) Exec(arg interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec", arg)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockNamedStatementMockRecorder) Exec(arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockNamedStatement)(nil).Exec), arg)
}
