package app_command

import (
	"github.com/xingjigongsi/carproject/framework/cobra"
)

var Demon bool = true

func CronCommand() *cobra.Command {
	startcommand := CronStartCommand()
	startcommand.Flags().BoolVarP(&Demon, "demon", "d", false, "start cron demon")
	cron.AddCommand(startcommand)
	cron.AddCommand(CronStopCommand())
	cron.AddCommand(CronStateCommand())
	return cron
}

var cron = &cobra.Command{
	Use: "cron",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
