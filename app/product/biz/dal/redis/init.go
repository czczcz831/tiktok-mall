package redis

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/product/conf"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	Rs          *redsync.Redsync
)

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	pool := goredis.NewPool(RedisClient) // or, pool := redigo.NewPool(...)
	Rs = redsync.New(pool)
}
