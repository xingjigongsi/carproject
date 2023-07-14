package redis

import "github.com/garyburd/redigo/redis"

const REDIS_NAME = "app:redis"

type RedisInterface interface {
	RedisPool() *redis.Pool
}
