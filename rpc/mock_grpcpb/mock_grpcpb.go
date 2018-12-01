// Code generated by MockGen. DO NOT EDIT.
// Source: grpc.pb.go

// Package mock_grpcpb is a generated GoMock package.
package mock_grpcpb

import (
	context "context"
	pb "github.com/coschain/contentos-go/rpc/pb"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockApiServiceClient is a mock of ApiServiceClient interface
type MockApiServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockApiServiceClientMockRecorder
}

// MockApiServiceClientMockRecorder is the mock recorder for MockApiServiceClient
type MockApiServiceClientMockRecorder struct {
	mock *MockApiServiceClient
}

// NewMockApiServiceClient creates a new mock instance
func NewMockApiServiceClient(ctrl *gomock.Controller) *MockApiServiceClient {
	mock := &MockApiServiceClient{ctrl: ctrl}
	mock.recorder = &MockApiServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiServiceClient) EXPECT() *MockApiServiceClientMockRecorder {
	return m.recorder
}

// GetAccountByName mocks base method
func (m *MockApiServiceClient) GetAccountByName(ctx context.Context, in *pb.GetAccountByNameRequest, opts ...grpc.CallOption) (*pb.AccountResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccountByName", varargs...)
	ret0, _ := ret[0].(*pb.AccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByName indicates an expected call of GetAccountByName
func (mr *MockApiServiceClientMockRecorder) GetAccountByName(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByName", reflect.TypeOf((*MockApiServiceClient)(nil).GetAccountByName), varargs...)
}

// GetAccountRewardByName mocks base method
func (m *MockApiServiceClient) GetAccountRewardByName(ctx context.Context, in *pb.GetAccountRewardByNameRequest, opts ...grpc.CallOption) (*pb.AccountRewardResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccountRewardByName", varargs...)
	ret0, _ := ret[0].(*pb.AccountRewardResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountRewardByName indicates an expected call of GetAccountRewardByName
func (mr *MockApiServiceClientMockRecorder) GetAccountRewardByName(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountRewardByName", reflect.TypeOf((*MockApiServiceClient)(nil).GetAccountRewardByName), varargs...)
}

// GetFollowerListByName mocks base method
func (m *MockApiServiceClient) GetFollowerListByName(ctx context.Context, in *pb.GetFollowerListByNameRequest, opts ...grpc.CallOption) (*pb.GetFollowerListByNameResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFollowerListByName", varargs...)
	ret0, _ := ret[0].(*pb.GetFollowerListByNameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowerListByName indicates an expected call of GetFollowerListByName
func (mr *MockApiServiceClientMockRecorder) GetFollowerListByName(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowerListByName", reflect.TypeOf((*MockApiServiceClient)(nil).GetFollowerListByName), varargs...)
}

// GetFollowingListByName mocks base method
func (m *MockApiServiceClient) GetFollowingListByName(ctx context.Context, in *pb.GetFollowingListByNameRequest, opts ...grpc.CallOption) (*pb.GetFollowingListByNameResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFollowingListByName", varargs...)
	ret0, _ := ret[0].(*pb.GetFollowingListByNameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowingListByName indicates an expected call of GetFollowingListByName
func (mr *MockApiServiceClientMockRecorder) GetFollowingListByName(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowingListByName", reflect.TypeOf((*MockApiServiceClient)(nil).GetFollowingListByName), varargs...)
}

// GetFollowCountByName mocks base method
func (m *MockApiServiceClient) GetFollowCountByName(ctx context.Context, in *pb.GetFollowCountByNameRequest, opts ...grpc.CallOption) (*pb.GetFollowCountByNameResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFollowCountByName", varargs...)
	ret0, _ := ret[0].(*pb.GetFollowCountByNameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowCountByName indicates an expected call of GetFollowCountByName
func (mr *MockApiServiceClientMockRecorder) GetFollowCountByName(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowCountByName", reflect.TypeOf((*MockApiServiceClient)(nil).GetFollowCountByName), varargs...)
}

// GetWitnessList mocks base method
func (m *MockApiServiceClient) GetWitnessList(ctx context.Context, in *pb.GetWitnessListRequest, opts ...grpc.CallOption) (*pb.GetWitnessListResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWitnessList", varargs...)
	ret0, _ := ret[0].(*pb.GetWitnessListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWitnessList indicates an expected call of GetWitnessList
func (mr *MockApiServiceClientMockRecorder) GetWitnessList(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWitnessList", reflect.TypeOf((*MockApiServiceClient)(nil).GetWitnessList), varargs...)
}

// GetPostListByCreated mocks base method
func (m *MockApiServiceClient) GetPostListByCreated(ctx context.Context, in *pb.GetPostListByCreatedRequest, opts ...grpc.CallOption) (*pb.GetPostListByCreatedResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPostListByCreated", varargs...)
	ret0, _ := ret[0].(*pb.GetPostListByCreatedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostListByCreated indicates an expected call of GetPostListByCreated
func (mr *MockApiServiceClientMockRecorder) GetPostListByCreated(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostListByCreated", reflect.TypeOf((*MockApiServiceClient)(nil).GetPostListByCreated), varargs...)
}

// GetReplyListByPostId mocks base method
func (m *MockApiServiceClient) GetReplyListByPostId(ctx context.Context, in *pb.GetReplyListByPostIdRequest, opts ...grpc.CallOption) (*pb.GetReplyListByPostIdResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReplyListByPostId", varargs...)
	ret0, _ := ret[0].(*pb.GetReplyListByPostIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReplyListByPostId indicates an expected call of GetReplyListByPostId
func (mr *MockApiServiceClientMockRecorder) GetReplyListByPostId(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplyListByPostId", reflect.TypeOf((*MockApiServiceClient)(nil).GetReplyListByPostId), varargs...)
}

// GetBlockTransactionsByNum mocks base method
func (m *MockApiServiceClient) GetBlockTransactionsByNum(ctx context.Context, in *pb.GetBlockTransactionsByNumRequest, opts ...grpc.CallOption) (*pb.GetBlockTransactionsByNumResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlockTransactionsByNum", varargs...)
	ret0, _ := ret[0].(*pb.GetBlockTransactionsByNumResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockTransactionsByNum indicates an expected call of GetBlockTransactionsByNum
func (mr *MockApiServiceClientMockRecorder) GetBlockTransactionsByNum(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockTransactionsByNum", reflect.TypeOf((*MockApiServiceClient)(nil).GetBlockTransactionsByNum), varargs...)
}

// GetTrxById mocks base method
func (m *MockApiServiceClient) GetTrxById(ctx context.Context, in *pb.GetTrxByIdRequest, opts ...grpc.CallOption) (*pb.GetTrxByIdResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTrxById", varargs...)
	ret0, _ := ret[0].(*pb.GetTrxByIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrxById indicates an expected call of GetTrxById
func (mr *MockApiServiceClientMockRecorder) GetTrxById(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrxById", reflect.TypeOf((*MockApiServiceClient)(nil).GetTrxById), varargs...)
}

// GetChainState mocks base method
func (m *MockApiServiceClient) GetChainState(ctx context.Context, in *pb.NonParamsRequest, opts ...grpc.CallOption) (*pb.GetChainStateResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChainState", varargs...)
	ret0, _ := ret[0].(*pb.GetChainStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainState indicates an expected call of GetChainState
func (mr *MockApiServiceClientMockRecorder) GetChainState(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainState", reflect.TypeOf((*MockApiServiceClient)(nil).GetChainState), varargs...)
}

// BroadcastTrx mocks base method
func (m *MockApiServiceClient) BroadcastTrx(ctx context.Context, in *pb.BroadcastTrxRequest, opts ...grpc.CallOption) (*pb.BroadcastTrxResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BroadcastTrx", varargs...)
	ret0, _ := ret[0].(*pb.BroadcastTrxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BroadcastTrx indicates an expected call of BroadcastTrx
func (mr *MockApiServiceClientMockRecorder) BroadcastTrx(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastTrx", reflect.TypeOf((*MockApiServiceClient)(nil).BroadcastTrx), varargs...)
}

// MockApiServiceServer is a mock of ApiServiceServer interface
type MockApiServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockApiServiceServerMockRecorder
}

// MockApiServiceServerMockRecorder is the mock recorder for MockApiServiceServer
type MockApiServiceServerMockRecorder struct {
	mock *MockApiServiceServer
}

// NewMockApiServiceServer creates a new mock instance
func NewMockApiServiceServer(ctrl *gomock.Controller) *MockApiServiceServer {
	mock := &MockApiServiceServer{ctrl: ctrl}
	mock.recorder = &MockApiServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiServiceServer) EXPECT() *MockApiServiceServerMockRecorder {
	return m.recorder
}

// GetAccountByName mocks base method
func (m *MockApiServiceServer) GetAccountByName(arg0 context.Context, arg1 *pb.GetAccountByNameRequest) (*pb.AccountResponse, error) {
	ret := m.ctrl.Call(m, "GetAccountByName", arg0, arg1)
	ret0, _ := ret[0].(*pb.AccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByName indicates an expected call of GetAccountByName
func (mr *MockApiServiceServerMockRecorder) GetAccountByName(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByName", reflect.TypeOf((*MockApiServiceServer)(nil).GetAccountByName), arg0, arg1)
}

// GetAccountRewardByName mocks base method
func (m *MockApiServiceServer) GetAccountRewardByName(arg0 context.Context, arg1 *pb.GetAccountRewardByNameRequest) (*pb.AccountRewardResponse, error) {
	ret := m.ctrl.Call(m, "GetAccountRewardByName", arg0, arg1)
	ret0, _ := ret[0].(*pb.AccountRewardResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountRewardByName indicates an expected call of GetAccountRewardByName
func (mr *MockApiServiceServerMockRecorder) GetAccountRewardByName(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountRewardByName", reflect.TypeOf((*MockApiServiceServer)(nil).GetAccountRewardByName), arg0, arg1)
}

// GetFollowerListByName mocks base method
func (m *MockApiServiceServer) GetFollowerListByName(arg0 context.Context, arg1 *pb.GetFollowerListByNameRequest) (*pb.GetFollowerListByNameResponse, error) {
	ret := m.ctrl.Call(m, "GetFollowerListByName", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetFollowerListByNameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowerListByName indicates an expected call of GetFollowerListByName
func (mr *MockApiServiceServerMockRecorder) GetFollowerListByName(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowerListByName", reflect.TypeOf((*MockApiServiceServer)(nil).GetFollowerListByName), arg0, arg1)
}

// GetFollowingListByName mocks base method
func (m *MockApiServiceServer) GetFollowingListByName(arg0 context.Context, arg1 *pb.GetFollowingListByNameRequest) (*pb.GetFollowingListByNameResponse, error) {
	ret := m.ctrl.Call(m, "GetFollowingListByName", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetFollowingListByNameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowingListByName indicates an expected call of GetFollowingListByName
func (mr *MockApiServiceServerMockRecorder) GetFollowingListByName(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowingListByName", reflect.TypeOf((*MockApiServiceServer)(nil).GetFollowingListByName), arg0, arg1)
}

// GetFollowCountByName mocks base method
func (m *MockApiServiceServer) GetFollowCountByName(arg0 context.Context, arg1 *pb.GetFollowCountByNameRequest) (*pb.GetFollowCountByNameResponse, error) {
	ret := m.ctrl.Call(m, "GetFollowCountByName", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetFollowCountByNameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowCountByName indicates an expected call of GetFollowCountByName
func (mr *MockApiServiceServerMockRecorder) GetFollowCountByName(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowCountByName", reflect.TypeOf((*MockApiServiceServer)(nil).GetFollowCountByName), arg0, arg1)
}

// GetWitnessList mocks base method
func (m *MockApiServiceServer) GetWitnessList(arg0 context.Context, arg1 *pb.GetWitnessListRequest) (*pb.GetWitnessListResponse, error) {
	ret := m.ctrl.Call(m, "GetWitnessList", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetWitnessListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWitnessList indicates an expected call of GetWitnessList
func (mr *MockApiServiceServerMockRecorder) GetWitnessList(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWitnessList", reflect.TypeOf((*MockApiServiceServer)(nil).GetWitnessList), arg0, arg1)
}

// GetPostListByCreated mocks base method
func (m *MockApiServiceServer) GetPostListByCreated(arg0 context.Context, arg1 *pb.GetPostListByCreatedRequest) (*pb.GetPostListByCreatedResponse, error) {
	ret := m.ctrl.Call(m, "GetPostListByCreated", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetPostListByCreatedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostListByCreated indicates an expected call of GetPostListByCreated
func (mr *MockApiServiceServerMockRecorder) GetPostListByCreated(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostListByCreated", reflect.TypeOf((*MockApiServiceServer)(nil).GetPostListByCreated), arg0, arg1)
}

// GetReplyListByPostId mocks base method
func (m *MockApiServiceServer) GetReplyListByPostId(arg0 context.Context, arg1 *pb.GetReplyListByPostIdRequest) (*pb.GetReplyListByPostIdResponse, error) {
	ret := m.ctrl.Call(m, "GetReplyListByPostId", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetReplyListByPostIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReplyListByPostId indicates an expected call of GetReplyListByPostId
func (mr *MockApiServiceServerMockRecorder) GetReplyListByPostId(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplyListByPostId", reflect.TypeOf((*MockApiServiceServer)(nil).GetReplyListByPostId), arg0, arg1)
}

// GetBlockTransactionsByNum mocks base method
func (m *MockApiServiceServer) GetBlockTransactionsByNum(arg0 context.Context, arg1 *pb.GetBlockTransactionsByNumRequest) (*pb.GetBlockTransactionsByNumResponse, error) {
	ret := m.ctrl.Call(m, "GetBlockTransactionsByNum", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetBlockTransactionsByNumResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockTransactionsByNum indicates an expected call of GetBlockTransactionsByNum
func (mr *MockApiServiceServerMockRecorder) GetBlockTransactionsByNum(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockTransactionsByNum", reflect.TypeOf((*MockApiServiceServer)(nil).GetBlockTransactionsByNum), arg0, arg1)
}

// GetTrxById mocks base method
func (m *MockApiServiceServer) GetTrxById(arg0 context.Context, arg1 *pb.GetTrxByIdRequest) (*pb.GetTrxByIdResponse, error) {
	ret := m.ctrl.Call(m, "GetTrxById", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetTrxByIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrxById indicates an expected call of GetTrxById
func (mr *MockApiServiceServerMockRecorder) GetTrxById(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrxById", reflect.TypeOf((*MockApiServiceServer)(nil).GetTrxById), arg0, arg1)
}

// GetChainState mocks base method
func (m *MockApiServiceServer) GetChainState(arg0 context.Context, arg1 *pb.NonParamsRequest) (*pb.GetChainStateResponse, error) {
	ret := m.ctrl.Call(m, "GetChainState", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetChainStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainState indicates an expected call of GetChainState
func (mr *MockApiServiceServerMockRecorder) GetChainState(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainState", reflect.TypeOf((*MockApiServiceServer)(nil).GetChainState), arg0, arg1)
}

// BroadcastTrx mocks base method
func (m *MockApiServiceServer) BroadcastTrx(arg0 context.Context, arg1 *pb.BroadcastTrxRequest) (*pb.BroadcastTrxResponse, error) {
	ret := m.ctrl.Call(m, "BroadcastTrx", arg0, arg1)
	ret0, _ := ret[0].(*pb.BroadcastTrxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BroadcastTrx indicates an expected call of BroadcastTrx
func (mr *MockApiServiceServerMockRecorder) BroadcastTrx(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastTrx", reflect.TypeOf((*MockApiServiceServer)(nil).BroadcastTrx), arg0, arg1)
}
