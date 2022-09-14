package discovery

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/log"
	"time"
)

func Register(srvName, addr, port string) error {
	etcdServer := "127.0.0.1:2379"                     //etcd服务的IP地址
	prefix := fmt.Sprintf("took/%s/", srvName)         //服务的目录
	serverInstance := fmt.Sprintf("%s:%s", addr, port) //当前实例Server的地址
	key := prefix + serverInstance                     //服务实例注册的路径
	value := fmt.Sprintf(`{"Addr":"%s"}`, serverInstance)
	ctx := context.Background()
	//etcd连接参数
	option := etcdv3.ClientOptions{DialTimeout: time.Second * 3, DialKeepAlive: time.Second * 3}
	//创建连接
	client, err := etcdv3.NewClient(ctx, []string{etcdServer}, option)
	if err != nil {
		return err
	}
	registrar := etcdv3.NewRegistrar(client, etcdv3.Service{Key: key, Value: value}, log.NewNopLogger())
	registrar.Register() //启动注册服务
	return nil
}
