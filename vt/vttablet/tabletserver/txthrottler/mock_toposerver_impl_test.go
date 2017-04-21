// Automatically generated by MockGen. DO NOT EDIT!
// Source: gopkg.in/sqle/vitess-go.v2/vt/topo (interfaces: Impl)

package txthrottler

import (
	gomock "github.com/golang/mock/gomock"
	topodata "gopkg.in/sqle/vitess-go.v2/vt/proto/topodata"
	vschema "gopkg.in/sqle/vitess-go.v2/vt/proto/vschema"
	topo "gopkg.in/sqle/vitess-go.v2/vt/topo"
	context "golang.org/x/net/context"
)

// Mock of Impl interface
type MockImpl struct {
	ctrl     *gomock.Controller
	recorder *_MockImplRecorder
}

// Recorder for MockImpl (not exported)
type _MockImplRecorder struct {
	mock *MockImpl
}

func NewMockImpl(ctrl *gomock.Controller) *MockImpl {
	mock := &MockImpl{ctrl: ctrl}
	mock.recorder = &_MockImplRecorder{mock}
	return mock
}

func (_m *MockImpl) EXPECT() *_MockImplRecorder {
	return _m.recorder
}

func (_m *MockImpl) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockImplRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockImpl) Create(_param0 context.Context, _param1 string, _param2 string, _param3 []byte) (topo.Version, error) {
	ret := _m.ctrl.Call(_m, "Create", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(topo.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) Create(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0, arg1, arg2, arg3)
}

func (_m *MockImpl) CreateKeyspace(_param0 context.Context, _param1 string, _param2 *topodata.Keyspace) error {
	ret := _m.ctrl.Call(_m, "CreateKeyspace", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplRecorder) CreateKeyspace(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateKeyspace", arg0, arg1, arg2)
}

func (_m *MockImpl) CreateShard(_param0 context.Context, _param1 string, _param2 string, _param3 *topodata.Shard) error {
	ret := _m.ctrl.Call(_m, "CreateShard", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplRecorder) CreateShard(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateShard", arg0, arg1, arg2, arg3)
}

func (_m *MockImpl) CreateTablet(_param0 context.Context, _param1 *topodata.Tablet) error {
	ret := _m.ctrl.Call(_m, "CreateTablet", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplRecorder) CreateTablet(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateTablet", arg0, arg1)
}

func (_m *MockImpl) Delete(_param0 context.Context, _param1 string, _param2 string, _param3 topo.Version) error {
	ret := _m.ctrl.Call(_m, "Delete", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0, arg1, arg2, arg3)
}

func (_m *MockImpl) DeleteKeyspace(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteKeyspace", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplRecorder) DeleteKeyspace(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteKeyspace", arg0, arg1)
}

func (_m *MockImpl) DeleteKeyspaceReplication(_param0 context.Context, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "DeleteKeyspaceReplication", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplRecorder) DeleteKeyspaceReplication(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteKeyspaceReplication", arg0, arg1, arg2)
}

func (_m *MockImpl) DeleteShard(_param0 context.Context, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "DeleteShard", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplRecorder) DeleteShard(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteShard", arg0, arg1, arg2)
}

func (_m *MockImpl) DeleteShardReplication(_param0 context.Context, _param1 string, _param2 string, _param3 string) error {
	ret := _m.ctrl.Call(_m, "DeleteShardReplication", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplRecorder) DeleteShardReplication(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteShardReplication", arg0, arg1, arg2, arg3)
}

func (_m *MockImpl) DeleteSrvKeyspace(_param0 context.Context, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "DeleteSrvKeyspace", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplRecorder) DeleteSrvKeyspace(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteSrvKeyspace", arg0, arg1, arg2)
}

func (_m *MockImpl) DeleteTablet(_param0 context.Context, _param1 *topodata.TabletAlias) error {
	ret := _m.ctrl.Call(_m, "DeleteTablet", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplRecorder) DeleteTablet(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteTablet", arg0, arg1)
}

func (_m *MockImpl) Get(_param0 context.Context, _param1 string, _param2 string) ([]byte, topo.Version, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0, _param1, _param2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(topo.Version)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockImplRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1, arg2)
}

func (_m *MockImpl) GetKeyspace(_param0 context.Context, _param1 string) (*topodata.Keyspace, int64, error) {
	ret := _m.ctrl.Call(_m, "GetKeyspace", _param0, _param1)
	ret0, _ := ret[0].(*topodata.Keyspace)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockImplRecorder) GetKeyspace(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetKeyspace", arg0, arg1)
}

func (_m *MockImpl) GetKeyspaces(_param0 context.Context) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetKeyspaces", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) GetKeyspaces(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetKeyspaces", arg0)
}

func (_m *MockImpl) GetKnownCells(_param0 context.Context) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetKnownCells", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) GetKnownCells(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetKnownCells", arg0)
}

func (_m *MockImpl) GetShard(_param0 context.Context, _param1 string, _param2 string) (*topodata.Shard, int64, error) {
	ret := _m.ctrl.Call(_m, "GetShard", _param0, _param1, _param2)
	ret0, _ := ret[0].(*topodata.Shard)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockImplRecorder) GetShard(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetShard", arg0, arg1, arg2)
}

func (_m *MockImpl) GetShardNames(_param0 context.Context, _param1 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetShardNames", _param0, _param1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) GetShardNames(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetShardNames", arg0, arg1)
}

func (_m *MockImpl) GetShardReplication(_param0 context.Context, _param1 string, _param2 string, _param3 string) (*topo.ShardReplicationInfo, error) {
	ret := _m.ctrl.Call(_m, "GetShardReplication", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(*topo.ShardReplicationInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) GetShardReplication(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetShardReplication", arg0, arg1, arg2, arg3)
}

func (_m *MockImpl) GetSrvKeyspace(_param0 context.Context, _param1 string, _param2 string) (*topodata.SrvKeyspace, error) {
	ret := _m.ctrl.Call(_m, "GetSrvKeyspace", _param0, _param1, _param2)
	ret0, _ := ret[0].(*topodata.SrvKeyspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) GetSrvKeyspace(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSrvKeyspace", arg0, arg1, arg2)
}

func (_m *MockImpl) GetSrvKeyspaceNames(_param0 context.Context, _param1 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetSrvKeyspaceNames", _param0, _param1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) GetSrvKeyspaceNames(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSrvKeyspaceNames", arg0, arg1)
}

func (_m *MockImpl) GetSrvVSchema(_param0 context.Context, _param1 string) (*vschema.SrvVSchema, error) {
	ret := _m.ctrl.Call(_m, "GetSrvVSchema", _param0, _param1)
	ret0, _ := ret[0].(*vschema.SrvVSchema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) GetSrvVSchema(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSrvVSchema", arg0, arg1)
}

func (_m *MockImpl) GetTablet(_param0 context.Context, _param1 *topodata.TabletAlias) (*topodata.Tablet, int64, error) {
	ret := _m.ctrl.Call(_m, "GetTablet", _param0, _param1)
	ret0, _ := ret[0].(*topodata.Tablet)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockImplRecorder) GetTablet(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTablet", arg0, arg1)
}

func (_m *MockImpl) GetTabletsByCell(_param0 context.Context, _param1 string) ([]*topodata.TabletAlias, error) {
	ret := _m.ctrl.Call(_m, "GetTabletsByCell", _param0, _param1)
	ret0, _ := ret[0].([]*topodata.TabletAlias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) GetTabletsByCell(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTabletsByCell", arg0, arg1)
}

func (_m *MockImpl) GetVSchema(_param0 context.Context, _param1 string) (*vschema.Keyspace, error) {
	ret := _m.ctrl.Call(_m, "GetVSchema", _param0, _param1)
	ret0, _ := ret[0].(*vschema.Keyspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) GetVSchema(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetVSchema", arg0, arg1)
}

func (_m *MockImpl) ListDir(_param0 context.Context, _param1 string, _param2 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "ListDir", _param0, _param1, _param2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) ListDir(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListDir", arg0, arg1, arg2)
}

func (_m *MockImpl) LockKeyspaceForAction(_param0 context.Context, _param1 string, _param2 string) (string, error) {
	ret := _m.ctrl.Call(_m, "LockKeyspaceForAction", _param0, _param1, _param2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) LockKeyspaceForAction(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LockKeyspaceForAction", arg0, arg1, arg2)
}

func (_m *MockImpl) LockShardForAction(_param0 context.Context, _param1 string, _param2 string, _param3 string) (string, error) {
	ret := _m.ctrl.Call(_m, "LockShardForAction", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) LockShardForAction(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LockShardForAction", arg0, arg1, arg2, arg3)
}

func (_m *MockImpl) NewMasterParticipation(_param0 string, _param1 string) (topo.MasterParticipation, error) {
	ret := _m.ctrl.Call(_m, "NewMasterParticipation", _param0, _param1)
	ret0, _ := ret[0].(topo.MasterParticipation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) NewMasterParticipation(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewMasterParticipation", arg0, arg1)
}

func (_m *MockImpl) SaveVSchema(_param0 context.Context, _param1 string, _param2 *vschema.Keyspace) error {
	ret := _m.ctrl.Call(_m, "SaveVSchema", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplRecorder) SaveVSchema(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveVSchema", arg0, arg1, arg2)
}

func (_m *MockImpl) UnlockKeyspaceForAction(_param0 context.Context, _param1 string, _param2 string, _param3 string) error {
	ret := _m.ctrl.Call(_m, "UnlockKeyspaceForAction", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplRecorder) UnlockKeyspaceForAction(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnlockKeyspaceForAction", arg0, arg1, arg2, arg3)
}

func (_m *MockImpl) UnlockShardForAction(_param0 context.Context, _param1 string, _param2 string, _param3 string, _param4 string) error {
	ret := _m.ctrl.Call(_m, "UnlockShardForAction", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplRecorder) UnlockShardForAction(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnlockShardForAction", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockImpl) Update(_param0 context.Context, _param1 string, _param2 string, _param3 []byte, _param4 topo.Version) (topo.Version, error) {
	ret := _m.ctrl.Call(_m, "Update", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(topo.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) Update(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Update", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockImpl) UpdateKeyspace(_param0 context.Context, _param1 string, _param2 *topodata.Keyspace, _param3 int64) (int64, error) {
	ret := _m.ctrl.Call(_m, "UpdateKeyspace", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) UpdateKeyspace(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateKeyspace", arg0, arg1, arg2, arg3)
}

func (_m *MockImpl) UpdateShard(_param0 context.Context, _param1 string, _param2 string, _param3 *topodata.Shard, _param4 int64) (int64, error) {
	ret := _m.ctrl.Call(_m, "UpdateShard", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) UpdateShard(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateShard", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockImpl) UpdateShardReplicationFields(_param0 context.Context, _param1 string, _param2 string, _param3 string, _param4 func(*topodata.ShardReplication) error) error {
	ret := _m.ctrl.Call(_m, "UpdateShardReplicationFields", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplRecorder) UpdateShardReplicationFields(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateShardReplicationFields", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockImpl) UpdateSrvKeyspace(_param0 context.Context, _param1 string, _param2 string, _param3 *topodata.SrvKeyspace) error {
	ret := _m.ctrl.Call(_m, "UpdateSrvKeyspace", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplRecorder) UpdateSrvKeyspace(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateSrvKeyspace", arg0, arg1, arg2, arg3)
}

func (_m *MockImpl) UpdateSrvVSchema(_param0 context.Context, _param1 string, _param2 *vschema.SrvVSchema) error {
	ret := _m.ctrl.Call(_m, "UpdateSrvVSchema", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockImplRecorder) UpdateSrvVSchema(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateSrvVSchema", arg0, arg1, arg2)
}

func (_m *MockImpl) UpdateTablet(_param0 context.Context, _param1 *topodata.Tablet, _param2 int64) (int64, error) {
	ret := _m.ctrl.Call(_m, "UpdateTablet", _param0, _param1, _param2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockImplRecorder) UpdateTablet(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateTablet", arg0, arg1, arg2)
}

func (_m *MockImpl) Watch(_param0 context.Context, _param1 string, _param2 string) (*topo.WatchData, <-chan *topo.WatchData, topo.CancelFunc) {
	ret := _m.ctrl.Call(_m, "Watch", _param0, _param1, _param2)
	ret0, _ := ret[0].(*topo.WatchData)
	ret1, _ := ret[1].(<-chan *topo.WatchData)
	ret2, _ := ret[2].(topo.CancelFunc)
	return ret0, ret1, ret2
}

func (_mr *_MockImplRecorder) Watch(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Watch", arg0, arg1, arg2)
}
