package user

import (
	"context"
	"fmt"
	proto2 "github.com/xingjigongsi/carproject/api/protobuf/user/v1/proto"
	"github.com/xingjigongsi/carproject/framework/components/parse"
	"github.com/xingjigongsi/carproject/framework/container"
)

type UserRegister struct {
	Container container.InterfaceContainer
}

func (c *UserRegister) RegisterUser(ctx context.Context, in *proto2.UserMessage) (*proto2.RegiterRegisterUser, error) {
	apply := c.Container.MustMake(parse.PASE_NAME).(*parse.ParseApply)
	getString, _ := apply.GetString("database.mongodb.mongodbUrl")
	fmt.Println(getString)
	fmt.Println("fdsdsfdsfdsfdsfdsf")
	return &proto2.RegiterRegisterUser{}, nil
}
