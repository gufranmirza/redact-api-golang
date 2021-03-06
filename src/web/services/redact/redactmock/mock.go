// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gufranmirza/redact-api-golang/src/web/services/redact (interfaces: Redact)

// Package redactmock is a generated GoMock package.
package redactmock

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockRedact is a mock of Redact interface
type MockRedact struct {
	ctrl     *gomock.Controller
	recorder *MockRedactMockRecorder
}

// MockRedactMockRecorder is the mock recorder for MockRedact
type MockRedactMockRecorder struct {
	mock *MockRedact
}

// NewMockRedact creates a new mock instance
func NewMockRedact(ctrl *gomock.Controller) *MockRedact {
	mock := &MockRedact{ctrl: ctrl}
	mock.recorder = &MockRedactMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRedact) EXPECT() *MockRedactMockRecorder {
	return m.recorder
}

// RedactJSON mocks base method
func (m *MockRedact) RedactJSON() http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RedactJSON")
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// RedactJSON indicates an expected call of RedactJSON
func (mr *MockRedactMockRecorder) RedactJSON() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedactJSON", reflect.TypeOf((*MockRedact)(nil).RedactJSON))
}
