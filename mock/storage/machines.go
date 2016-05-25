// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/docker/libmachete/storage (interfaces: Machines)

package storage

import (
	storage "github.com/docker/libmachete/storage"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Machines interface
type MockMachines struct {
	ctrl     *gomock.Controller
	recorder *_MockMachinesRecorder
}

// Recorder for MockMachines (not exported)
type _MockMachinesRecorder struct {
	mock *MockMachines
}

func NewMockMachines(ctrl *gomock.Controller) *MockMachines {
	mock := &MockMachines{ctrl: ctrl}
	mock.recorder = &_MockMachinesRecorder{mock}
	return mock
}

func (_m *MockMachines) EXPECT() *_MockMachinesRecorder {
	return _m.recorder
}

func (_m *MockMachines) Delete(_param0 storage.MachineID) error {
	ret := _m.ctrl.Call(_m, "Delete", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMachinesRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}

func (_m *MockMachines) GetDetails(_param0 storage.MachineID, _param1 interface{}) error {
	ret := _m.ctrl.Call(_m, "GetDetails", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMachinesRecorder) GetDetails(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDetails", arg0, arg1)
}

func (_m *MockMachines) GetRecord(_param0 storage.MachineID) (*storage.MachineRecord, error) {
	ret := _m.ctrl.Call(_m, "GetRecord", _param0)
	ret0, _ := ret[0].(*storage.MachineRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMachinesRecorder) GetRecord(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRecord", arg0)
}

func (_m *MockMachines) List() ([]storage.MachineID, error) {
	ret := _m.ctrl.Call(_m, "List")
	ret0, _ := ret[0].([]storage.MachineID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMachinesRecorder) List() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List")
}

func (_m *MockMachines) Save(_param0 storage.MachineRecord, _param1 interface{}) error {
	ret := _m.ctrl.Call(_m, "Save", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMachinesRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0, arg1)
}
