package netServer

import "google.golang.org/grpc"

type NetApply struct {
	GrpcServer *grpc.Server
}

func BindServer(params ...interface{}) (interface{}, error) {
	grpcserver := params[0].(*grpc.Server)
	return &NetApply{grpcserver}, nil
}
