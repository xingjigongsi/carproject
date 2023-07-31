package app_command

import (
	"github.com/xingjigongsi/carproject/framework/cobra"
)

var grpc = &cobra.Command{
	Use:   "grpc",
	Short: "grpc 服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Help()
		}
		return cmd.Execute()
	},
}

func GrpcCommand() *cobra.Command {
	grpc.AddCommand(GrpcStart())
	return grpc
}
