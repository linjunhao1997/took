package grpc

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	fileconsole "took/pkg/fileconsole/api/v1/grpc/proto"
	"took/pkg/fileconsole/endpoint"
	"took/pkg/fileconsole/service"
)

type GrpcServer struct {
	loadFileHandler  grpc.Handler
	loadFilesHandler grpc.Handler
}

func (s *GrpcServer) LoadFile(ctx context.Context, req *fileconsole.LoadFileRequest) (*fileconsole.FileResponse, error) {
	_, rsp, err := s.loadFileHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*fileconsole.FileResponse), err
}

func (s *GrpcServer) LoadFiles(ctx context.Context, req *fileconsole.LoadFilesRequest) (*fileconsole.FilesResponse, error) {
	_, rsp, err := s.loadFilesHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*fileconsole.FilesResponse), err
}

func NewGrpcServiceServer(svc service.Service, opts ...grpc.ServerOption) fileconsole.FileServiceServer {

	loadFileHandler := grpc.NewServer(
		endpoint.MakeLoadFileEndpoint(svc),
		decodeLoadFileRequestByGrpc,
		encodeLoadFileResponseByGrpc,
		opts...,
	)

	grpcServer := new(GrpcServer)
	grpcServer.loadFileHandler = loadFileHandler

	return grpcServer
}
