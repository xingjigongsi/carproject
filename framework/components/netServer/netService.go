package netServer

import (
	"github.com/xingjigongsi/carproject/framework/container"
	"google.golang.org/grpc"
)

type NetService struct {
	Grpcserver *grpc.Server
}

func (netservice *NetService) Register(contain container.InterfaceContainer) container.NewInstance {
	return BindServer
}

func (netservice *NetService) Params(contain container.InterfaceContainer) []interface{} {
	result := []interface{}{}
	if netservice.Grpcserver == nil {
		result = append(result, grpc.NewServer())
	} else {
		result = append(result, netservice.Grpcserver)
	}
	return result
}

func (netservice *NetService) ApplyInit(contain container.InterfaceContainer) error {
	return nil
}

func (netservice *NetService) Name() string {
	return NET_NAME
}
func (netservice *NetService) IsDefer() bool {
	return false
}
