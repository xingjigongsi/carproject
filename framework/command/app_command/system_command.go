package app_command

import (
	"fmt"
	"github.com/xingjigongsi/carproject/api/protobuf/user/v1/proto"
	"github.com/xingjigongsi/carproject/internal/grpc/server/user"
	"log"
	"net"
	"os"
	"path"
	"strconv"
	"syscall"
	"time"

	"github.com/erikdubbelboer/gspt"
	"github.com/sevlyar/go-daemon"
	"google.golang.org/grpc"

	"github.com/xingjigongsi/carproject/framework/cobra"
	"github.com/xingjigongsi/carproject/framework/container"
	"github.com/xingjigongsi/carproject/framework/util"
)

var Port string

func SystemCommand() *cobra.Command {
	var system = &cobra.Command{
		Use:   "system",
		Short: "系统相关命令",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	system.AddCommand(startSystemCommand())
	system.AddCommand(stopSystemCommand())
	return system
}

func startSystemCommand() *cobra.Command {
	var startSystem = &cobra.Command{
		Use:   "start",
		Short: "启动系统",
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
			systempid := path.Join(folderPid, "system.pid")
			systemlog := path.Join(foldrlog, "system.log")
			pid := strconv.Itoa(os.Getpid())
			if Demon {
				context := &daemon.Context{
					PidFileName: systempid,
					PidFilePerm: 0644,
					LogFileName: systemlog,
					LogFilePerm: 0640,
					WorkDir:     workdir,
					Args:        []string{"", "system", "start", "-d=true"},
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
				gspt.SetProcTitle("system")
				StartgRpc()
				return nil
			}
			os.WriteFile(systempid, []byte(pid), os.ModePerm)
			gspt.SetProcTitle("system")
			StartgRpc()
			return nil
		},
	}
	startSystem.Flags().BoolVarP(&Demon, "demon", "d", false, "系统后台执行")
	startSystem.Flags().StringVarP(&Port, "port", "p", "8081", "系统端口号")
	return startSystem
}

func stopSystemCommand() *cobra.Command {
	var stoptSystem = &cobra.Command{
		Use:   "stop",
		Short: "关闭系统",
		RunE: func(cmd *cobra.Command, args []string) error {
			containe := cmd.Get()
			app := containe.MustMake(container.APPKEY).(*container.AppApply)
			workdir := app.BaseFolder()
			folderPid := path.Join(path.Join(workdir, "pid"), "system.pid")
			pid, err := os.ReadFile(folderPid)
			if err != nil || len(pid) < 0 {
				return err
			}
			atoi, err := strconv.Atoi(string(pid))
			if err != nil {
				return err
			}
			for i := 0; i < 10; i++ {
				process, err := os.FindProcess(atoi)
				if err != nil {
					break
				}
				err = process.Signal(syscall.Signal(0))
				if err != nil {
					break
				}
				err = syscall.Kill(atoi, syscall.SIGTERM|syscall.SIGQUIT)
				if err != nil {
					return err
				}
				time.Sleep(1 * time.Second)
			}
			err = os.WriteFile(folderPid, []byte{}, os.ModePerm)
			if err != nil {
				return err
			}
			fmt.Println("stop pid:", atoi)
			return nil
		},
	}
	return stoptSystem
}

func StartgRpc() {
	listen, err := net.Listen("tcp", ":8099")
	if err != nil {
		log.Fatalf("%v", err)
	}
	g := grpc.NewServer()
	proto.RegisterUserServiceServer(g, &user.UserRegister{})

	g.Serve(listen)
}
