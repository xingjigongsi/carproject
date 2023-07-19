package app_command

import (
	"fmt"
	"github.com/xingjigongsi/carproject/framework/cobra"
	"github.com/xingjigongsi/carproject/framework/components/parse"
	"os"
)

func ConfigCommand() *cobra.Command {
	var config = &cobra.Command{
		Use:   "config",
		Short: "./main config",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				fmt.Println("参数不能为空")
				os.Exit(-1)
				cmd.Help()
			}
			contain := cmd.Get()
			parseApply := contain.MustMake(parse.PASE_NAME).(*parse.ParseApply)
			parse, err := parseApply.GetString(args[0])
			if err != nil {
				return err
			}
			fmt.Println(parse)
			return nil
		},
	}
	return config
}
