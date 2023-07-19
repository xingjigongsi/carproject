package app_command

import (
	"fmt"
	"github.com/xingjigongsi/carproject/framework/cobra"
	"github.com/xingjigongsi/carproject/framework/container"
	"os"
	"path"
	"strconv"
	"syscall"
)

func CronStopCommand() *cobra.Command {
	var cron = &cobra.Command{
		Use:   "stop",
		Short: "cron 关闭",
		RunE: func(cmd *cobra.Command, args []string) error {
			containe := cmd.Get()
			app := containe.MustMake(container.APPKEY).(*container.AppApply)
			workdir := app.BaseFolder()
			folderPid := path.Join(path.Join(workdir, "pid"), "cron.pid")
			pid, err := os.ReadFile(folderPid)
			if err != nil {
				return err
			}
			atoi, err := strconv.Atoi(string(pid))
			if err != nil {
				return err
			}
			err = syscall.Kill(atoi, syscall.SIGTERM|syscall.SIGQUIT)
			if err != nil {
				return err
			}
			err = os.WriteFile(folderPid, []byte{}, os.ModePerm)
			if err != nil {
				return err
			}
			fmt.Println("stop pid:", atoi)
			//cmd.Root().Cron.Stop()
			return nil
		},
	}
	return cron
}
