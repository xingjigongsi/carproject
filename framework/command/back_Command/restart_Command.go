package back_Command

import (
	"github.com/xingjigongsi/carproject/framework/cobra"
	"log"
)

var Port string

func CommandRestart() *cobra.Command {
	restart := &cobra.Command{
		Use:   "BackEnd",
		Short: "后端启动",
		RunE: func(cmd *cobra.Command, args []string) error {
			backend := NewBackend(cmd.Get())
			go backend.MoniterFolder()
			err := backend.StartBackend()
			if err != nil {
				log.Printf("%v", err)
			}
			select {}
			return nil
		},
	}
	restart.Flags().StringVarP(&Port, "port", "p", "8080", "服务启动的端口")
	return restart
}
