package main

import (
	"github.com/xingjigongsi/carproject/framework/command"
	"github.com/xingjigongsi/carproject/framework/components/log"
	"github.com/xingjigongsi/carproject/framework/components/mongodb"
	"github.com/xingjigongsi/carproject/framework/components/parse"
	"github.com/xingjigongsi/carproject/framework/components/redis"
	"github.com/xingjigongsi/carproject/framework/container"
)

func main() {
	newContainer := container.NewContainer()
	newContainer.Bind(&container.AppApplyService{})
	newContainer.Bind(&parse.ParseService{})
	newContainer.Bind(&mongodb.MongoDbService{})
	newContainer.Bind(&log.LogService{})
	newContainer.Bind(&redis.RedisService{})
	command.RunCommand(newContainer)
}
