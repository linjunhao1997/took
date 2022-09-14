package account

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"
	"took/pkg/account/user"
)

type loadUserRequest struct {
	Id int
}

type loadUserResponse struct {
	User *user.User `json:"file,omitempty"`
	Err  error      `json:"error,omitempty"`
}

func makeLoadUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logrus.Info("收到请求")
		req := request.(loadUserRequest)
		user, err := s.LoadUser(ctx, req.Id)
		if err != nil {
			return nil, err
		}
		return loadUserResponse{User: user, Err: err}, nil
	}
}
