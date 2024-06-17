// Code generated by MockGen. DO NOT EDIT.
// Source: internal/commands/builder_inspect.go

// Package mock_commands is a generated GoMock package.
package testmocks

import (
	reflect "reflect"

	client "github.com/buildpacks/pack/pkg/client"

	gomock "github.com/golang/mock/gomock"
)

// MockBuilderInspector is a mock of BuilderInspector interface.
type MockBuilderInspector struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderInspectorMockRecorder
}

// MockBuilderInspectorMockRecorder is the mock recorder for MockBuilderInspector.
type MockBuilderInspectorMockRecorder struct {
	mock *MockBuilderInspector
}

// NewMockBuilderInspector creates a new mock instance.
func NewMockBuilderInspector(ctrl *gomock.Controller) *MockBuilderInspector {
	mock := &MockBuilderInspector{ctrl: ctrl}
	mock.recorder = &MockBuilderInspectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuilderInspector) EXPECT() *MockBuilderInspectorMockRecorder {
	return m.recorder
}

// InspectBuilder mocks base method.
func (m *MockBuilderInspector) InspectBuilder(name string, daemon bool, modifiers ...client.BuilderInspectionModifier) (*client.BuilderInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, daemon}
	for _, a := range modifiers {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InspectBuilder", varargs...)
	ret0, _ := ret[0].(*client.BuilderInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectBuilder indicates an expected call of InspectBuilder.
func (mr *MockBuilderInspectorMockRecorder) InspectBuilder(name, daemon interface{}, modifiers ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, daemon}, modifiers...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectBuilder", reflect.TypeOf((*MockBuilderInspector)(nil).InspectBuilder), varargs...)
}
