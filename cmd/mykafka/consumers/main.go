package main

import (
	"context"

	"github.com/1037group/dousheng/cmd/api/biz/rpc"
	"github.com/1037group/dousheng/dal/db"
	"github.com/1037group/dousheng/dal/redis"
)

func Init() {
	db.Init()
	redis.Init()
	rpc.Init()
}

func main() {
	Init()
	ctx := context.Background()

	go GetFavoriteActionConsumer(ctx)
	go GetCommentActionConsumer(ctx)
	go GetRelationActionConsumer(ctx)

	ExecuteVideoCron(ctx)

	run := true

	for run {

	}
}
