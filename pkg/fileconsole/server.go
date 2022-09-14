package fileconsole

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"took/pkg/fileconsole/fileconsolepb"
)

type GrpcServer struct {
	loadFileHandler  grpc.Handler
	loadFilesHandler grpc.Handler
}

func (s *GrpcServer) LoadFile(ctx context.Context, req *fileconsolepb.LoadFileRequest) (*fileconsolepb.FileResponse, error) {
	_, rsp, err := s.loadFileHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*fileconsolepb.FileResponse), err
}

func (s *GrpcServer) LoadFiles(ctx context.Context, req *fileconsolepb.LoadFilesRequest) (*fileconsolepb.FilesResponse, error) {
	_, rsp, err := s.loadFilesHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*fileconsolepb.FilesResponse), err
}

func NewGrpcServiceServer(svc Service, opts ...grpc.ServerOption) fileconsolepb.FileServiceServer {

	loadFileHandler := grpc.NewServer(
		makeLoadFileEndpoint(svc),
		decodeLoadFileRequestByGrpc,
		encodeLoadFileResponseByGrpc,
		opts...,
	)

	grpcServer := new(GrpcServer)
	grpcServer.loadFileHandler = loadFileHandler

	return grpcServer
}
