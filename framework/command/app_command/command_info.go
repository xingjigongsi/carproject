package app_command

import (
	"fmt"
	"github.com/xingjigongsi/carproject/framework/cobra"
)

func CommandInformation() *cobra.Command {
	Command.AddCommand(CommandList())
	return Command
}

func CommandList() *cobra.Command {
	list := &cobra.Command{
		Use:   "list",
		Short: "系统相关的command",
		RunE: func(cmd *cobra.Command, args []string) error {
			root := cmd.Root().Commands()
			for i, v := range root {
				fmt.Printf("%d\t%s\t%s\n", i+1, v.Name(), v.Short)
			}
			return nil
		},
	}
	return list
}

var Command = &cobra.Command{
	Use:   "command",
	Short: "command 相关的命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			cmd.Help()
		}
		return nil
	},
}
