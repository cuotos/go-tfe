// Code generated by MockGen. DO NOT EDIT.
// Source: oauth_token.go

// Package tfe is a generated GoMock package.
package tfe

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOAuthTokens is a mock of OAuthTokens interface.
type MockOAuthTokens struct {
	ctrl     *gomock.Controller
	recorder *MockOAuthTokensMockRecorder
}

// MockOAuthTokensMockRecorder is the mock recorder for MockOAuthTokens.
type MockOAuthTokensMockRecorder struct {
	mock *MockOAuthTokens
}

// NewMockOAuthTokens creates a new mock instance.
func NewMockOAuthTokens(ctrl *gomock.Controller) *MockOAuthTokens {
	mock := &MockOAuthTokens{ctrl: ctrl}
	mock.recorder = &MockOAuthTokensMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOAuthTokens) EXPECT() *MockOAuthTokensMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockOAuthTokens) Delete(ctx context.Context, oAuthTokenID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, oAuthTokenID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOAuthTokensMockRecorder) Delete(ctx, oAuthTokenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOAuthTokens)(nil).Delete), ctx, oAuthTokenID)
}

// List mocks base method.
func (m *MockOAuthTokens) List(ctx context.Context, organization string, options OAuthTokenListOptions) (*OAuthTokenList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, organization, options)
	ret0, _ := ret[0].(*OAuthTokenList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockOAuthTokensMockRecorder) List(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOAuthTokens)(nil).List), ctx, organization, options)
}

// Read mocks base method.
func (m *MockOAuthTokens) Read(ctx context.Context, oAuthTokenID string) (*OAuthToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, oAuthTokenID)
	ret0, _ := ret[0].(*OAuthToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockOAuthTokensMockRecorder) Read(ctx, oAuthTokenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockOAuthTokens)(nil).Read), ctx, oAuthTokenID)
}

// Update mocks base method.
func (m *MockOAuthTokens) Update(ctx context.Context, oAuthTokenID string, options OAuthTokenUpdateOptions) (*OAuthToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, oAuthTokenID, options)
	ret0, _ := ret[0].(*OAuthToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOAuthTokensMockRecorder) Update(ctx, oAuthTokenID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOAuthTokens)(nil).Update), ctx, oAuthTokenID, options)
}