package run

import (
	"fmt"
	"strings"
	"took/pkg/account"
	"took/pkg/fileconsole"
)

type runFuncMap map[string]func(srvName, addr, port string, errCh chan<- error)

var Funcs runFuncMap

func (runFunc runFuncMap) RunServer(srvName, addr, port, t string, errCh chan error) {
	t = strings.ToLower(t)
	if t == "http" {
		runFunc.runHttpServer(srvName, addr, port, errCh)
	} else if t == "grpc" {
		runFunc.runGrpcServer(srvName, addr, port, errCh)
	} else {
		errCh <- fmt.Errorf("no exist type: %s", t)
	}
}

func (runFunc runFuncMap) runHttpServer(srvName, addr, port string, errCh chan error) {
	srv := "http_" + srvName
	runFunc.runServer(srv, addr, port, errCh)
}

func (runFunc runFuncMap) runGrpcServer(srvName, addr, port string, errCh chan error) {
	srv := "grpc_" + srvName
	runFunc.runServer(srv, addr, port, errCh)
}

func (runFunc runFuncMap) runServer(srv, addr, port string, errCh chan error) {
	if _, ok := runFunc[srv]; !ok {
		errCh <- fmt.Errorf("no exist srv: %s", srv)
	}
	runFunc[srv](srv, addr, port, errCh)
}

func init() {
	Funcs = runFuncMap{}
	Funcs["grpc_"+account.ServiceName] = account.RunGrpc
	Funcs["grpc_"+fileconsole.ServiceName] = fileconsole.RunGrpc
}
