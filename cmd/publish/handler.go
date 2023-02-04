package main

import (
	"context"
	"github.com/1037group/dousheng/dal/db"
	douyin_publish "github.com/1037group/dousheng/kitex_gen/douyin_publish"
	"github.com/1037group/dousheng/kitex_gen/douyin_user"
	"github.com/1037group/dousheng/pack"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/1037group/dousheng/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *douyin_publish.PublishActionRequest) (resp *douyin_publish.PublishActionResponse, err error) {
	klog.CtxInfof(ctx, "[PublishAction] %+v", req)

	t := time.Now()
	video := sql.Video{
		UserId:             req.ReqUserId,
		VideoPlayUrl:       req.VideoPlayUrl,
		VideoCoverUrl:      req.VideoCoverUrl,
		VideoFavoriteCount: 0,
		VideoCommentCount:  0,
		VideoTitle:         req.Title,
		Ctime:              t,
		Utime:              t,
	}

	err = db.CreateVideo(ctx, db.DB, &video)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, errno.ParamErr
	}

	msg := "publish success"
	return &douyin_publish.PublishActionResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
	}, nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *douyin_publish.PublishListRequest) (resp *douyin_publish.PublishListResponse, err error) {
	klog.CtxInfof(ctx, "[PublishListRequest] %+v", req)

	res, err := db.MGetVideosByUserId(ctx, db.DB, &req.UserId)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}

	userMap := make(map[int64]douyin_user.User)
	var userIDs []int64

	for _, m := range res {
		userIDs = append(userIDs, m.UserId)
	}

	users, err := db.MGetUserByID(ctx, db.DB, userIDs)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}

	for _, m := range users {
		isFollow, err := db.CheckFollow(ctx, db.DB, req.ReqUserId, m.UserId)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return nil, err
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

	return &douyin_publish.PublishListResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		VideoList:  videoList,
	}, nil
}
