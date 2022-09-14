package account

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"took/pkg/account/accountpb"
)

type GrpcServer struct {
	loadUserHandler  grpc.Handler
	loadUsersHandler grpc.Handler
}

func (s *GrpcServer) LoadUser(ctx context.Context, req *accountpb.LoadUserRequest) (*accountpb.UserResponse, error) {
	_, rsp, err := s.loadUserHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*accountpb.UserResponse), err
}

func (s *GrpcServer) LoadUsers(ctx context.Context, req *accountpb.LoadUsersRequest) (*accountpb.UsersResponse, error) {
	_, rsp, err := s.loadUsersHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*accountpb.UsersResponse), err
}

func NewGrpcServiceServer(svc Service, opts ...grpc.ServerOption) accountpb.AccountServiceServer {

	loadUserHandler := grpc.NewServer(
		makeLoadUserEndpoint(svc),
		decodeLoadUserRequestByGrpc,
		encodeLoadUserResponseByGrpc,
		opts...,
	)

	grpcServer := new(GrpcServer)
	grpcServer.loadUserHandler = loadUserHandler

	return grpcServer
}
