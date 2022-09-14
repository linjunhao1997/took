package util

import (
	"took/pkg/account/accountpb"
	"took/pkg/account/user"
)

func ConvertFormUserProto(proto *accountpb.User) *user.User {
	user := &user.User{
		Id:       int(proto.Id),
		Username: proto.Username,
		Phone:    proto.Phone,
		Email:    proto.Email,
	}
	return user
}
func ConvertFromUser(user *user.User) *accountpb.User {
	return &accountpb.User{
		Id:        int32(user.Id),
		Username:  user.Username,
		Password:  user.Password,
		Phone:     user.Phone,
		Email:     user.Email,
		Disabled:  user.Disabled >= 0,
		CreatedAt: "",
	}
}
