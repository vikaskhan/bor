// Code generated by MockGen. DO NOT EDIT.
// Source: consensus/bor/api/caller.go

// Package api is a generated GoMock package.
package api

import (
	context "context"
	reflect "reflect"

	common "github.com/ethereum/go-ethereum/common"
	hexutil "github.com/ethereum/go-ethereum/common/hexutil"
	state "github.com/ethereum/go-ethereum/core/state"
	ethapi "github.com/ethereum/go-ethereum/internal/ethapi"
	override "github.com/ethereum/go-ethereum/internal/ethapi/override"
	rpc "github.com/ethereum/go-ethereum/rpc"
	gomock "github.com/golang/mock/gomock"
)

// MockCaller is a mock of Caller interface.
type MockCaller struct {
	ctrl     *gomock.Controller
	recorder *MockCallerMockRecorder
}

// MockCallerMockRecorder is the mock recorder for MockCaller.
type MockCallerMockRecorder struct {
	mock *MockCaller
}

// NewMockCaller creates a new mock instance.
func NewMockCaller(ctrl *gomock.Controller) *MockCaller {
	mock := &MockCaller{ctrl: ctrl}
	mock.recorder = &MockCallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCaller) EXPECT() *MockCallerMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *MockCaller) Call(ctx context.Context, args ethapi.TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, overrides *override.StateOverride, blockOverrides *override.BlockOverrides) (hexutil.Bytes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", ctx, args, blockNrOrHash, overrides, blockOverrides)
	ret0, _ := ret[0].(hexutil.Bytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call.
func (mr *MockCallerMockRecorder) Call(ctx, args, blockNrOrHash, overrides, blockOverrides interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockCaller)(nil).Call), ctx, args, blockNrOrHash, overrides, blockOverrides)
}

// CallWithState mocks base method.
func (m *MockCaller) CallWithState(ctx context.Context, args ethapi.TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, state *state.StateDB, overrides *override.StateOverride, blockOverrides *override.BlockOverrides) (hexutil.Bytes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallWithState", ctx, args, blockNrOrHash, state, overrides, blockOverrides)
	ret0, _ := ret[0].(hexutil.Bytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallWithState indicates an expected call of CallWithState.
func (mr *MockCallerMockRecorder) CallWithState(ctx, args, blockNrOrHash, state, overrides, blockOverrides interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallWithState", reflect.TypeOf((*MockCaller)(nil).CallWithState), ctx, args, blockNrOrHash, state, overrides, blockOverrides)
}

// GetBalance mocks base method.
func (m *MockCaller) GetBalance(ctx context.Context, address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Big, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx, address, blockNrOrHash)
	ret0, _ := ret[0].(*hexutil.Big)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockCallerMockRecorder) GetBalance(ctx, address, blockNrOrHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockCaller)(nil).GetBalance), ctx, address, blockNrOrHash)
}
