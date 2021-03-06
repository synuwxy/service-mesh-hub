// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_selection is a generated GoMock package.
package mock_selection

import (
	reflect "reflect"

	appmesh "github.com/aws/aws-sdk-go/service/appmesh"
	eks "github.com/aws/aws-sdk-go/service/eks"
	gomock "github.com/golang/mock/gomock"
	types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	selection "github.com/solo-io/service-mesh-hub/pkg/aws/selection"
)

// MockAwsSelector is a mock of AwsSelector interface.
type MockAwsSelector struct {
	ctrl     *gomock.Controller
	recorder *MockAwsSelectorMockRecorder
}

// MockAwsSelectorMockRecorder is the mock recorder for MockAwsSelector.
type MockAwsSelectorMockRecorder struct {
	mock *MockAwsSelector
}

// NewMockAwsSelector creates a new mock instance.
func NewMockAwsSelector(ctrl *gomock.Controller) *MockAwsSelector {
	mock := &MockAwsSelector{ctrl: ctrl}
	mock.recorder = &MockAwsSelectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAwsSelector) EXPECT() *MockAwsSelectorMockRecorder {
	return m.recorder
}

// ResourceSelectorsByRegion mocks base method.
func (m *MockAwsSelector) ResourceSelectorsByRegion(resourceSelectors []*types.SettingsSpec_AwsAccount_ResourceSelector) (selection.AwsSelectorsByRegion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceSelectorsByRegion", resourceSelectors)
	ret0, _ := ret[0].(selection.AwsSelectorsByRegion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResourceSelectorsByRegion indicates an expected call of ResourceSelectorsByRegion.
func (mr *MockAwsSelectorMockRecorder) ResourceSelectorsByRegion(resourceSelectors interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceSelectorsByRegion", reflect.TypeOf((*MockAwsSelector)(nil).ResourceSelectorsByRegion), resourceSelectors)
}

// AwsSelectorsForAllRegions mocks base method.
func (m *MockAwsSelector) AwsSelectorsForAllRegions() selection.AwsSelectorsByRegion {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AwsSelectorsForAllRegions")
	ret0, _ := ret[0].(selection.AwsSelectorsByRegion)
	return ret0
}

// AwsSelectorsForAllRegions indicates an expected call of AwsSelectorsForAllRegions.
func (mr *MockAwsSelectorMockRecorder) AwsSelectorsForAllRegions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AwsSelectorsForAllRegions", reflect.TypeOf((*MockAwsSelector)(nil).AwsSelectorsForAllRegions))
}

// IsDiscoverAll mocks base method.
func (m *MockAwsSelector) IsDiscoverAll(discoverySettings *types.SettingsSpec_AwsAccount_DiscoverySelector) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDiscoverAll", discoverySettings)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDiscoverAll indicates an expected call of IsDiscoverAll.
func (mr *MockAwsSelectorMockRecorder) IsDiscoverAll(discoverySettings interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDiscoverAll", reflect.TypeOf((*MockAwsSelector)(nil).IsDiscoverAll), discoverySettings)
}

// AppMeshMatchedBySelectors mocks base method.
func (m *MockAwsSelector) AppMeshMatchedBySelectors(appmeshRef *appmesh.MeshRef, appmeshTags []*appmesh.TagRef, selectors []*types.SettingsSpec_AwsAccount_ResourceSelector) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppMeshMatchedBySelectors", appmeshRef, appmeshTags, selectors)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppMeshMatchedBySelectors indicates an expected call of AppMeshMatchedBySelectors.
func (mr *MockAwsSelectorMockRecorder) AppMeshMatchedBySelectors(appmeshRef, appmeshTags, selectors interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppMeshMatchedBySelectors", reflect.TypeOf((*MockAwsSelector)(nil).AppMeshMatchedBySelectors), appmeshRef, appmeshTags, selectors)
}

// EKSMatchedBySelectors mocks base method.
func (m *MockAwsSelector) EKSMatchedBySelectors(eksCluster *eks.Cluster, selectors []*types.SettingsSpec_AwsAccount_ResourceSelector) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EKSMatchedBySelectors", eksCluster, selectors)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EKSMatchedBySelectors indicates an expected call of EKSMatchedBySelectors.
func (mr *MockAwsSelectorMockRecorder) EKSMatchedBySelectors(eksCluster, selectors interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EKSMatchedBySelectors", reflect.TypeOf((*MockAwsSelector)(nil).EKSMatchedBySelectors), eksCluster, selectors)
}
