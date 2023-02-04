package main

import (
	"context"
	"time"

	"github.com/1037group/dousheng/dal/db"
	douyin_favorite "github.com/1037group/dousheng/kitex_gen/douyin_favorite"
	"github.com/1037group/dousheng/pack"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/1037group/dousheng/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *douyin_favorite.FavoriteActionRequest) (resp *douyin_favorite.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	klog.CtxInfof(ctx, "[FavoriteAction] %+v", req)
	t := time.Now()
	favorite := sql.Favorite{
		UserId:     req.UserId,
		VideoId:    req.VideoId,
		IsFavorite: 1, //1代表点赞，0代表无点赞
		DelState:   0,
		Utime:      t,
	}
	//做一次gorm查询目前这个视频有没有被这个人点赞,然后再执行操作
	res, err := db.MIsFavoriteByUserId(ctx, db.DB, &req.UserId, &req.VideoId)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, errno.ParamErr
	}
	klog.CtxInfof(ctx, "[FavoriteAction][res==nil] %v", res == nil)

	//该用户已经点赞了该视频
	var isFavor int64
	for _, m := range res {
		isFavor = int64(m.IsFavorite)
	}
	klog.CtxInfof(ctx, "[FavoriteAction][isFavor] %v", isFavor)
	if (isFavor == 1 && req.ActionType == 1) || (isFavor == 0 && req.ActionType == 2) {
		msg := "Duplicate action type for this user"
		return &douyin_favorite.FavoriteActionResponse{
			StatusCode: 0,
			StatusMsg:  &msg,
		}, nil
	}
	if res == nil && req.ActionType == 1 { //创建点赞
		err = db.CreateFavorite(ctx, &favorite)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return nil, errno.ParamErr
		}
	} else if res != nil && req.ActionType == 1 { //某个用户已经取消了点赞，并想重新点赞
		err = db.DoFavorite(ctx, db.DB, req.UserId, req.VideoId, t)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return nil, errno.ParamErr
		}
	} else if req.ActionType == 2 { //取消点赞
		err = db.CancelFavorite(ctx, db.DB, req.UserId, req.VideoId, t)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return nil, errno.ParamErr
		}
	} else {
		msg := "Incorrect action type"
		return &douyin_favorite.FavoriteActionResponse{
			StatusCode: 0,
			StatusMsg:  &msg,
		}, nil
	}
	msg := "Favorite success"
	return &douyin_favorite.FavoriteActionResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
	}, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *douyin_favorite.FavoriteListRequest) (resp *douyin_favorite.FavoriteListResponse, err error) {
	// TODO: Your code here...
	klog.CtxInfof(ctx, "[FavoriteList] %+v", req)

	//首先查favorite表，根据点赞的UserId获取对应的Favorite表
	res, err := db.MGetVideoIdtByUserId(ctx, db.DB, &req.UserId)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}
	var videoIDs []int64
	//将获取的VideoId打包成数组
	for _, m := range res {
		videoIDs = append(videoIDs, m.VideoId)
	}

	//查Video表，根据VideoId获取对应的Video信息
	videos, err := db.MGetVideosByVideoId(ctx, db.DB, videoIDs)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}
	var userIDs []int64
	for _, m := range videos {
		userIDs = append(userIDs, m.UserId)
	}

	userMap := make(map[int64]sql.User)
	//查User表，根据UserId获取对应的user信息
	users, err := db.MGetUserByID(ctx, db.DB, userIDs)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}
	//打包成信息
	for _, m := range users {
		userMap[m.UserId] = *m
	}
	videoList := pack.Videos(videos, userMap)
	klog.CtxInfof(ctx, "[videoList] %+v", videoList)
	msg := "FavoriteList success"
	return &douyin_favorite.FavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
		VideoList:  videoList,
	}, nil
}
