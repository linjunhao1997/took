package fileconsole

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"runtime/debug"
	"time"
	grpc_account "took/pkg/account/api/v1/grpc"
	grpc_fileconsole "took/pkg/fileconsole/api/v1/grpc"
	fileconsole "took/pkg/fileconsole/api/v1/grpc/proto"
	"took/pkg/fileconsole/domain/file"
	"took/pkg/fileconsole/service"
	"took/pkg/util"
	"took/pkg/util/discovery"
)

func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()

	return handler(ctx, req)
}

func runWeb(addr string, errc <-chan error) {

}

func RunGrpc(srvName, addr, port string, errCh chan<- error) {
	var opts = []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			RecoveryInterceptor,
		),
	}

	var grpcServer = grpc.NewServer(opts...)
	db, err := util.NewDB("root:123456@tcp(192.168.100.100:3306)/fileconsole?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		errCh <- err
		return
	}

	if err = discovery.Register(srvName, addr, port); err != nil {
		errCh <- err
		return
	}
	accountService, err := grpc_account.NewGrpcService(3 * time.Second)
	if err != nil {
		errCh <- err
		return
	}
	fileconsole.RegisterFileServiceServer(grpcServer, grpc_fileconsole.NewGrpcServiceServer(service.NewFileConsoleService(file.NewFileRepository(db), accountService)))

	listener, err := net.Listen("tcp", addr+":"+port)
	if err != nil {
		errCh <- err
		return
	}

	if err = grpcServer.Serve(listener); err != nil {
		errCh <- err
		return
	}
}
