// Code generated by MockGen. DO NOT EDIT.
// Source: Domain/promote.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
)

// MockPromoteUseCase is a mock of PromoteUseCase interface.
type MockPromoteUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockPromoteUseCaseMockRecorder
}

// MockPromoteUseCaseMockRecorder is the mock recorder for MockPromoteUseCase.
type MockPromoteUseCaseMockRecorder struct {
	mock *MockPromoteUseCase
}

// NewMockPromoteUseCase creates a new mock instance.
func NewMockPromoteUseCase(ctrl *gomock.Controller) *MockPromoteUseCase {
	mock := &MockPromoteUseCase{ctrl: ctrl}
	mock.recorder = &MockPromoteUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPromoteUseCase) EXPECT() *MockPromoteUseCaseMockRecorder {
	return m.recorder
}

// PromoteUser mocks base method.
func (m *MockPromoteUseCase) PromoteUser(userID string, ctx context.Context) *domain.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PromoteUser", userID, ctx)
	ret0, _ := ret[0].(*domain.ErrorResponse)
	return ret0
}

// PromoteUser indicates an expected call of PromoteUser.
func (mr *MockPromoteUseCaseMockRecorder) PromoteUser(userID, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PromoteUser", reflect.TypeOf((*MockPromoteUseCase)(nil).PromoteUser), userID, ctx)
}