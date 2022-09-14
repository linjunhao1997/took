package account

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"time"
	"took/pkg/account/accountpb"
	"took/pkg/account/user"
	"took/pkg/account/util"
)

func NewGrpcService(callTimeout time.Duration) (Service, error) {

	cli, err := clientv3.NewFromURL("http://192.168.31.117:2379")
	if err != nil {
		return nil, err
	}
	builder, err := resolver.NewBuilder(cli)
	if err != nil {
		return nil, err
	}
	s := &GrpcService{callTimeout: callTimeout}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	s.connection, err = grpc.DialContext(ctx, "etcd:///took/account",
		grpc.WithInsecure(),
		grpc.WithResolvers(builder),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
	)
	if err != nil {
		return nil, err
	}

	s.accountClient = accountpb.NewAccountServiceClient(s.connection)
	return s, nil
}

type GrpcService struct {
	accountClient accountpb.AccountServiceClient
	connection    *grpc.ClientConn
	callTimeout   time.Duration
}

func (g *GrpcService) LoadUser(ctx context.Context, id int) (*user.User, error) {
	ctx, cancel := context.WithTimeout(ctx, g.callTimeout)
	defer cancel()

	request := &accountpb.LoadUserRequest{
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
