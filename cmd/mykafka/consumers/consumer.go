package main

import (
	"context"
	"encoding/json"
	"github.com/1037group/dousheng/cmd/api/biz/rpc"
	"github.com/1037group/dousheng/dal/db"
	"github.com/1037group/dousheng/dal/redis"
	"github.com/1037group/dousheng/kitex_gen/douyin_favorite"
	"github.com/1037group/dousheng/pkg/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"strconv"
	"time"
)

func GetFavoriteActionConsumer(ctx context.Context) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": consts.KafkaHost,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{consts.TopicFavoriteAction}, nil)

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			klog.CtxInfof(ctx, "Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			var req douyin_favorite.FavoriteActionRequest
			err := json.Unmarshal(msg.Value, &req)
			if err != nil {
				klog.CtxErrorf(ctx, err.Error())
				continue
			}

			// favorite表中的变动直接由rpc处理
			_, err = rpc.FavoriteAction(ctx, &req)
			if err != nil {
				klog.CtxErrorf(ctx, err.Error())
				continue
			}

			// video表中的计数由redis暂存，然后由cronjob异步更新
			var operatorType int
			if req.ActionType == 1 {
				operatorType = redis.Add
			} else {
				operatorType = redis.Sub
			}

			// 暂时只用0分片
			hashKey := redis.GetHashKeyName(redis.ModelNameVideo, 0)
			hashFieldName := redis.GetHashFieldName(req.VideoId, redis.VideoFavoriteCount)

			_, err = redis.Get(ctx, hashKey, hashFieldName)
			// redis中不存在时从mysql中获取
			if err != nil {
				favoriteCount, err := db.GetFavoriteCount(ctx, db.DB, req.VideoId)
				if err != nil {
					klog.CtxErrorf(ctx, err.Error())
					continue
				}
				klog.CtxInfof(ctx, "[GetFavoriteActionConsumer] favoriteCount: %+v", favoriteCount)
				redis.Set(ctx, hashKey, hashFieldName, strconv.FormatInt(favoriteCount+1, 10))
			} else {
				// 存在则更新redis的值
				_, err = redis.UpdateCachedCount(ctx, redis.ModelNameFavorite, redis.VideoFavoriteCount, req.VideoId, operatorType)
				if err != nil {
					klog.CtxErrorf(ctx, err.Error())
					continue
				}
			}
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			klog.CtxErrorf(ctx, "Consumer error: %v (%v)\n", err, msg)
		}
	}
}
