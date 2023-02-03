package redis

import (
	"context"
	"fmt"
	"github.com/1037group/dousheng/pkg/consts"
	"github.com/bsm/redislock"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var locker *redislock.Client

func Init() {

	rdb = redis.NewClient(&redis.Options{Addr: consts.RedisIp + ":" + consts.RedisPort, Password: ""})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis connect error.")
		panic(err)
	}
	locker = redislock.New(rdb)

	klog.CtxInfof(context.Background(), "Redis initialized successfully.")
}
