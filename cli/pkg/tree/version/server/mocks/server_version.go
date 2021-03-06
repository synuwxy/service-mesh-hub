// Code generated by MockGen. DO NOT EDIT.
// Source: ./server_version.go

// Package mock_server is a generated GoMock package.
package mock_server

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	server "github.com/solo-io/service-mesh-hub/cli/pkg/tree/version/server"
	v1 "k8s.io/api/apps/v1"
)

// MockDeploymentClient is a mock of DeploymentClient interface.
type MockDeploymentClient struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentClientMockRecorder
}

// MockDeploymentClientMockRecorder is the mock recorder for MockDeploymentClient.
type MockDeploymentClientMockRecorder struct {
	mock *MockDeploymentClient
}

// NewMockDeploymentClient creates a new mock instance.
func NewMockDeploymentClient(ctrl *gomock.Controller) *MockDeploymentClient {
	mock := &MockDeploymentClient{ctrl: ctrl}
	mock.recorder = &MockDeploymentClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentClient) EXPECT() *MockDeploymentClientMockRecorder {
	return m.recorder
}

// GetDeployments mocks base method.
func (m *MockDeploymentClient) GetDeployments(namespace, labelSelector string) (*v1.DeploymentList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployments", namespace, labelSelector)
	ret0, _ := ret[0].(*v1.DeploymentList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployments indicates an expected call of GetDeployments.
func (mr *MockDeploymentClientMockRecorder) GetDeployments(namespace, labelSelector interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployments", reflect.TypeOf((*MockDeploymentClient)(nil).GetDeployments), namespace, labelSelector)
}

// MockServerVersionClient is a mock of ServerVersionClient interface.
type MockServerVersionClient struct {
	ctrl     *gomock.Controller
	recorder *MockServerVersionClientMockRecorder
}

// MockServerVersionClientMockRecorder is the mock recorder for MockServerVersionClient.
type MockServerVersionClientMockRecorder struct {
	mock *MockServerVersionClient
}

// NewMockServerVersionClient creates a new mock instance.
func NewMockServerVersionClient(ctrl *gomock.Controller) *MockServerVersionClient {
	mock := &MockServerVersionClient{ctrl: ctrl}
	mock.recorder = &MockServerVersionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerVersionClient) EXPECT() *MockServerVersionClientMockRecorder {
	return m.recorder
}

// GetServerVersion mocks base method.
func (m *MockServerVersionClient) GetServerVersion() (*server.ServerVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerVersion")
	ret0, _ := ret[0].(*server.ServerVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerVersion indicates an expected call of GetServerVersion.
func (mr *MockServerVersionClientMockRecorder) GetServerVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerVersion", reflect.TypeOf((*MockServerVersionClient)(nil).GetServerVersion))
}
