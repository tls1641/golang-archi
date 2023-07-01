// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package architecture is a generated GoMock package.
package architecture

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAccessor is a mock of Accessor interface.
type MockAccessor struct {
	ctrl     *gomock.Controller
	recorder *MockAccessorMockRecorder
}

// MockAccessorMockRecorder is the mock recorder for MockAccessor.
type MockAccessorMockRecorder struct {
	mock *MockAccessor
}

// NewMockAccessor creates a new mock instance.
func NewMockAccessor(ctrl *gomock.Controller) *MockAccessor {
	mock := &MockAccessor{ctrl: ctrl}
	mock.recorder = &MockAccessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessor) EXPECT() *MockAccessorMockRecorder {
	return m.recorder
}

// Retrieve mocks base method.
func (m *MockAccessor) Retrieve(n int) Person {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Retrieve", n)
	ret0, _ := ret[0].(Person)
	return ret0
}

// Retrieve indicates an expected call of Retrieve.
func (mr *MockAccessorMockRecorder) Retrieve(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retrieve", reflect.TypeOf((*MockAccessor)(nil).Retrieve), n)
}

// Save mocks base method.
func (m *MockAccessor) Save(n int, p Person) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Save", n, p)
}

// Save indicates an expected call of Save.
func (mr *MockAccessorMockRecorder) Save(n, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockAccessor)(nil).Save), n, p)
}