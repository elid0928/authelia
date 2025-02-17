// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: AccessTokenStrategy)
//
// Generated by this command:
//
//	mockgen -package mocks -destination fosite_access_token_strategy.go -mock_names Provider=MockAccessTokenStrategy github.com/ory/fosite/handler/oauth2 AccessTokenStrategy
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	fosite "github.com/ory/fosite"
	gomock "go.uber.org/mock/gomock"
)

// MockAccessTokenStrategy is a mock of AccessTokenStrategy interface.
type MockAccessTokenStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockAccessTokenStrategyMockRecorder
}

// MockAccessTokenStrategyMockRecorder is the mock recorder for MockAccessTokenStrategy.
type MockAccessTokenStrategyMockRecorder struct {
	mock *MockAccessTokenStrategy
}

// NewMockAccessTokenStrategy creates a new mock instance.
func NewMockAccessTokenStrategy(ctrl *gomock.Controller) *MockAccessTokenStrategy {
	mock := &MockAccessTokenStrategy{ctrl: ctrl}
	mock.recorder = &MockAccessTokenStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessTokenStrategy) EXPECT() *MockAccessTokenStrategyMockRecorder {
	return m.recorder
}

// AccessTokenSignature mocks base method.
func (m *MockAccessTokenStrategy) AccessTokenSignature(arg0 context.Context, arg1 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessTokenSignature", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// AccessTokenSignature indicates an expected call of AccessTokenSignature.
func (mr *MockAccessTokenStrategyMockRecorder) AccessTokenSignature(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessTokenSignature", reflect.TypeOf((*MockAccessTokenStrategy)(nil).AccessTokenSignature), arg0, arg1)
}

// GenerateAccessToken mocks base method.
func (m *MockAccessTokenStrategy) GenerateAccessToken(arg0 context.Context, arg1 fosite.Requester) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockAccessTokenStrategyMockRecorder) GenerateAccessToken(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockAccessTokenStrategy)(nil).GenerateAccessToken), arg0, arg1)
}

// ValidateAccessToken mocks base method.
func (m *MockAccessTokenStrategy) ValidateAccessToken(arg0 context.Context, arg1 fosite.Requester, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAccessToken", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateAccessToken indicates an expected call of ValidateAccessToken.
func (mr *MockAccessTokenStrategyMockRecorder) ValidateAccessToken(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAccessToken", reflect.TypeOf((*MockAccessTokenStrategy)(nil).ValidateAccessToken), arg0, arg1, arg2)
}
