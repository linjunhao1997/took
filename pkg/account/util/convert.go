package util

import (
	"took/pkg/account/api/v1/grpc/proto"
	"took/pkg/account/domain/user"
)

func ConvertFormUserProto(proto *account.User) *user.User {
	result := &user.User{
		Id:       int(proto.Id),
		Username: proto.Username,
		Phone:    proto.Phone,
		Email:    proto.Email,
	}
	return result
}
func ConvertFromUser(user *user.User) *account.User {
	return &account.User{
		Id:        int32(user.Id),
		Username:  user.Username,
		Password:  user.Password,
		Phone:     user.Phone,
		Email:     user.Email,
		Disabled:  user.Disabled >= 0,
		CreatedAt: "",
	}
}
