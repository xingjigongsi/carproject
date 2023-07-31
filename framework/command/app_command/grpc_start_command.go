package app_command

import (
	"github.com/xingjigongsi/carproject/framework/cobra"
	"github.com/xingjigongsi/carproject/framework/components/netServer"
	"net"
)

func GrpcStart() *cobra.Command {
	grpcStart := &cobra.Command{
		Use:   "start",
		Short: "grpc 启动",
		RunE: func(cmd *cobra.Command, args []string) error {
			container := cmd.Get()
			netserver := container.MustMake(netServer.NET_NAME).(*netServer.NetApply)
			server := netserver.GrpcServer
			listen, err := net.Listen("tcp", ":"+Port)
			if err != nil {
				return err
			}
			if err := server.Serve(listen); err != nil {
				return err
			}
			return nil
		},
	}
	grpcStart.Flags().BoolVarP(&Demon, "demon", "d", false, "grpc后台执行")
	grpcStart.Flags().StringVarP(&Port, "port", "p", "8081", "grpc后台端口号")
	return grpcStart
}
