// Code generated by MockGen. DO NOT EDIT.
// Source: ./store/store.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(name string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", name)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), name)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(id int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", id)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), id)
}