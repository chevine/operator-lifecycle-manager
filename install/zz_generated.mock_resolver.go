// Automatically generated by MockGen. DO NOT EDIT!
// Source: install/resolver.go

package install

import (
	v1alpha1 "github.com/coreos-inc/alm/apis/clusterserviceversion/v1alpha1"
	client "github.com/coreos-inc/operator-client/pkg/client"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Mock of Strategy interface
type MockStrategy struct {
	ctrl     *gomock.Controller
	recorder *_MockStrategyRecorder
}

// Recorder for MockStrategy (not exported)
type _MockStrategyRecorder struct {
	mock *MockStrategy
}

func NewMockStrategy(ctrl *gomock.Controller) *MockStrategy {
	mock := &MockStrategy{ctrl: ctrl}
	mock.recorder = &_MockStrategyRecorder{mock}
	return mock
}

func (_m *MockStrategy) EXPECT() *_MockStrategyRecorder {
	return _m.recorder
}

func (_m *MockStrategy) GetStrategyName() string {
	ret := _m.ctrl.Call(_m, "GetStrategyName")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockStrategyRecorder) GetStrategyName() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetStrategyName")
}

// Mock of StrategyInstaller interface
type MockStrategyInstaller struct {
	ctrl     *gomock.Controller
	recorder *_MockStrategyInstallerRecorder
}

// Recorder for MockStrategyInstaller (not exported)
type _MockStrategyInstallerRecorder struct {
	mock *MockStrategyInstaller
}

func NewMockStrategyInstaller(ctrl *gomock.Controller) *MockStrategyInstaller {
	mock := &MockStrategyInstaller{ctrl: ctrl}
	mock.recorder = &_MockStrategyInstallerRecorder{mock}
	return mock
}

func (_m *MockStrategyInstaller) EXPECT() *_MockStrategyInstallerRecorder {
	return _m.recorder
}

func (_m *MockStrategyInstaller) Install(strategy Strategy) error {
	ret := _m.ctrl.Call(_m, "Install", strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStrategyInstallerRecorder) Install(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Install", arg0)
}

func (_m *MockStrategyInstaller) CheckInstalled(strategy Strategy) (bool, error) {
	ret := _m.ctrl.Call(_m, "CheckInstalled", strategy)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStrategyInstallerRecorder) CheckInstalled(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CheckInstalled", arg0)
}

// Mock of StrategyResolverInterface interface
type MockStrategyResolverInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockStrategyResolverInterfaceRecorder
}

// Recorder for MockStrategyResolverInterface (not exported)
type _MockStrategyResolverInterfaceRecorder struct {
	mock *MockStrategyResolverInterface
}

func NewMockStrategyResolverInterface(ctrl *gomock.Controller) *MockStrategyResolverInterface {
	mock := &MockStrategyResolverInterface{ctrl: ctrl}
	mock.recorder = &_MockStrategyResolverInterfaceRecorder{mock}
	return mock
}

func (_m *MockStrategyResolverInterface) EXPECT() *_MockStrategyResolverInterfaceRecorder {
	return _m.recorder
}

func (_m *MockStrategyResolverInterface) UnmarshalStrategy(s v1alpha1.NamedInstallStrategy) (Strategy, error) {
	ret := _m.ctrl.Call(_m, "UnmarshalStrategy", s)
	ret0, _ := ret[0].(Strategy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStrategyResolverInterfaceRecorder) UnmarshalStrategy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnmarshalStrategy", arg0)
}

func (_m *MockStrategyResolverInterface) InstallerForStrategy(strategyName string, opClient client.Interface, ownerMeta v1.ObjectMeta) StrategyInstaller {
	ret := _m.ctrl.Call(_m, "InstallerForStrategy", strategyName, opClient, ownerMeta)
	ret0, _ := ret[0].(StrategyInstaller)
	return ret0
}

func (_mr *_MockStrategyResolverInterfaceRecorder) InstallerForStrategy(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InstallerForStrategy", arg0, arg1, arg2)
}