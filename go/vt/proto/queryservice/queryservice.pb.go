// Code generated by protoc-gen-go. DO NOT EDIT.
// source: queryservice.proto

package queryservice // import "vitess.io/vitess/go/vt/proto/queryservice"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import binlogdata "vitess.io/vitess/go/vt/proto/binlogdata"
import query "vitess.io/vitess/go/vt/proto/query"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Execute executes the specified SQL query (might be in a
	// transaction context, if Query.transaction_id is set).
	Execute(ctx context.Context, in *query.ExecuteRequest, opts ...grpc.CallOption) (*query.ExecuteResponse, error)
	// ExecuteBatch executes a list of queries, and returns the result
	// for each query.
	ExecuteBatch(ctx context.Context, in *query.ExecuteBatchRequest, opts ...grpc.CallOption) (*query.ExecuteBatchResponse, error)
	// StreamExecute executes a streaming query. Use this method if the
	// query returns a large number of rows. The first QueryResult will
	// contain the Fields, subsequent QueryResult messages will contain
	// the rows.
	StreamExecute(ctx context.Context, in *query.StreamExecuteRequest, opts ...grpc.CallOption) (Query_StreamExecuteClient, error)
	// Begin a transaction.
	Begin(ctx context.Context, in *query.BeginRequest, opts ...grpc.CallOption) (*query.BeginResponse, error)
	// Commit a transaction.
	Commit(ctx context.Context, in *query.CommitRequest, opts ...grpc.CallOption) (*query.CommitResponse, error)
	// Rollback a transaction.
	Rollback(ctx context.Context, in *query.RollbackRequest, opts ...grpc.CallOption) (*query.RollbackResponse, error)
	// Prepare preares a transaction.
	Prepare(ctx context.Context, in *query.PrepareRequest, opts ...grpc.CallOption) (*query.PrepareResponse, error)
	// CommitPrepared commits a prepared transaction.
	CommitPrepared(ctx context.Context, in *query.CommitPreparedRequest, opts ...grpc.CallOption) (*query.CommitPreparedResponse, error)
	// RollbackPrepared rolls back a prepared transaction.
	RollbackPrepared(ctx context.Context, in *query.RollbackPreparedRequest, opts ...grpc.CallOption) (*query.RollbackPreparedResponse, error)
	// CreateTransaction creates the metadata for a 2pc transaction.
	CreateTransaction(ctx context.Context, in *query.CreateTransactionRequest, opts ...grpc.CallOption) (*query.CreateTransactionResponse, error)
	// StartCommit initiates a commit for a 2pc transaction.
	StartCommit(ctx context.Context, in *query.StartCommitRequest, opts ...grpc.CallOption) (*query.StartCommitResponse, error)
	// SetRollback marks the 2pc transaction for rollback.
	SetRollback(ctx context.Context, in *query.SetRollbackRequest, opts ...grpc.CallOption) (*query.SetRollbackResponse, error)
	// ConcludeTransaction marks the 2pc transaction as resolved.
	ConcludeTransaction(ctx context.Context, in *query.ConcludeTransactionRequest, opts ...grpc.CallOption) (*query.ConcludeTransactionResponse, error)
	// ReadTransaction returns the 2pc transaction info.
	ReadTransaction(ctx context.Context, in *query.ReadTransactionRequest, opts ...grpc.CallOption) (*query.ReadTransactionResponse, error)
	// BeginExecute executes a begin and the specified SQL query.
	BeginExecute(ctx context.Context, in *query.BeginExecuteRequest, opts ...grpc.CallOption) (*query.BeginExecuteResponse, error)
	// BeginExecuteBatch executes a begin and a list of queries.
	BeginExecuteBatch(ctx context.Context, in *query.BeginExecuteBatchRequest, opts ...grpc.CallOption) (*query.BeginExecuteBatchResponse, error)
	// MessageStream streams messages from a message table.
	MessageStream(ctx context.Context, in *query.MessageStreamRequest, opts ...grpc.CallOption) (Query_MessageStreamClient, error)
	// MessageAck acks messages for a table.
	MessageAck(ctx context.Context, in *query.MessageAckRequest, opts ...grpc.CallOption) (*query.MessageAckResponse, error)
	// SplitQuery is the API to facilitate MapReduce-type iterations
	// over large data sets (like full table dumps).
	SplitQuery(ctx context.Context, in *query.SplitQueryRequest, opts ...grpc.CallOption) (*query.SplitQueryResponse, error)
	// StreamHealth runs a streaming RPC to the tablet, that returns the
	// current health of the tablet on a regular basis.
	StreamHealth(ctx context.Context, in *query.StreamHealthRequest, opts ...grpc.CallOption) (Query_StreamHealthClient, error)
	// UpdateStream asks the server to return a stream of the updates that have been applied to its database.
	UpdateStream(ctx context.Context, in *query.UpdateStreamRequest, opts ...grpc.CallOption) (Query_UpdateStreamClient, error)
	// VStream streams vreplication events.
	VStream(ctx context.Context, in *binlogdata.VStreamRequest, opts ...grpc.CallOption) (Query_VStreamClient, error)
}

type queryClient struct {
	cc *grpc.ClientConn
}

func NewQueryClient(cc *grpc.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Execute(ctx context.Context, in *query.ExecuteRequest, opts ...grpc.CallOption) (*query.ExecuteResponse, error) {
	out := new(query.ExecuteResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ExecuteBatch(ctx context.Context, in *query.ExecuteBatchRequest, opts ...grpc.CallOption) (*query.ExecuteBatchResponse, error) {
	out := new(query.ExecuteBatchResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/ExecuteBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StreamExecute(ctx context.Context, in *query.StreamExecuteRequest, opts ...grpc.CallOption) (Query_StreamExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Query_serviceDesc.Streams[0], "/queryservice.Query/StreamExecute", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryStreamExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_StreamExecuteClient interface {
	Recv() (*query.StreamExecuteResponse, error)
	grpc.ClientStream
}

type queryStreamExecuteClient struct {
	grpc.ClientStream
}

func (x *queryStreamExecuteClient) Recv() (*query.StreamExecuteResponse, error) {
	m := new(query.StreamExecuteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryClient) Begin(ctx context.Context, in *query.BeginRequest, opts ...grpc.CallOption) (*query.BeginResponse, error) {
	out := new(query.BeginResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/Begin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Commit(ctx context.Context, in *query.CommitRequest, opts ...grpc.CallOption) (*query.CommitResponse, error) {
	out := new(query.CommitResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Rollback(ctx context.Context, in *query.RollbackRequest, opts ...grpc.CallOption) (*query.RollbackResponse, error) {
	out := new(query.RollbackResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/Rollback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Prepare(ctx context.Context, in *query.PrepareRequest, opts ...grpc.CallOption) (*query.PrepareResponse, error) {
	out := new(query.PrepareResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/Prepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CommitPrepared(ctx context.Context, in *query.CommitPreparedRequest, opts ...grpc.CallOption) (*query.CommitPreparedResponse, error) {
	out := new(query.CommitPreparedResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/CommitPrepared", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RollbackPrepared(ctx context.Context, in *query.RollbackPreparedRequest, opts ...grpc.CallOption) (*query.RollbackPreparedResponse, error) {
	out := new(query.RollbackPreparedResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/RollbackPrepared", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CreateTransaction(ctx context.Context, in *query.CreateTransactionRequest, opts ...grpc.CallOption) (*query.CreateTransactionResponse, error) {
	out := new(query.CreateTransactionResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StartCommit(ctx context.Context, in *query.StartCommitRequest, opts ...grpc.CallOption) (*query.StartCommitResponse, error) {
	out := new(query.StartCommitResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/StartCommit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SetRollback(ctx context.Context, in *query.SetRollbackRequest, opts ...grpc.CallOption) (*query.SetRollbackResponse, error) {
	out := new(query.SetRollbackResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/SetRollback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ConcludeTransaction(ctx context.Context, in *query.ConcludeTransactionRequest, opts ...grpc.CallOption) (*query.ConcludeTransactionResponse, error) {
	out := new(query.ConcludeTransactionResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/ConcludeTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ReadTransaction(ctx context.Context, in *query.ReadTransactionRequest, opts ...grpc.CallOption) (*query.ReadTransactionResponse, error) {
	out := new(query.ReadTransactionResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/ReadTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BeginExecute(ctx context.Context, in *query.BeginExecuteRequest, opts ...grpc.CallOption) (*query.BeginExecuteResponse, error) {
	out := new(query.BeginExecuteResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/BeginExecute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BeginExecuteBatch(ctx context.Context, in *query.BeginExecuteBatchRequest, opts ...grpc.CallOption) (*query.BeginExecuteBatchResponse, error) {
	out := new(query.BeginExecuteBatchResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/BeginExecuteBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MessageStream(ctx context.Context, in *query.MessageStreamRequest, opts ...grpc.CallOption) (Query_MessageStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Query_serviceDesc.Streams[1], "/queryservice.Query/MessageStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryMessageStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_MessageStreamClient interface {
	Recv() (*query.MessageStreamResponse, error)
	grpc.ClientStream
}

type queryMessageStreamClient struct {
	grpc.ClientStream
}

func (x *queryMessageStreamClient) Recv() (*query.MessageStreamResponse, error) {
	m := new(query.MessageStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryClient) MessageAck(ctx context.Context, in *query.MessageAckRequest, opts ...grpc.CallOption) (*query.MessageAckResponse, error) {
	out := new(query.MessageAckResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/MessageAck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SplitQuery(ctx context.Context, in *query.SplitQueryRequest, opts ...grpc.CallOption) (*query.SplitQueryResponse, error) {
	out := new(query.SplitQueryResponse)
	err := c.cc.Invoke(ctx, "/queryservice.Query/SplitQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StreamHealth(ctx context.Context, in *query.StreamHealthRequest, opts ...grpc.CallOption) (Query_StreamHealthClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Query_serviceDesc.Streams[2], "/queryservice.Query/StreamHealth", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryStreamHealthClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_StreamHealthClient interface {
	Recv() (*query.StreamHealthResponse, error)
	grpc.ClientStream
}

type queryStreamHealthClient struct {
	grpc.ClientStream
}

func (x *queryStreamHealthClient) Recv() (*query.StreamHealthResponse, error) {
	m := new(query.StreamHealthResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryClient) UpdateStream(ctx context.Context, in *query.UpdateStreamRequest, opts ...grpc.CallOption) (Query_UpdateStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Query_serviceDesc.Streams[3], "/queryservice.Query/UpdateStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryUpdateStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_UpdateStreamClient interface {
	Recv() (*query.UpdateStreamResponse, error)
	grpc.ClientStream
}

type queryUpdateStreamClient struct {
	grpc.ClientStream
}

func (x *queryUpdateStreamClient) Recv() (*query.UpdateStreamResponse, error) {
	m := new(query.UpdateStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryClient) VStream(ctx context.Context, in *binlogdata.VStreamRequest, opts ...grpc.CallOption) (Query_VStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Query_serviceDesc.Streams[4], "/queryservice.Query/VStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryVStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_VStreamClient interface {
	Recv() (*binlogdata.VStreamResponse, error)
	grpc.ClientStream
}

type queryVStreamClient struct {
	grpc.ClientStream
}

func (x *queryVStreamClient) Recv() (*binlogdata.VStreamResponse, error) {
	m := new(binlogdata.VStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Execute executes the specified SQL query (might be in a
	// transaction context, if Query.transaction_id is set).
	Execute(context.Context, *query.ExecuteRequest) (*query.ExecuteResponse, error)
	// ExecuteBatch executes a list of queries, and returns the result
	// for each query.
	ExecuteBatch(context.Context, *query.ExecuteBatchRequest) (*query.ExecuteBatchResponse, error)
	// StreamExecute executes a streaming query. Use this method if the
	// query returns a large number of rows. The first QueryResult will
	// contain the Fields, subsequent QueryResult messages will contain
	// the rows.
	StreamExecute(*query.StreamExecuteRequest, Query_StreamExecuteServer) error
	// Begin a transaction.
	Begin(context.Context, *query.BeginRequest) (*query.BeginResponse, error)
	// Commit a transaction.
	Commit(context.Context, *query.CommitRequest) (*query.CommitResponse, error)
	// Rollback a transaction.
	Rollback(context.Context, *query.RollbackRequest) (*query.RollbackResponse, error)
	// Prepare preares a transaction.
	Prepare(context.Context, *query.PrepareRequest) (*query.PrepareResponse, error)
	// CommitPrepared commits a prepared transaction.
	CommitPrepared(context.Context, *query.CommitPreparedRequest) (*query.CommitPreparedResponse, error)
	// RollbackPrepared rolls back a prepared transaction.
	RollbackPrepared(context.Context, *query.RollbackPreparedRequest) (*query.RollbackPreparedResponse, error)
	// CreateTransaction creates the metadata for a 2pc transaction.
	CreateTransaction(context.Context, *query.CreateTransactionRequest) (*query.CreateTransactionResponse, error)
	// StartCommit initiates a commit for a 2pc transaction.
	StartCommit(context.Context, *query.StartCommitRequest) (*query.StartCommitResponse, error)
	// SetRollback marks the 2pc transaction for rollback.
	SetRollback(context.Context, *query.SetRollbackRequest) (*query.SetRollbackResponse, error)
	// ConcludeTransaction marks the 2pc transaction as resolved.
	ConcludeTransaction(context.Context, *query.ConcludeTransactionRequest) (*query.ConcludeTransactionResponse, error)
	// ReadTransaction returns the 2pc transaction info.
	ReadTransaction(context.Context, *query.ReadTransactionRequest) (*query.ReadTransactionResponse, error)
	// BeginExecute executes a begin and the specified SQL query.
	BeginExecute(context.Context, *query.BeginExecuteRequest) (*query.BeginExecuteResponse, error)
	// BeginExecuteBatch executes a begin and a list of queries.
	BeginExecuteBatch(context.Context, *query.BeginExecuteBatchRequest) (*query.BeginExecuteBatchResponse, error)
	// MessageStream streams messages from a message table.
	MessageStream(*query.MessageStreamRequest, Query_MessageStreamServer) error
	// MessageAck acks messages for a table.
	MessageAck(context.Context, *query.MessageAckRequest) (*query.MessageAckResponse, error)
	// SplitQuery is the API to facilitate MapReduce-type iterations
	// over large data sets (like full table dumps).
	SplitQuery(context.Context, *query.SplitQueryRequest) (*query.SplitQueryResponse, error)
	// StreamHealth runs a streaming RPC to the tablet, that returns the
	// current health of the tablet on a regular basis.
	StreamHealth(*query.StreamHealthRequest, Query_StreamHealthServer) error
	// UpdateStream asks the server to return a stream of the updates that have been applied to its database.
	UpdateStream(*query.UpdateStreamRequest, Query_UpdateStreamServer) error
	// VStream streams vreplication events.
	VStream(*binlogdata.VStreamRequest, Query_VStreamServer) error
}

func RegisterQueryServer(s *grpc.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Execute(ctx, req.(*query.ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ExecuteBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.ExecuteBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ExecuteBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/ExecuteBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ExecuteBatch(ctx, req.(*query.ExecuteBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_StreamExecute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(query.StreamExecuteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).StreamExecute(m, &queryStreamExecuteServer{stream})
}

type Query_StreamExecuteServer interface {
	Send(*query.StreamExecuteResponse) error
	grpc.ServerStream
}

type queryStreamExecuteServer struct {
	grpc.ServerStream
}

func (x *queryStreamExecuteServer) Send(m *query.StreamExecuteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Query_Begin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.BeginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Begin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/Begin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Begin(ctx, req.(*query.BeginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Commit(ctx, req.(*query.CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Rollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.RollbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Rollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/Rollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Rollback(ctx, req.(*query.RollbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.PrepareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Prepare(ctx, req.(*query.PrepareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CommitPrepared_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.CommitPreparedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CommitPrepared(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/CommitPrepared",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CommitPrepared(ctx, req.(*query.CommitPreparedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RollbackPrepared_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.RollbackPreparedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RollbackPrepared(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/RollbackPrepared",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RollbackPrepared(ctx, req.(*query.RollbackPreparedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CreateTransaction(ctx, req.(*query.CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_StartCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.StartCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).StartCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/StartCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).StartCommit(ctx, req.(*query.StartCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SetRollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.SetRollbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SetRollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/SetRollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SetRollback(ctx, req.(*query.SetRollbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ConcludeTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.ConcludeTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ConcludeTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/ConcludeTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ConcludeTransaction(ctx, req.(*query.ConcludeTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ReadTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.ReadTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ReadTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/ReadTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ReadTransaction(ctx, req.(*query.ReadTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BeginExecute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.BeginExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BeginExecute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/BeginExecute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BeginExecute(ctx, req.(*query.BeginExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BeginExecuteBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.BeginExecuteBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BeginExecuteBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/BeginExecuteBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BeginExecuteBatch(ctx, req.(*query.BeginExecuteBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(query.MessageStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).MessageStream(m, &queryMessageStreamServer{stream})
}

type Query_MessageStreamServer interface {
	Send(*query.MessageStreamResponse) error
	grpc.ServerStream
}

type queryMessageStreamServer struct {
	grpc.ServerStream
}

func (x *queryMessageStreamServer) Send(m *query.MessageStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Query_MessageAck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.MessageAckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MessageAck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/MessageAck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MessageAck(ctx, req.(*query.MessageAckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SplitQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(query.SplitQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SplitQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryservice.Query/SplitQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SplitQuery(ctx, req.(*query.SplitQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_StreamHealth_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(query.StreamHealthRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).StreamHealth(m, &queryStreamHealthServer{stream})
}

type Query_StreamHealthServer interface {
	Send(*query.StreamHealthResponse) error
	grpc.ServerStream
}

type queryStreamHealthServer struct {
	grpc.ServerStream
}

func (x *queryStreamHealthServer) Send(m *query.StreamHealthResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Query_UpdateStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(query.UpdateStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).UpdateStream(m, &queryUpdateStreamServer{stream})
}

type Query_UpdateStreamServer interface {
	Send(*query.UpdateStreamResponse) error
	grpc.ServerStream
}

type queryUpdateStreamServer struct {
	grpc.ServerStream
}

func (x *queryUpdateStreamServer) Send(m *query.UpdateStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Query_VStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(binlogdata.VStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).VStream(m, &queryVStreamServer{stream})
}

type Query_VStreamServer interface {
	Send(*binlogdata.VStreamResponse) error
	grpc.ServerStream
}

type queryVStreamServer struct {
	grpc.ServerStream
}

func (x *queryVStreamServer) Send(m *binlogdata.VStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "queryservice.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _Query_Execute_Handler,
		},
		{
			MethodName: "ExecuteBatch",
			Handler:    _Query_ExecuteBatch_Handler,
		},
		{
			MethodName: "Begin",
			Handler:    _Query_Begin_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _Query_Commit_Handler,
		},
		{
			MethodName: "Rollback",
			Handler:    _Query_Rollback_Handler,
		},
		{
			MethodName: "Prepare",
			Handler:    _Query_Prepare_Handler,
		},
		{
			MethodName: "CommitPrepared",
			Handler:    _Query_CommitPrepared_Handler,
		},
		{
			MethodName: "RollbackPrepared",
			Handler:    _Query_RollbackPrepared_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _Query_CreateTransaction_Handler,
		},
		{
			MethodName: "StartCommit",
			Handler:    _Query_StartCommit_Handler,
		},
		{
			MethodName: "SetRollback",
			Handler:    _Query_SetRollback_Handler,
		},
		{
			MethodName: "ConcludeTransaction",
			Handler:    _Query_ConcludeTransaction_Handler,
		},
		{
			MethodName: "ReadTransaction",
			Handler:    _Query_ReadTransaction_Handler,
		},
		{
			MethodName: "BeginExecute",
			Handler:    _Query_BeginExecute_Handler,
		},
		{
			MethodName: "BeginExecuteBatch",
			Handler:    _Query_BeginExecuteBatch_Handler,
		},
		{
			MethodName: "MessageAck",
			Handler:    _Query_MessageAck_Handler,
		},
		{
			MethodName: "SplitQuery",
			Handler:    _Query_SplitQuery_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamExecute",
			Handler:       _Query_StreamExecute_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "MessageStream",
			Handler:       _Query_MessageStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamHealth",
			Handler:       _Query_StreamHealth_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateStream",
			Handler:       _Query_UpdateStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "VStream",
			Handler:       _Query_VStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "queryservice.proto",
}

func init() { proto.RegisterFile("queryservice.proto", fileDescriptor_queryservice_17509881eb07629d) }

var fileDescriptor_queryservice_17509881eb07629d = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x95, 0xdf, 0x6b, 0xd4, 0x40,
	0x10, 0xc7, 0xf5, 0xa1, 0xad, 0x4c, 0x4f, 0xad, 0x5b, 0xab, 0x36, 0xad, 0x6d, 0xed, 0x9b, 0x08,
	0x17, 0x51, 0x41, 0x28, 0xf8, 0xd0, 0x0b, 0x16, 0xa5, 0xf8, 0xeb, 0xce, 0x16, 0xf1, 0x41, 0xd8,
	0x4b, 0x86, 0x33, 0x34, 0x97, 0x4d, 0x93, 0xbd, 0x43, 0xff, 0x6a, 0xff, 0x05, 0x31, 0x9b, 0x99,
	0xec, 0xee, 0x25, 0xbe, 0xdd, 0x7e, 0xbf, 0x33, 0x1f, 0x26, 0x3b, 0x37, 0xb3, 0x20, 0xae, 0x17,
	0x58, 0xfe, 0xae, 0xb0, 0x5c, 0xa6, 0x31, 0x0e, 0x8b, 0x52, 0x69, 0x25, 0x06, 0xb6, 0x16, 0x6c,
	0xd6, 0x27, 0x63, 0x05, 0x5b, 0xd3, 0x34, 0xcf, 0xd4, 0x2c, 0x91, 0x5a, 0x1a, 0xe5, 0xc5, 0x9f,
	0x01, 0xac, 0x7d, 0xf9, 0x17, 0x21, 0x4e, 0x60, 0xe3, 0xed, 0x2f, 0x8c, 0x17, 0x1a, 0xc5, 0xce,
	0xd0, 0x24, 0x35, 0xe7, 0x31, 0x5e, 0x2f, 0xb0, 0xd2, 0xc1, 0x03, 0x5f, 0xae, 0x0a, 0x95, 0x57,
	0x78, 0x7c, 0x43, 0xbc, 0x87, 0x41, 0x23, 0x8e, 0xa4, 0x8e, 0x7f, 0x8a, 0xc0, 0x8d, 0xac, 0x45,
	0xa2, 0xec, 0x75, 0x7a, 0x8c, 0xfa, 0x08, 0xb7, 0x27, 0xba, 0x44, 0x39, 0xa7, 0x62, 0x28, 0xde,
	0x51, 0x09, 0xb6, 0xdf, 0x6d, 0x12, 0xed, 0xf9, 0x4d, 0xf1, 0x0a, 0xd6, 0x46, 0x38, 0x4b, 0x73,
	0xb1, 0xdd, 0x84, 0xd6, 0x27, 0xca, 0xbf, 0xef, 0x8a, 0x5c, 0xc5, 0x6b, 0x58, 0x8f, 0xd4, 0x7c,
	0x9e, 0x6a, 0x41, 0x11, 0xe6, 0x48, 0x79, 0x3b, 0x9e, 0xca, 0x89, 0x6f, 0xe0, 0xd6, 0x58, 0x65,
	0xd9, 0x54, 0xc6, 0x57, 0x82, 0xee, 0x8b, 0x04, 0x4a, 0x7e, 0xb8, 0xa2, 0x73, 0xfa, 0x09, 0x6c,
	0x7c, 0x2e, 0xb1, 0x90, 0x65, 0xdb, 0x84, 0xe6, 0xec, 0x37, 0x81, 0x65, 0xce, 0xfd, 0x04, 0x77,
	0x4c, 0x39, 0x8d, 0x95, 0x88, 0x7d, 0xa7, 0x4a, 0x92, 0x89, 0xf4, 0xb8, 0xc7, 0x65, 0xe0, 0x05,
	0x6c, 0x51, 0x89, 0x8c, 0x3c, 0xf0, 0x6a, 0xf7, 0xa1, 0x87, 0xbd, 0x3e, 0x63, 0xbf, 0xc1, 0xbd,
	0xa8, 0x44, 0xa9, 0xf1, 0x6b, 0x29, 0xf3, 0x4a, 0xc6, 0x3a, 0x55, 0xb9, 0xa0, 0xbc, 0x15, 0x87,
	0xc0, 0x47, 0xfd, 0x01, 0x4c, 0x3e, 0x83, 0xcd, 0x89, 0x96, 0xa5, 0x6e, 0x5a, 0xb7, 0xcb, 0x7f,
	0x0e, 0xd6, 0x88, 0x16, 0x74, 0x59, 0x0e, 0x07, 0x35, 0xf7, 0x91, 0x39, 0xad, 0xb6, 0xc2, 0xb1,
	0x2d, 0xe6, 0xfc, 0x80, 0xed, 0x48, 0xe5, 0x71, 0xb6, 0x48, 0x9c, 0x6f, 0x7d, 0xc2, 0x17, 0xbf,
	0xe2, 0x11, 0xf7, 0xf8, 0x7f, 0x21, 0xcc, 0x1f, 0xc3, 0xdd, 0x31, 0xca, 0xc4, 0x66, 0x53, 0x53,
	0x3d, 0x9d, 0xb8, 0x07, 0x7d, 0xb6, 0x3d, 0xca, 0xf5, 0x30, 0xd0, 0xf8, 0x05, 0xf6, 0x84, 0x78,
	0xd3, 0xb7, 0xd7, 0xe9, 0xd9, 0x8d, 0xb6, 0x1d, 0xb3, 0x1a, 0x0e, 0x3b, 0x72, 0x9c, 0xfd, 0x70,
	0xd4, 0x1f, 0x60, 0x2f, 0x89, 0x0f, 0x58, 0x55, 0x72, 0x86, 0x66, 0xf0, 0x79, 0x49, 0x38, 0xaa,
	0xbf, 0x24, 0x3c, 0xd3, 0x5a, 0x12, 0x11, 0x40, 0x63, 0x9e, 0xc6, 0x57, 0xe2, 0x91, 0x1b, 0x7f,
	0xda, 0xb6, 0x7b, 0xb7, 0xc3, 0xe1, 0xa2, 0x22, 0x80, 0x49, 0x91, 0xa5, 0xda, 0xac, 0x53, 0x82,
	0xb4, 0x92, 0x0f, 0xb1, 0x1d, 0x86, 0x9c, 0xc3, 0xc0, 0xd4, 0xf7, 0x0e, 0x65, 0xa6, 0xdb, 0x4d,
	0x6a, 0x8b, 0xfe, 0xf5, 0xbb, 0x9e, 0xf5, 0x59, 0xe7, 0x30, 0xb8, 0x28, 0x12, 0xa9, 0xe9, 0x96,
	0x08, 0x66, 0x8b, 0x3e, 0xcc, 0xf5, 0x2c, 0xd8, 0x19, 0x6c, 0x5c, 0x32, 0xc7, 0x7a, 0x47, 0x2e,
	0x7d, 0x4e, 0x97, 0xd7, 0x72, 0x46, 0xcf, 0xbe, 0x3f, 0x5d, 0xa6, 0x1a, 0xab, 0x6a, 0x98, 0xaa,
	0xd0, 0xfc, 0x0a, 0x67, 0x2a, 0x5c, 0xea, 0xb0, 0x7e, 0x91, 0x42, 0xfb, 0xf5, 0x9a, 0xae, 0xd7,
	0xda, 0xcb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x47, 0x8e, 0x80, 0xe8, 0x06, 0x00, 0x00,
}
