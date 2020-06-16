// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/openshift-azure/pkg/util/azureclient/graphrbac (interfaces: ApplicationsClient,GroupsClient,ServicePrincipalsClient)

// Package mock_graphrbac is a generated GoMock package.
package mock_graphrbac

import (
	context "context"
	reflect "reflect"

	graphrbac "github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
)

// MockApplicationsClient is a mock of ApplicationsClient interface.
type MockApplicationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationsClientMockRecorder
}

// MockApplicationsClientMockRecorder is the mock recorder for MockApplicationsClient.
type MockApplicationsClientMockRecorder struct {
	mock *MockApplicationsClient
}

// NewMockApplicationsClient creates a new mock instance.
func NewMockApplicationsClient(ctrl *gomock.Controller) *MockApplicationsClient {
	mock := &MockApplicationsClient{ctrl: ctrl}
	mock.recorder = &MockApplicationsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationsClient) EXPECT() *MockApplicationsClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockApplicationsClient) Create(arg0 context.Context, arg1 graphrbac.ApplicationCreateParameters) (graphrbac.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(graphrbac.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockApplicationsClientMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockApplicationsClient)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockApplicationsClient) Delete(arg0 context.Context, arg1 string) (autorest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(autorest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockApplicationsClientMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockApplicationsClient)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockApplicationsClient) Get(arg0 context.Context, arg1 string) (graphrbac.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(graphrbac.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockApplicationsClientMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockApplicationsClient)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockApplicationsClient) List(arg0 context.Context, arg1 string) (graphrbac.ApplicationListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(graphrbac.ApplicationListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockApplicationsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockApplicationsClient)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockApplicationsClient) Patch(arg0 context.Context, arg1 string, arg2 graphrbac.ApplicationUpdateParameters) (autorest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", arg0, arg1, arg2)
	ret0, _ := ret[0].(autorest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockApplicationsClientMockRecorder) Patch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockApplicationsClient)(nil).Patch), arg0, arg1, arg2)
}

// MockGroupsClient is a mock of GroupsClient interface.
type MockGroupsClient struct {
	ctrl     *gomock.Controller
	recorder *MockGroupsClientMockRecorder
}

// MockGroupsClientMockRecorder is the mock recorder for MockGroupsClient.
type MockGroupsClientMockRecorder struct {
	mock *MockGroupsClient
}

// NewMockGroupsClient creates a new mock instance.
func NewMockGroupsClient(ctrl *gomock.Controller) *MockGroupsClient {
	mock := &MockGroupsClient{ctrl: ctrl}
	mock.recorder = &MockGroupsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupsClient) EXPECT() *MockGroupsClientMockRecorder {
	return m.recorder
}

// GetGroupMembers mocks base method.
func (m *MockGroupsClient) GetGroupMembers(arg0 context.Context, arg1 string) ([]graphrbac.BasicDirectoryObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupMembers", arg0, arg1)
	ret0, _ := ret[0].([]graphrbac.BasicDirectoryObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupMembers indicates an expected call of GetGroupMembers.
func (mr *MockGroupsClientMockRecorder) GetGroupMembers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMembers", reflect.TypeOf((*MockGroupsClient)(nil).GetGroupMembers), arg0, arg1)
}

// MockServicePrincipalsClient is a mock of ServicePrincipalsClient interface.
type MockServicePrincipalsClient struct {
	ctrl     *gomock.Controller
	recorder *MockServicePrincipalsClientMockRecorder
}

// MockServicePrincipalsClientMockRecorder is the mock recorder for MockServicePrincipalsClient.
type MockServicePrincipalsClientMockRecorder struct {
	mock *MockServicePrincipalsClient
}

// NewMockServicePrincipalsClient creates a new mock instance.
func NewMockServicePrincipalsClient(ctrl *gomock.Controller) *MockServicePrincipalsClient {
	mock := &MockServicePrincipalsClient{ctrl: ctrl}
	mock.recorder = &MockServicePrincipalsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicePrincipalsClient) EXPECT() *MockServicePrincipalsClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockServicePrincipalsClient) Create(arg0 context.Context, arg1 graphrbac.ServicePrincipalCreateParameters) (graphrbac.ServicePrincipal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(graphrbac.ServicePrincipal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServicePrincipalsClientMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServicePrincipalsClient)(nil).Create), arg0, arg1)
}

// List mocks base method.
func (m *MockServicePrincipalsClient) List(arg0 context.Context, arg1 string) (graphrbac.ServicePrincipalListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(graphrbac.ServicePrincipalListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockServicePrincipalsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServicePrincipalsClient)(nil).List), arg0, arg1)
}
