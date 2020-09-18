// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/drain (interfaces: NodeDrainStrategyBuilder)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/openshift/managed-upgrade-operator/pkg/apis/upgrade/v1alpha1"
	drain "github.com/openshift/managed-upgrade-operator/pkg/drain"
	v1 "k8s.io/api/core/v1"
	reflect "reflect"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockNodeDrainStrategyBuilder is a mock of NodeDrainStrategyBuilder interface
type MockNodeDrainStrategyBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockNodeDrainStrategyBuilderMockRecorder
}

// MockNodeDrainStrategyBuilderMockRecorder is the mock recorder for MockNodeDrainStrategyBuilder
type MockNodeDrainStrategyBuilderMockRecorder struct {
	mock *MockNodeDrainStrategyBuilder
}

// NewMockNodeDrainStrategyBuilder creates a new mock instance
func NewMockNodeDrainStrategyBuilder(ctrl *gomock.Controller) *MockNodeDrainStrategyBuilder {
	mock := &MockNodeDrainStrategyBuilder{ctrl: ctrl}
	mock.recorder = &MockNodeDrainStrategyBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeDrainStrategyBuilder) EXPECT() *MockNodeDrainStrategyBuilderMockRecorder {
	return m.recorder
}

// NewNodeDrainStrategy mocks base method
func (m *MockNodeDrainStrategyBuilder) NewNodeDrainStrategy(arg0 client.Client, arg1 *v1alpha1.UpgradeConfig, arg2 *v1.Node, arg3 *drain.NodeDrain) (drain.NodeDrainStrategy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewNodeDrainStrategy", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(drain.NodeDrainStrategy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewNodeDrainStrategy indicates an expected call of NewNodeDrainStrategy
func (mr *MockNodeDrainStrategyBuilderMockRecorder) NewNodeDrainStrategy(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewNodeDrainStrategy", reflect.TypeOf((*MockNodeDrainStrategyBuilder)(nil).NewNodeDrainStrategy), arg0, arg1, arg2, arg3)
}
