package netServer

import (
	"github.com/xingjigongsi/carproject/api/protobuf/user/v1/proto"
	"github.com/xingjigongsi/carproject/framework/container"
	"github.com/xingjigongsi/carproject/internal/grpc/server/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer(container container.InterfaceContainer) (*grpc.Server, error) {
	server := grpc.NewServer()
	proto.RegisterUserServiceServer(server, &user.UserRegister{Container: container})
	reflection.Register(server)
	return server, nil
}
