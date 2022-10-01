package grpc

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"time"
	account "took/pkg/account/api/v1/grpc/proto"
	"took/pkg/account/domain/user"
	"took/pkg/account/service"
	"took/pkg/account/util"
)

func NewGrpcService(callTimeout time.Duration) (service.AccountService, error) {

	clientConfig := clientv3.Config{
		Endpoints:   []string{"http://192.168.1.101:2379"},
		DialTimeout: 2 * time.Second,
	}

	cli, err := clientv3.New(clientConfig)
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to etcd")
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = cli.Status(timeoutCtx, clientConfig.Endpoints[0])
	if err != nil {
		return nil, errors.Wrapf(err, "error checking etcd status: %v", err)
	}

	builder, err := resolver.NewBuilder(cli)
	if err != nil {
		return nil, err
	}
	s := &GrpcService{callTimeout: callTimeout}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	s.connection, err = grpc.DialContext(ctx, "etcd:///took/grpc_account",
		grpc.WithInsecure(),
		grpc.WithResolvers(builder),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	s.accountClient = account.NewAccountServiceClient(s.connection)
	return s, nil
}

type GrpcService struct {
	accountClient account.AccountServiceClient
	connection    *grpc.ClientConn
	callTimeout   time.Duration
}

func (g *GrpcService) LoadUser(ctx context.Context, id int) (*user.User, error) {
	ctx, cancel := context.WithTimeout(ctx, g.callTimeout)
	defer cancel()

	request := &account.LoadUserRequest{
		Id: int32(id),
	}
	resp, err := g.accountClient.LoadUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return util.ConvertFormUserProto(resp.Data), nil
}

func (g *GrpcService) LoadUsers(ctx context.Context, id ...int) ([]*user.User, error) {
	return nil, nil
}
