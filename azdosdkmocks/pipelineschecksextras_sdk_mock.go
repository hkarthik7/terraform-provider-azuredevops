// Code generated by MockGen. DO NOT EDIT.
// Source: D:/workspace/GolandProjects/terraform-provider-azuredevops/azuredevops/utils/sdk/pipelineschecksextras (interfaces: Client)

// Package azdosdkmocks is a generated GoMock package.
package azdosdkmocks

import (
	context "context"
	reflect "reflect"

	pipelineschecksextras "github.com/microsoft/terraform-provider-azuredevops/azuredevops/utils/sdk/pipelineschecksextras"
	gomock "go.uber.org/mock/gomock"
)

// MockPipelineschecksextrasClient is a mock of Client interface.
type MockPipelineschecksextrasClient struct {
	ctrl     *gomock.Controller
	recorder *MockPipelineschecksextrasClientMockRecorder
	isgomock struct{}
}

// MockPipelineschecksextrasClientMockRecorder is the mock recorder for MockPipelineschecksextrasClient.
type MockPipelineschecksextrasClientMockRecorder struct {
	mock *MockPipelineschecksextrasClient
}

// NewMockPipelineschecksextrasClient creates a new mock instance.
func NewMockPipelineschecksextrasClient(ctrl *gomock.Controller) *MockPipelineschecksextrasClient {
	mock := &MockPipelineschecksextrasClient{ctrl: ctrl}
	mock.recorder = &MockPipelineschecksextrasClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelineschecksextrasClient) EXPECT() *MockPipelineschecksextrasClientMockRecorder {
	return m.recorder
}

// AddCheckConfiguration mocks base method.
func (m *MockPipelineschecksextrasClient) AddCheckConfiguration(arg0 context.Context, arg1 pipelineschecksextras.AddCheckConfigurationArgs) (*pipelineschecksextras.CheckConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCheckConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*pipelineschecksextras.CheckConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCheckConfiguration indicates an expected call of AddCheckConfiguration.
func (mr *MockPipelineschecksextrasClientMockRecorder) AddCheckConfiguration(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCheckConfiguration", reflect.TypeOf((*MockPipelineschecksextrasClient)(nil).AddCheckConfiguration), arg0, arg1)
}

// DeleteCheckConfiguration mocks base method.
func (m *MockPipelineschecksextrasClient) DeleteCheckConfiguration(arg0 context.Context, arg1 pipelineschecksextras.DeleteCheckConfigurationArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCheckConfiguration", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCheckConfiguration indicates an expected call of DeleteCheckConfiguration.
func (mr *MockPipelineschecksextrasClientMockRecorder) DeleteCheckConfiguration(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCheckConfiguration", reflect.TypeOf((*MockPipelineschecksextrasClient)(nil).DeleteCheckConfiguration), arg0, arg1)
}

// EvaluateCheckSuite mocks base method.
func (m *MockPipelineschecksextrasClient) EvaluateCheckSuite(arg0 context.Context, arg1 pipelineschecksextras.EvaluateCheckSuiteArgs) (*pipelineschecksextras.CheckSuite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EvaluateCheckSuite", arg0, arg1)
	ret0, _ := ret[0].(*pipelineschecksextras.CheckSuite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EvaluateCheckSuite indicates an expected call of EvaluateCheckSuite.
func (mr *MockPipelineschecksextrasClientMockRecorder) EvaluateCheckSuite(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvaluateCheckSuite", reflect.TypeOf((*MockPipelineschecksextrasClient)(nil).EvaluateCheckSuite), arg0, arg1)
}

// GetCheckConfiguration mocks base method.
func (m *MockPipelineschecksextrasClient) GetCheckConfiguration(arg0 context.Context, arg1 pipelineschecksextras.GetCheckConfigurationArgs) (*pipelineschecksextras.CheckConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*pipelineschecksextras.CheckConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckConfiguration indicates an expected call of GetCheckConfiguration.
func (mr *MockPipelineschecksextrasClientMockRecorder) GetCheckConfiguration(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckConfiguration", reflect.TypeOf((*MockPipelineschecksextrasClient)(nil).GetCheckConfiguration), arg0, arg1)
}

// GetCheckConfigurationsOnResource mocks base method.
func (m *MockPipelineschecksextrasClient) GetCheckConfigurationsOnResource(arg0 context.Context, arg1 pipelineschecksextras.GetCheckConfigurationsOnResourceArgs) (*[]pipelineschecksextras.CheckConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckConfigurationsOnResource", arg0, arg1)
	ret0, _ := ret[0].(*[]pipelineschecksextras.CheckConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckConfigurationsOnResource indicates an expected call of GetCheckConfigurationsOnResource.
func (mr *MockPipelineschecksextrasClientMockRecorder) GetCheckConfigurationsOnResource(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckConfigurationsOnResource", reflect.TypeOf((*MockPipelineschecksextrasClient)(nil).GetCheckConfigurationsOnResource), arg0, arg1)
}

// GetCheckSuite mocks base method.
func (m *MockPipelineschecksextrasClient) GetCheckSuite(arg0 context.Context, arg1 pipelineschecksextras.GetCheckSuiteArgs) (*pipelineschecksextras.CheckSuite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckSuite", arg0, arg1)
	ret0, _ := ret[0].(*pipelineschecksextras.CheckSuite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckSuite indicates an expected call of GetCheckSuite.
func (mr *MockPipelineschecksextrasClientMockRecorder) GetCheckSuite(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckSuite", reflect.TypeOf((*MockPipelineschecksextrasClient)(nil).GetCheckSuite), arg0, arg1)
}

// QueryCheckConfigurationsOnResources mocks base method.
func (m *MockPipelineschecksextrasClient) QueryCheckConfigurationsOnResources(arg0 context.Context, arg1 pipelineschecksextras.QueryCheckConfigurationsOnResourcesArgs) (*[]pipelineschecksextras.CheckConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryCheckConfigurationsOnResources", arg0, arg1)
	ret0, _ := ret[0].(*[]pipelineschecksextras.CheckConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryCheckConfigurationsOnResources indicates an expected call of QueryCheckConfigurationsOnResources.
func (mr *MockPipelineschecksextrasClientMockRecorder) QueryCheckConfigurationsOnResources(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryCheckConfigurationsOnResources", reflect.TypeOf((*MockPipelineschecksextrasClient)(nil).QueryCheckConfigurationsOnResources), arg0, arg1)
}

// UpdateCheckConfiguration mocks base method.
func (m *MockPipelineschecksextrasClient) UpdateCheckConfiguration(arg0 context.Context, arg1 pipelineschecksextras.UpdateCheckConfigurationArgs) (*pipelineschecksextras.CheckConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCheckConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*pipelineschecksextras.CheckConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCheckConfiguration indicates an expected call of UpdateCheckConfiguration.
func (mr *MockPipelineschecksextrasClientMockRecorder) UpdateCheckConfiguration(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCheckConfiguration", reflect.TypeOf((*MockPipelineschecksextrasClient)(nil).UpdateCheckConfiguration), arg0, arg1)
}
