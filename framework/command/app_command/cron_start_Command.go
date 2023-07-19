package app_command

import (
	"fmt"
	"github.com/erikdubbelboer/gspt"
	daemon "github.com/sevlyar/go-daemon"
	"github.com/xingjigongsi/carproject/framework/cobra"
	"github.com/xingjigongsi/carproject/framework/container"
	"github.com/xingjigongsi/carproject/framework/util"
	"log"
	"os"
	"path"
	"strconv"
)

func CronStartCommand() *cobra.Command {
	var cron = &cobra.Command{
		Use:   "start",
		Short: "cron 启动",
		RunE: func(cmd *cobra.Command, args []string) error {
			containe := cmd.Get()
			app := containe.MustMake(container.APPKEY).(*container.AppApply)
			workdir := app.BaseFolder()
			folderPid := path.Join(workdir, "pid")
			foldrlog := path.Join(workdir, "log")
			if !util.PathIsExist(folderPid) {
				err := os.Mkdir(folderPid, os.ModePerm)
				if err != nil {
					panic("进程id文件夹创建失败")
				}
			}
			if !util.PathIsExist(foldrlog) {
				err := os.Mkdir(foldrlog, os.ModePerm)
				if err != nil {
					panic("日志文件夹创建失败")
				}
			}
			cronpid := path.Join(folderPid, "cron.pid")
			cronlog := path.Join(foldrlog, "cron.log")
			pid := strconv.Itoa(os.Getpid())
			if Demon {
				context := &daemon.Context{
					PidFileName: cronpid,
					PidFilePerm: 0644,
					LogFileName: cronlog,
					LogFilePerm: 0640,
					WorkDir:     workdir,
					Args:        []string{"", "cron", "start", "-d=true"},
					Umask:       027,
				}
				child, err := context.Reborn()
				if err != nil {
					return err
				}
				if child != nil {
					fmt.Println("父进程ID", child.Pid)
					return nil
				}
				defer func(context *daemon.Context) {
					err := context.Release()
					if err != nil {
						log.Printf("释放失败:%s", err.Error())
					}
					log.Printf("释放成功!!!")
				}(context)
				gspt.SetProcTitle("cron")
				cmd.Root().Cron.Run()
				return nil
			}
			os.WriteFile(cronpid, []byte(pid), os.ModePerm)
			gspt.SetProcTitle("cron")
			cmd.Root().Cron.Run()
			return nil
		},
	}
	return cron
}
