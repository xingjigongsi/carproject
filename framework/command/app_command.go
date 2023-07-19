package command

import (
	"github.com/xingjigongsi/carproject/framework/cobra"
	"github.com/xingjigongsi/carproject/framework/command/app_command"
)

func AppCommand(command *cobra.Command) {
	command.AddCommand(app_command.ConfigCommand())
	command.AddCommand(app_command.CronCommand())

}
