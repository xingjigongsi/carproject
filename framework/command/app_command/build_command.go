package app_command

import (
	"fmt"
	"github.com/xingjigongsi/carproject/framework/cobra"
	"os/exec"
)

func BuildCommand() *cobra.Command {
	var buildCommand = &cobra.Command{
		Use:   "build",
		Short: "服务的编译命令",
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := exec.LookPath("go")
			if err != nil {
				return err
			}
			command := exec.Command(path, "build", "-o", "main", "./")
			output, err := command.CombinedOutput()
			if err != nil {
				fmt.Println(output)
				return err
			}
			return nil
		},
	}
	return buildCommand
}
