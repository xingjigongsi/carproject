package redis

import (
	"carproject/framework/container"
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisApply struct {
	container.InterfaceContainer
	Addr        string
	Password    string
	DB          int
	MaxIdle     int
	IdleTimeout time.Duration
}

func GetConnetct(params ...interface{}) (interface{}, error) {
	containe := params[0].(container.InterfaceContainer)
	addr := params[1].(string)
	password := params[2].(string)
	db := params[3].(int)
	MaxIdle := params[4].(int)
	timeouitem := params[5].(int)
	timeout := time.Duration(timeouitem) * time.Second
	return &RedisApply{containe, addr, password, db, MaxIdle, timeout}, nil
}

func (redisApply *RedisApply) RedisPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisApply.Addr)
		},
		MaxIdle:     redisApply.MaxIdle,
		IdleTimeout: redisApply.IdleTimeout,
	}
}
