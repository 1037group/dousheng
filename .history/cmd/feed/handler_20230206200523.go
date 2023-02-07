package main

import (
	"context"
	"time"

	"github.com/1037group/dousheng/dal/db"
	douyin_feed "github.com/1037group/dousheng/kitex_gen/douyin_feed"
	"github.com/1037group/dousheng/kitex_gen/douyin_user"
	"github.com/1037group/dousheng/pack"
	"github.com/1037group/dousheng/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *douyin_feed.FeedRequest) (resp *douyin_feed.FeedResponse, err error) {

	klog.CtxInfof(ctx, "[FeedRequest] %+v", req)

	curTime := time.Now().UnixMilli()
	if req.LatestTime != nil {
		curTime = *req.LatestTime
	}
	res, err := db.MGetVideosByLastTime(ctx, db.DB, &curTime)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}
	userMap := make(map[int64]douyin_user.User)
	var userIDs []int64
	for _, m := range res {
		userIDs = append(userIDs, m.UserId)
		res, err := db.MIsFavoriteByUserId(ctx, db.DB, &req.ReqUserId, &m.VideoId)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return nil, errno.ParamErr
		}
	}
	res, err := db.MIsFavoriteByUserId(ctx, db.DB, &req.ReqUserId, &req.)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, errno.ParamErr
	}


	users, err := db.MGetUserByID(ctx, db.DB, userIDs)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}

	for _, m := range users {
		var isFollow bool

		// 设置了token时检验是否follow
		if req.ReqUserId != nil {
			isFollow, err = db.CheckFollow(ctx, db.DB, *req.ReqUserId, m.UserId)
			if err != nil {
				klog.CtxErrorf(ctx, err.Error())
				return nil, err
			}
		}

		userMap[m.UserId] = douyin_user.User{
			Id:            m.UserId,
			Name:          m.UserName,
			FollowCount:   &m.UserFollowCount,
			FollowerCount: &m.UserFollowerCount,
			IsFollow:      isFollow,
		}
	}

	videoList := pack.Videos(res, userMap)

	return &douyin_feed.FeedResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		VideoList:  videoList,
		NextTime:   &curTime,
	}, nil
}
