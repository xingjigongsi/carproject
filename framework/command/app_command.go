package command

import (
	"github.com/xingjigongsi/carproject/framework/cobra"
	"github.com/xingjigongsi/carproject/framework/command/app_command"
	"github.com/xingjigongsi/carproject/framework/command/back_Command"
)

func AppCommand(command *cobra.Command) {
	command.AddCommand(app_command.ConfigCommand())
	command.AddCommand(app_command.CronCommand())
	command.AddCommand(app_command.BuildCommand())
	command.AddCommand(app_command.CommandInformation())
	command.AddCommand(app_command.SystemCommand())
	command.AddCommand(back_Command.CommandRestart())
	command.AddCommand(app_command.GrpcCommand())
}
