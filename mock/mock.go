// Automatically generated by MockGen. DO NOT EDIT!
// Source: interface.go

package mock_gocbadaptor

import (
	gocb "github.com/couchbase/gocb"
	conf "github.com/deadcheat/gocbadaptor/conf"
	gomock "github.com/golang/mock/gomock"
)

// Mock of CouchAdapter interface
type MockCouchAdapter struct {
	ctrl     *gomock.Controller
	recorder *_MockCouchAdapterRecorder
}

// Recorder for MockCouchAdapter (not exported)
type _MockCouchAdapterRecorder struct {
	mock *MockCouchAdapter
}

func NewMockCouchAdapter(ctrl *gomock.Controller) *MockCouchAdapter {
	mock := &MockCouchAdapter{ctrl: ctrl}
	mock.recorder = &_MockCouchAdapterRecorder{mock}
	return mock
}

func (_m *MockCouchAdapter) EXPECT() *_MockCouchAdapterRecorder {
	return _m.recorder
}

func (_m *MockCouchAdapter) Open(couchenv *conf.Env) *gocb.Bucket {
	ret := _m.ctrl.Call(_m, "Open", couchenv)
	ret0, _ := ret[0].(*gocb.Bucket)
	return ret0
}

func (_mr *_MockCouchAdapterRecorder) Open(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Open", arg0)
}

func (_m *MockCouchAdapter) Get(b *gocb.Bucket, key string) (gocb.Cas, []byte, bool) {
	ret := _m.ctrl.Call(_m, "Get", b, key)
	ret0, _ := ret[0].(gocb.Cas)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

func (_mr *_MockCouchAdapterRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1)
}

func (_m *MockCouchAdapter) Insert(b *gocb.Bucket, key string, data []byte, expiry uint32) (gocb.Cas, bool) {
	ret := _m.ctrl.Call(_m, "Insert", b, key, data, expiry)
	ret0, _ := ret[0].(gocb.Cas)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockCouchAdapterRecorder) Insert(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Insert", arg0, arg1, arg2, arg3)
}

func (_m *MockCouchAdapter) Upsert(b *gocb.Bucket, key string, data []byte, expiry uint32) (gocb.Cas, bool) {
	ret := _m.ctrl.Call(_m, "Upsert", b, key, data, expiry)
	ret0, _ := ret[0].(gocb.Cas)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockCouchAdapterRecorder) Upsert(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Upsert", arg0, arg1, arg2, arg3)
}
