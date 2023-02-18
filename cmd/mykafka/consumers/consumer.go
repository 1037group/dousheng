package main

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/1037group/dousheng/cmd/api/biz/rpc"
	"github.com/1037group/dousheng/dal/db"
	"github.com/1037group/dousheng/dal/redis"
	"github.com/1037group/dousheng/kitex_gen/douyin_comment"
	"github.com/1037group/dousheng/kitex_gen/douyin_favorite"
	"github.com/1037group/dousheng/kitex_gen/douyin_relation"
	"github.com/1037group/dousheng/pkg/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// 将评论数量放在消息队列中处理,处理方式和点赞类似
func GetCommentActionConsumer(ctx context.Context) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": consts.KafkaHost,
		"group.id":          "myGroup",
		"auto.offset.reset": "latest",
	})

	if err != nil {
		panic(err)
	}
	c.SubscribeTopics([]string{consts.TopicCommentAction}, nil)
	run := true
	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			klog.CtxInfof(ctx, "Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			var req douyin_comment.CommentActionRequest
			err := json.Unmarshal(msg.Value, &req)
			if err != nil {
				klog.CtxErrorf(ctx, err.Error())
				continue
			}

			// comment表中的变动直接由rpc处理，这里的rpc放在了api service里面
			// _, err = rpc.CommentAction(ctx, &req)
			// if err != nil {
			// 	klog.CtxErrorf(ctx, err.Error())
			// 	continue
			// }

			// video表中VideoCommentCount的计数由redis暂存，然后由cronjob异步更新
			var operatorType int
			if req.ActionType == 1 {
				operatorType = redis.Add
			} else {
				operatorType = redis.Sub
			}

			// 暂时只用0分片
			hashKey := redis.GetHashKeyName(redis.ModelNameVideo, 0)
			hashFieldName := redis.GetHashFieldName(req.VideoId, redis.VideoCommentCount)

			_, err = redis.Get(ctx, hashKey, hashFieldName)
			// redis中不存在时从mysql中获取
			if err != nil {
				commentCount, err := db.GetCommentCount(ctx, db.DB, req.VideoId)
				if err != nil {
					klog.CtxErrorf(ctx, err.Error())
					continue
				}
				klog.CtxInfof(ctx, "[GetCommentActionConsumer] CommentCount: %+v", commentCount)
				if operatorType == redis.Add {
					redis.Set(ctx, hashKey, hashFieldName, strconv.FormatInt(commentCount+1, 10))
				} else {
					redis.Set(ctx, hashKey, hashFieldName, strconv.FormatInt(commentCount-1, 10))
				}
			} else {
				// 存在则更新redis的值
				_, err = redis.UpdateCachedCount(ctx, redis.ModelNameComment, redis.VideoCommentCount, req.VideoId, operatorType)
				if err != nil {
					klog.CtxErrorf(ctx, err.Error())
					continue
				}
			}
		} else if !err.(kafka.Error).IsTimeout() {
			klog.CtxErrorf(ctx, "Consumer error: %v (%v)\n", err, msg)
		}
	}
}

// 将关注放在消息队列中处理
func GetRelationActionConsumer(ctx context.Context) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": consts.KafkaHost,
		"group.id":          "myGroup",
		"auto.offset.reset": "latest",
	})
	if err != nil {
		panic(err)
	}
	c.SubscribeTopics([]string{consts.TopicRelationAction}, nil)
	run := true
	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			klog.CtxInfof(ctx, "Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			var req douyin_relation.RelationActionRequest
			err := json.Unmarshal(msg.Value, &req)
			if err != nil {
				klog.CtxErrorf(ctx, err.Error())
				continue
			}
			//relation表的变动直接由rpc处理
			// _, err = rpc.RelationAction(ctx, &req)
			// if err != nil {
			// 	klog.CtxErrorf(ctx, err.Error())
			// 	continue
			// }

			//user表中follow和follower的计数由redis缓存，并用cronjob异步更新
			var operatorType int
			if req.ActionType == 1 {
				operatorType = redis.Add
			} else {
				operatorType = redis.Sub
			}
			hashKey := redis.GetHashKeyName(redis.ModelNameUser, 0)

			//关注的发起者，他的followcount增加了，而被关注的人，followerCount增加了
			hashFieldName1 := redis.GetHashFieldName(req.ReqUserId, redis.UserFollowCount)
			klog.CtxInfof(ctx, "[GetHashFieldName] hashFieldName1: %+v", hashFieldName1)
			hashFieldName2 := redis.GetHashFieldName(req.ToUserId, redis.UserFollowerCount)
			klog.CtxInfof(ctx, "[GetHashFieldName] hashFieldName2: %+v", hashFieldName2)

			_, err = redis.Get(ctx, hashKey, hashFieldName1)
			if err != nil {
				followCount, err := db.GetFollowCount(ctx, db.DB, req.ReqUserId)
				if err != nil {
					klog.CtxErrorf(ctx, err.Error())
					continue
				}
				klog.CtxInfof(ctx, "[GetRelationActionConsumer] followCount: %+v", followCount)
				if operatorType == redis.Add {
					redis.Set(ctx, hashKey, hashFieldName1, strconv.FormatInt(followCount+1, 10))
				} else {
					redis.Set(ctx, hashKey, hashFieldName1, strconv.FormatInt(followCount-1, 10))
				}
			} else {
				// 存在则更新redis的值
				_, err = redis.UpdateCachedCount(ctx, redis.ModelNameUser, redis.UserFollowCount, req.ReqUserId, operatorType)
				if err != nil {
					klog.CtxErrorf(ctx, err.Error())
					continue
				}
			}
			_, err = redis.Get(ctx, hashKey, hashFieldName2)
			if err != nil {
				followerCount, err := db.GetFollowerCount(ctx, db.DB, req.ToUserId)
				if err != nil {
					klog.CtxErrorf(ctx, err.Error())
					continue
				}
				klog.CtxInfof(ctx, "[GetRelationActionConsumer] followerCount: %+v", followerCount)
				if operatorType == redis.Add {
					redis.Set(ctx, hashKey, hashFieldName2, strconv.FormatInt(followerCount+1, 10))
				} else {
					redis.Set(ctx, hashKey, hashFieldName2, strconv.FormatInt(followerCount-1, 10))
				}
			} else {
				// 存在则更新redis的值
				_, err = redis.UpdateCachedCount(ctx, redis.ModelNameUser, redis.UserFollowerCount, req.ToUserId, operatorType)
				if err != nil {
					klog.CtxErrorf(ctx, err.Error())
					continue
				}
			}

		} else if !err.(kafka.Error).IsTimeout() {
			klog.CtxErrorf(ctx, "Consumer error: %v (%v)\n", err, msg)
		}
	}

}

func GetFavoriteActionConsumer(ctx context.Context) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": consts.KafkaHost,
		"group.id":          "myGroup",
		"auto.offset.reset": "latest",
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
				if operatorType == redis.Add {
					redis.Set(ctx, hashKey, hashFieldName, strconv.FormatInt(favoriteCount+1, 10))
				} else {
					redis.Set(ctx, hashKey, hashFieldName, strconv.FormatInt(favoriteCount-1, 10))
				}
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
