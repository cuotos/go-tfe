// Code generated by MockGen. DO NOT EDIT.
// Source: team_project_access.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tfe "github.com/hashicorp/go-tfe"
)

// MockTeamProjectAccesses is a mock of TeamProjectAccesses interface.
type MockTeamProjectAccesses struct {
	ctrl     *gomock.Controller
	recorder *MockTeamProjectAccessesMockRecorder
}

// MockTeamProjectAccessesMockRecorder is the mock recorder for MockTeamProjectAccesses.
type MockTeamProjectAccessesMockRecorder struct {
	mock *MockTeamProjectAccesses
}

// NewMockTeamProjectAccesses creates a new mock instance.
func NewMockTeamProjectAccesses(ctrl *gomock.Controller) *MockTeamProjectAccesses {
	mock := &MockTeamProjectAccesses{ctrl: ctrl}
	mock.recorder = &MockTeamProjectAccessesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamProjectAccesses) EXPECT() *MockTeamProjectAccessesMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockTeamProjectAccesses) Add(ctx context.Context, options tfe.TeamProjectAccessAddOptions) (*tfe.TeamProjectAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, options)
	ret0, _ := ret[0].(*tfe.TeamProjectAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockTeamProjectAccessesMockRecorder) Add(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTeamProjectAccesses)(nil).Add), ctx, options)
}

// List mocks base method.
func (m *MockTeamProjectAccesses) List(ctx context.Context, options tfe.TeamProjectAccessListOptions) (*tfe.TeamProjectAccessList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, options)
	ret0, _ := ret[0].(*tfe.TeamProjectAccessList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTeamProjectAccessesMockRecorder) List(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTeamProjectAccesses)(nil).List), ctx, options)
}

// Read mocks base method.
func (m *MockTeamProjectAccesses) Read(ctx context.Context, teamProjectAccessID string) (*tfe.TeamProjectAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, teamProjectAccessID)
	ret0, _ := ret[0].(*tfe.TeamProjectAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockTeamProjectAccessesMockRecorder) Read(ctx, teamProjectAccessID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockTeamProjectAccesses)(nil).Read), ctx, teamProjectAccessID)
}

// Remove mocks base method.
func (m *MockTeamProjectAccesses) Remove(ctx context.Context, teamProjectAccessID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, teamProjectAccessID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockTeamProjectAccessesMockRecorder) Remove(ctx, teamProjectAccessID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockTeamProjectAccesses)(nil).Remove), ctx, teamProjectAccessID)
}

// Update mocks base method.
func (m *MockTeamProjectAccesses) Update(ctx context.Context, teamProjectAccessID string, options tfe.TeamProjectAccessUpdateOptions) (*tfe.TeamProjectAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, teamProjectAccessID, options)
	ret0, _ := ret[0].(*tfe.TeamProjectAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTeamProjectAccessesMockRecorder) Update(ctx, teamProjectAccessID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTeamProjectAccesses)(nil).Update), ctx, teamProjectAccessID, options)
}