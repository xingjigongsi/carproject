package command

import (
	"fmt"
	"github.com/xingjigongsi/carproject/framework/cobra"
	"github.com/xingjigongsi/carproject/framework/command/cron_command"
	"time"
)

var fun = func() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
var fun1 = func() {
	fmt.Println("fsdfsddfsdfsdfsfsdfdsfsd")
}
var fun2 = func() {
	fmt.Println("liuxingyu")
}

func CronCommand(cmd *cobra.Command) {
	cron_command.CronSelectCommand(cmd, "time_print", "0/7 * * * * *", fun, 1*time.Second)
	cron_command.CronSelectCommand(cmd, "time_print_1", "0/12 * * * * *", fun1, 1*time.Second)
}
