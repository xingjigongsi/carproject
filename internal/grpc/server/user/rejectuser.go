package user

import (
	proto2 "carproject/api/protobuf/user/v1/proto"
	"context"
	"google.golang.org/grpc"
)

type UserRegister struct {
}

func (c *UserRegister) RegisterUser(ctx context.Context, in *proto2.UserMessage, opts ...grpc.CallOption) (*proto2.RegiterRegisterUser, error) {

}
