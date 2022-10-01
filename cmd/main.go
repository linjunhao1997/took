package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"took/pkg/util/run"
)

var (
	name = flag.String("name", "", "input service name")
	addr = flag.String("addr", "127.0.0.1", "input addr")
	port = flag.String("port", "9090", "input port")
	t    = flag.String("type", "", "input service type, eg:http,grpc")
)

func main() {

	errCh := make(chan error)
	flag.Parse()

	go run.Funcs.RunServer(*name, *addr, *port, *t, errCh)
	logrus.WithField("error", <-errCh).Info("exit")
}
