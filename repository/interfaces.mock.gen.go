// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetUserByPhone mocks base method.
func (m *MockRepositoryInterface) GetUserByPhone(ctx context.Context, phone string) (User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPhone", ctx, phone)
	ret0, _ := ret[0].(User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPhone indicates an expected call of GetUserByPhone.
func (mr *MockRepositoryInterfaceMockRecorder) GetUserByPhone(ctx, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPhone", reflect.TypeOf((*MockRepositoryInterface)(nil).GetUserByPhone), ctx, phone)
}

// SaveUser mocks base method.
func (m *MockRepositoryInterface) SaveUser(ctx context.Context, user *User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUser indicates an expected call of SaveUser.
func (mr *MockRepositoryInterfaceMockRecorder) SaveUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUser", reflect.TypeOf((*MockRepositoryInterface)(nil).SaveUser), ctx, user)
}