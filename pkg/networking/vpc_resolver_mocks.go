// Code generated by MockGen. DO NOT EDIT.
// Source: sigs.k8s.io/aws-load-balancer-controller/pkg/networking (interfaces: VPCResolver)

// Package networking is a generated GoMock package.
package networking

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockVPCResolver is a mock of VPCResolver interface.
type MockVPCResolver struct {
	ctrl     *gomock.Controller
	recorder *MockVPCResolverMockRecorder
}

// MockVPCResolverMockRecorder is the mock recorder for MockVPCResolver.
type MockVPCResolverMockRecorder struct {
	mock *MockVPCResolver
}

// NewMockVPCResolver creates a new mock instance.
func NewMockVPCResolver(ctrl *gomock.Controller) *MockVPCResolver {
	mock := &MockVPCResolver{ctrl: ctrl}
	mock.recorder = &MockVPCResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVPCResolver) EXPECT() *MockVPCResolverMockRecorder {
	return m.recorder
}

// ResolveCIDRs mocks base method.
func (m *MockVPCResolver) ResolveCIDRs(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveCIDRs", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveCIDRs indicates an expected call of ResolveCIDRs.
func (mr *MockVPCResolverMockRecorder) ResolveCIDRs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveCIDRs", reflect.TypeOf((*MockVPCResolver)(nil).ResolveCIDRs), arg0)
}

// ResolveIPv6CIDRs mocks base method.
func (m *MockVPCResolver) ResolveIPv6CIDRs(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveIPv6CIDRs", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveIPv6CIDRs indicates an expected call of ResolveIPv6CIDRs.
func (mr *MockVPCResolverMockRecorder) ResolveIPv6CIDRs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveIPv6CIDRs", reflect.TypeOf((*MockVPCResolver)(nil).ResolveIPv6CIDRs), arg0)
}
