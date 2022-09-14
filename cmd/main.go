package main

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"took/pkg/account"
	"took/pkg/fileconsole"
)

func main() {
	errc := make(chan error)

	if len(os.Args) > 2 {
		serviceName := os.Args[1]
		if serviceName == "fileconsole" {
			go fileconsole.RunGrpc("fileconsole", "127.0.0.1", os.Args[2], errc)
		} else if serviceName == "account" {
			go account.RunGrpc("account", "127.0.0.1", os.Args[2], errc)
		} else {
			logrus.WithField("error", errors.New("service not exist")).Info("exit")
		}
	} else {
		logrus.WithField("error", errors.New("please input service name")).Info("exit")
	}

	logrus.WithField("error", <-errc).Info("exit")
}
