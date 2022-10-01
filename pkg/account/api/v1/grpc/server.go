package grpc

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	account "took/pkg/account/api/v1/grpc/proto"
	"took/pkg/account/endpoint"
	"took/pkg/account/service"
)

type GrpcServer struct {
	loadUserHandler  grpc.Handler
	loadUsersHandler grpc.Handler
}

func (s *GrpcServer) LoadUser(ctx context.Context, req *account.LoadUserRequest) (*account.UserResponse, error) {
	_, rsp, err := s.loadUserHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*account.UserResponse), err
}

func (s *GrpcServer) LoadUsers(ctx context.Context, req *account.LoadUsersRequest) (*account.UsersResponse, error) {
	_, rsp, err := s.loadUsersHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*account.UsersResponse), err
}

func NewGrpcServiceServer(svc service.AccountService, opts ...grpc.ServerOption) account.AccountServiceServer {

	loadUserHandler := grpc.NewServer(
		endpoint.MakeLoadUserEndpoint(svc),
		decodeLoadUserRequestByGrpc,
		encodeLoadUserResponseByGrpc,
		opts...,
	)

	grpcServer := new(GrpcServer)
	grpcServer.loadUserHandler = loadUserHandler

	return grpcServer
}
