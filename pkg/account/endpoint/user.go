package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"took/pkg/account/domain/user"
	"took/pkg/account/service"
)

type LoadUserRequest struct {
	Id int
}

type LoadUserResponse struct {
	User *user.User `json:"file,omitempty"`
	Err  error      `json:"error,omitempty"`
}

func MakeLoadUserEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoadUserRequest)
		user, err := s.LoadUser(ctx, req.Id)
		if err != nil {
			return nil, err
		}
		return LoadUserResponse{User: user, Err: err}, nil
	}
}
