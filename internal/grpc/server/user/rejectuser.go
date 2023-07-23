package user

import (
	"context"
	"fmt"
	proto2 "github.com/xingjigongsi/carproject/api/protobuf/user/v1/proto"
)

type UserRegister struct {
}

func (c *UserRegister) RegisterUser(ctx context.Context, in *proto2.UserMessage) (*proto2.RegiterRegisterUser, error) {
	fmt.Println("omyererfsfs")
	fmt.Println("fdsfdfsdfdsfsfd")
	return &proto2.RegiterRegisterUser{}, nil
}
