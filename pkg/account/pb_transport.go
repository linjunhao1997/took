package account

import (
	"context"
	"fmt"
	"took/pkg/account/accountpb"
)

func decodeLoadUserRequestByGrpc(_ context.Context, req interface{}) (interface{}, error) {
	accountpb, ok := req.(*accountpb.LoadUserRequest)
	if !ok {
		return nil, fmt.Errorf("grpc server decode request error")
	}
	request := loadUserRequest{
		Id: int(accountpb.Id),
	}
	return request, nil
}

func encodeLoadUserResponseByGrpc(_ context.Context, response interface{}) (interface{}, error) {
	data, ok := response.(loadUserResponse)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response error (%T)", data)
	}

	resp := &accountpb.UserResponse{
		Data: &accountpb.User{
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
