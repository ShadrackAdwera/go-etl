// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ShadrackAdwera/go-etl/db/sqlc (interfaces: TxStore)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockTxStore is a mock of TxStore interface.
type MockTxStore struct {
	ctrl     *gomock.Controller
	recorder *MockTxStoreMockRecorder
}

// MockTxStoreMockRecorder is the mock recorder for MockTxStore.
type MockTxStoreMockRecorder struct {
	mock *MockTxStore
}

// NewMockTxStore creates a new mock instance.
func NewMockTxStore(ctrl *gomock.Controller) *MockTxStore {
	mock := &MockTxStore{ctrl: ctrl}
	mock.recorder = &MockTxStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTxStore) EXPECT() *MockTxStoreMockRecorder {
	return m.recorder
}

// CreateFile mocks base method.
func (m *MockTxStore) CreateFile(arg0 context.Context, arg1 db.CreateFileParams) (db.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFile", arg0, arg1)
	ret0, _ := ret[0].(db.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFile indicates an expected call of CreateFile.
func (mr *MockTxStoreMockRecorder) CreateFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFile", reflect.TypeOf((*MockTxStore)(nil).CreateFile), arg0, arg1)
}

// CreateMatchData mocks base method.
func (m *MockTxStore) CreateMatchData(arg0 context.Context, arg1 db.CreateMatchDataParams) (db.MatchDatum, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMatchData", arg0, arg1)
	ret0, _ := ret[0].(db.MatchDatum)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMatchData indicates an expected call of CreateMatchData.
func (mr *MockTxStoreMockRecorder) CreateMatchData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMatchData", reflect.TypeOf((*MockTxStore)(nil).CreateMatchData), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockTxStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockTxStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockTxStore)(nil).CreateSession), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockTxStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockTxStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockTxStore)(nil).CreateUser), arg0, arg1)
}

// DeleteMatchData mocks base method.
func (m *MockTxStore) DeleteMatchData(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMatchData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMatchData indicates an expected call of DeleteMatchData.
func (mr *MockTxStoreMockRecorder) DeleteMatchData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMatchData", reflect.TypeOf((*MockTxStore)(nil).DeleteMatchData), arg0, arg1)
}

// FindUserByEmail mocks base method.
func (m *MockTxStore) FindUserByEmail(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockTxStoreMockRecorder) FindUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockTxStore)(nil).FindUserByEmail), arg0, arg1)
}

// GetFiles mocks base method.
func (m *MockTxStore) GetFiles(arg0 context.Context, arg1 int64) ([]db.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFiles", arg0, arg1)
	ret0, _ := ret[0].([]db.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFiles indicates an expected call of GetFiles.
func (mr *MockTxStoreMockRecorder) GetFiles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFiles", reflect.TypeOf((*MockTxStore)(nil).GetFiles), arg0, arg1)
}

// GetMatchData mocks base method.
func (m *MockTxStore) GetMatchData(arg0 context.Context, arg1 int64) ([]db.MatchDatum, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchData", arg0, arg1)
	ret0, _ := ret[0].([]db.MatchDatum)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatchData indicates an expected call of GetMatchData.
func (mr *MockTxStoreMockRecorder) GetMatchData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchData", reflect.TypeOf((*MockTxStore)(nil).GetMatchData), arg0, arg1)
}

// GetMatchDataById mocks base method.
func (m *MockTxStore) GetMatchDataById(arg0 context.Context, arg1 int64) (db.MatchDatum, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchDataById", arg0, arg1)
	ret0, _ := ret[0].(db.MatchDatum)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatchDataById indicates an expected call of GetMatchDataById.
func (mr *MockTxStoreMockRecorder) GetMatchDataById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchDataById", reflect.TypeOf((*MockTxStore)(nil).GetMatchDataById), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockTxStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockTxStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockTxStore)(nil).GetSession), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockTxStore) GetUser(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockTxStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockTxStore)(nil).GetUser), arg0, arg1)
}

// UpdateMatchData mocks base method.
func (m *MockTxStore) UpdateMatchData(arg0 context.Context, arg1 db.UpdateMatchDataParams) (db.MatchDatum, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMatchData", arg0, arg1)
	ret0, _ := ret[0].(db.MatchDatum)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMatchData indicates an expected call of UpdateMatchData.
func (mr *MockTxStoreMockRecorder) UpdateMatchData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMatchData", reflect.TypeOf((*MockTxStore)(nil).UpdateMatchData), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockTxStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockTxStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockTxStore)(nil).UpdateUser), arg0, arg1)
}
