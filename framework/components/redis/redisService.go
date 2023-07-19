package redis

import (
	"github.com/xingjigongsi/carproject/framework/components/parse"
	"github.com/xingjigongsi/carproject/framework/container"
)

type RedisService struct {
}

func (redisService *RedisService) Register(contain container.InterfaceContainer) container.NewInstance {
	return GetConnetct
}

func (redisService *RedisService) Params(contain container.InterfaceContainer) []interface{} {
	mustMake := contain.MustMake(parse.PASE_NAME).(*parse.ParseApply)
	redisAddr := ""
	if addr, err := mustMake.GetString("redis.Addr"); err == nil {
		redisAddr = addr
	}
	passWord := ""
	if pwd, err := mustMake.GetString("redis.Password"); err == nil {
		passWord = pwd
	}
	selectDB := 0
	if db, err := mustMake.GetInt("redis.DB"); err == nil {
		selectDB = db
	}
	MaxIdle := 3
	if val, err := mustMake.GetInt("redis.MaxIdle"); err == nil {
		MaxIdle = val
	}
	timeout := 0
	if timeoutval, err := mustMake.GetInt("redis.Timeout"); err == nil {
		timeout = timeoutval
	}
	return []interface{}{contain, redisAddr, passWord, selectDB, MaxIdle, timeout}
}
func (redisService *RedisService) ApplyInit(contain container.InterfaceContainer) error {
	return nil
}
func (redisService *RedisService) Name() string {
	return REDIS_NAME
}
func (redisService *RedisService) IsDefer() bool {
	return false
}
