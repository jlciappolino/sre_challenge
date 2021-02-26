package infra

import "github.com/go-redis/redis/v8"

type RedisConn struct{
	*redis.Client
}

func NewRedisConn() RedisConn {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return RedisConn{rdb}
}


