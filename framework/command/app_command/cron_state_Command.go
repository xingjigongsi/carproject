package app_command

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"syscall"

	"github.com/xingjigongsi/carproject/framework/cobra"
	"github.com/xingjigongsi/carproject/framework/container"
)

func CronStateCommand() *cobra.Command {
	var cron = &cobra.Command{
		Use:   "state",
		Short: "cron 状态",
		RunE: func(cmd *cobra.Command, args []string) error {
			containe := cmd.Get()
			app := containe.MustMake(container.APPKEY).(*container.AppApply)
			workdir := app.BaseFolder()
			folderPid := path.Join(path.Join(workdir, "pid"), "cron.pid")
			pid, err := os.ReadFile(folderPid)
			if err != nil {
				return err
			}
			if len(pid) <= 0 {
				fmt.Println("cron 已经下线")
				return nil
			}
			atoi, err := strconv.Atoi(string(pid))
			if err != nil {
				return err
			}
			process, err := os.FindProcess(atoi)
			if err != nil {
				return err
			}
			if process == nil {
				fmt.Println("cron 已经下线")
				return nil
			}
			err = process.Signal(syscall.Signal(0))
			if err != nil {
				fmt.Println("cron 已经下线")

			}
			fmt.Println("cron 在线上")
			return nil
		},
	}
	return cron
}
