package redis

import (
	"context"
	"fmt"
	"github.com/1037group/dousheng/pkg/consts"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init() {

	rdb = redis.NewClient(&redis.Options{Addr: consts.RedisIp + ":" + consts.RedisPort, Password: ""})
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("redis connect error.")
		panic(err)
	}
	hlog.CtxInfof(context.Background(), "Redis initialized successfully.")
}
