package main

import (
	"context"
	"github.com/1037group/dousheng/dal/db"
	douyin_feed "github.com/1037group/dousheng/kitex_gen/douyin_feed"
	"github.com/1037group/dousheng/pack"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *douyin_feed.FeedRequest) (resp *douyin_feed.FeedResponse, err error) {

	klog.CtxInfof(ctx, "FeedRequest %+v", req)

	curTime := time.Now().UnixMilli()
	if req.LatestTime != nil {
		curTime = *req.LatestTime
	}
	res, err := db.MGetVideos(ctx, &curTime)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}
	klog.CtxInfof(ctx, "res: %+v", res)

	userMap := make(map[int64]sql.User)
	var userIDs []int64
	for _, m := range res {
		userIDs = append(userIDs, m.UserId)
	}

	users, err := db.MGetUserByID(ctx, userIDs)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}

	for _, m := range users {
		userMap[m.UserId] = *m
	}

	videoList := pack.Videos(res, userMap)

	return &douyin_feed.FeedResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		VideoList:  videoList,
		NextTime:   &curTime,
	}, nil
}
