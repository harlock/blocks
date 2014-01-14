// Automatically generated by MockGen. DO NOT EDIT!
// Source: routable.go

package mocks

import (
	. "github.com/harlock/blocks"
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of Pather interface
type MockPather struct {
	ctrl     *gomock.Controller
	recorder *_MockPatherRecorder
}

// Recorder for MockPather (not exported)
type _MockPatherRecorder struct {
	mock *MockPather
}

func NewMockPather(ctrl *gomock.Controller) *MockPather {
	mock := &MockPather{ctrl: ctrl}
	mock.recorder = &_MockPatherRecorder{mock}
	return mock
}

func (_m *MockPather) EXPECT() *_MockPatherRecorder {
	return _m.recorder
}

func (_m *MockPather) Path() string {
	ret := _m.ctrl.Call(_m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockPatherRecorder) Path() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Path")
}

func (_m *MockPather) Method() string {
	ret := _m.ctrl.Call(_m, "Method")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockPatherRecorder) Method() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Method")
}

// Mock of Routable interface
type MockRoutable struct {
	ctrl     *gomock.Controller
	recorder *_MockRoutableRecorder
}

// Recorder for MockRoutable (not exported)
type _MockRoutableRecorder struct {
	mock *MockRoutable
}

func NewMockRoutable(ctrl *gomock.Controller) *MockRoutable {
	mock := &MockRoutable{ctrl: ctrl}
	mock.recorder = &_MockRoutableRecorder{mock}
	return mock
}

func (_m *MockRoutable) EXPECT() *_MockRoutableRecorder {
	return _m.recorder
}

func (_m *MockRoutable) Path() string {
	ret := _m.ctrl.Call(_m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockRoutableRecorder) Path() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Path")
}

func (_m *MockRoutable) Find(_param0 Pather) (*Route, bool) {
	ret := _m.ctrl.Call(_m, "Find", _param0)
	ret0, _ := ret[0].(*Route)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockRoutableRecorder) Find(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Find", arg0)
}

func (_m *MockRoutable) Match(_param0 Pather) bool {
	ret := _m.ctrl.Call(_m, "Match", _param0)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockRoutableRecorder) Match(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Match", arg0)
}

func (_m *MockRoutable) Inspect() {
	_m.ctrl.Call(_m, "Inspect")
}

func (_mr *_MockRoutableRecorder) Inspect() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Inspect")
}

func (_m *MockRoutable) Get(_param0 string, _param1 interface{}, _param2 string) *Route {
	ret := _m.ctrl.Call(_m, "Get", _param0, _param1, _param2)
	ret0, _ := ret[0].(*Route)
	return ret0
}

func (_mr *_MockRoutableRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1, arg2)
}

func (_m *MockRoutable) Resources(_param0 interface{}) Routable {
	ret := _m.ctrl.Call(_m, "Resources", _param0)
	ret0, _ := ret[0].(Routable)
	return ret0
}

func (_mr *_MockRoutableRecorder) Resources(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Resources", arg0)
}

func (_m *MockRoutable) Namespace(_param0 string) Routable {
	ret := _m.ctrl.Call(_m, "Namespace", _param0)
	ret0, _ := ret[0].(Routable)
	return ret0
}

func (_mr *_MockRoutableRecorder) Namespace(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Namespace", arg0)
}
