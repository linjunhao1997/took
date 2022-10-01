package grpc

import (
	"context"
	"fmt"
	account "took/pkg/account/api/v1/grpc/proto"
	"took/pkg/account/endpoint"
)

func decodeLoadUserRequestByGrpc(_ context.Context, req interface{}) (interface{}, error) {
	accountpb, ok := req.(*account.LoadUserRequest)
	if !ok {
		return nil, fmt.Errorf("grpc server decode request error")
	}
	request := endpoint.LoadUserRequest{
		Id: int(accountpb.Id),
	}
	return request, nil
}

func encodeLoadUserResponseByGrpc(_ context.Context, response interface{}) (interface{}, error) {
	data, ok := response.(endpoint.LoadUserResponse)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response error (%T)", data)
	}

	resp := &account.UserResponse{
		Data: &account.User{
			Id:       int32(data.User.Id),
			Username: data.User.Username,
			Phone:    data.User.Phone,
			Email:    data.User.Email,
			Disabled: int(data.User.Disabled) == 0,
			//CreatedAt: data.CreatedAt,
		},
	}

	return resp, nil
}
