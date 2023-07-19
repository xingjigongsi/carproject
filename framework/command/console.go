package command

import (
	"github.com/xingjigongsi/carproject/framework/cobra"
	"github.com/xingjigongsi/carproject/framework/container"
)

func RunCommand(container *container.Container) error {

	var rootCmd = &cobra.Command{
		Use: "app",
		RunE: func(cmd *cobra.Command, args []string) error {
			//cmd.InitDefaultHelpFlag()
			return nil
		},
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
	rootCmd.Set(container)
	AppCommand(rootCmd)
	CronCommand(rootCmd)
	return rootCmd.Execute()
}
