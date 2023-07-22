package test

import (
	"context"
	"fmt"
	proto2 "github.com/xingjigongsi/carproject/api/protobuf/user/v1/proto"
	"github.com/xingjigongsi/carproject/framework/components/redis"
	"google.golang.org/grpc"
	"os"
	"testing"

	"github.com/xingjigongsi/carproject/framework/components/log"
	"github.com/xingjigongsi/carproject/framework/components/mongodb"
	"github.com/xingjigongsi/carproject/framework/components/parse"
	"github.com/xingjigongsi/carproject/framework/container"
)

func TestParseYml(t *testing.T) {
	//pkg.NewParseFile("../configs", ".").GetString("database.mongodb.mongodbUrl")
	//fmt.Println(getString["path"])
	newContainer := container.NewContainer()
	newContainer.Bind(&container.AppApplyService{})
	//mustMake := newContainer.MustMake(container.APPKEY)
	//service := mustMake.(container.ApplyService)
	newContainer.Bind(&parse.ParseService{})
	//apply := newContainer.MustMake(moudle.PASE_NAME).(*moudle.ParseApply)
	newContainer.Bind(&mongodb.MongoDbService{})
	////fmt.Println(apply.GetString("database.mongodb.mongodbUrl"))
	//apply := newContainer.MustMake(mongodb.MONDBAPP).(*mongodb.MongodbApply)
	//_, err := apply.MongodbClient()
	//fmt.Println(err)

	newContainer.Bind(&log.LogService{})
	newContainer.Bind(&redis.RedisService{})

	apply := newContainer.MustMake(redis.REDIS_NAME).(*redis.RedisApply)
	pool := apply.RedisPool()
	fmt.Println(pool.Get().Do("Get", "a"))

	//fmt.Println(logs)
	//fmt.Println("sfddfdfsfssf")
	//err := logs.Logf(context.Background(), "测试数据", map[string]interface{}{})
	//fmt.Println(err)
}

func TestGrpc(t *testing.T) {
	dial, err := grpc.Dial(":8099", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer dial.Close()
	serviceClient := proto2.NewUserServiceClient(dial)
	serviceClient.RegisterUser(context.Background(), &proto2.UserMessage{})

}
