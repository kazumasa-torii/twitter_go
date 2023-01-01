// Code generated by MockGen. DO NOT EDIT.
// Source: token.go

// Package mock_token is a generated GoMock package.
package mock_token

import (
	domain "api/server/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTokenizer is a mock of Tokenizer interface.
type MockTokenizer struct {
	ctrl     *gomock.Controller
	recorder *MockTokenizerMockRecorder
}

// MockTokenizerMockRecorder is the mock recorder for MockTokenizer.
type MockTokenizerMockRecorder struct {
	mock *MockTokenizer
}

// NewMockTokenizer creates a new mock instance.
func NewMockTokenizer(ctrl *gomock.Controller) *MockTokenizer {
	mock := &MockTokenizer{ctrl: ctrl}
	mock.recorder = &MockTokenizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenizer) EXPECT() *MockTokenizerMockRecorder {
	return m.recorder
}

// New mocks base method.
func (m *MockTokenizer) New(arg0 domain.User) (domain.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", arg0)
	ret0, _ := ret[0].(domain.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New.
func (mr *MockTokenizerMockRecorder) New(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockTokenizer)(nil).New), arg0)
}

// Verify mocks base method.
func (m *MockTokenizer) Verify(arg0 domain.Token) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockTokenizerMockRecorder) Verify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockTokenizer)(nil).Verify), arg0)
}
