package cron_command

import (
	"errors"
	"github.com/robfig/cron/v3"
	"github.com/xingjigongsi/carproject/framework/cobra"
	"github.com/xingjigongsi/carproject/framework/components/redis"
	"log"
	"time"
)

func CronSelectCommand(cmd *cobra.Command, serverName string, spec string, fun func(), t time.Duration) error {
	root := cmd.Root()
	if root.Cron == nil {
		root.Cron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
		root.CronList = []cobra.CronList{}
		root.CronListMap = make(map[string]struct{})
	}
	if _, ok := root.CronListMap[serverName]; ok {
		return errors.New("cron 的 " + serverName + "重复")
	}
	root.CronList = append(root.CronList, cobra.CronList{Command: cmd, ServerName: serverName, Spec: spec})
	root.CronListMap[serverName] = struct{}{}
	containe := root.Containe
	var command cobra.Command
	command = *cmd
	command.SetParent()
	command.SetArgEmpty()
	root.Cron.AddFunc(spec, func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		redisApply := containe.MustMake(redis.REDIS_NAME).(*redis.RedisApply)
		redis := redisApply.RedisPool()
		reply, err := redis.Get().Do("set", serverName, "1", "EX", t.Seconds(), "NX")
		if err != nil {
			return
		}
		if s, ok := reply.(string); ok && s == "OK" {
			fun()
			command.ExecuteContext(root.Context())
		}
	})
	return nil
}
