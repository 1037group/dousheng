package main

import (
	"context"
	"time"

	"github.com/1037group/dousheng/dal/db"
	douyin_favorite "github.com/1037group/dousheng/kitex_gen/douyin_favorite"
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
		IsFavorite: 0, //1代表点赞，0代表无点赞
		DelState:   0,
		Utime:      t,
	}
	//应该做一次gorm查询目前这个视频有没有被这个人点赞，然后再执行操作
	if req.ActionType == 1 { //点赞
		favorite.IsFavorite = 1
		err = db.CreateFavorite(ctx, &favorite)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return nil, errno.ParamErr
		}
	} else if req.ActionType == 2 { //取消点赞
		favorite.IsFavorite = 0
		err = db.CancelFavorite(ctx, req.UserId, req.VideoId)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return nil, errno.ParamErr
		}
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

	res, err := db.MGetVideosByUserId(ctx, db.DB, &req.UserId)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}

	userMap := make(map[int64]sql.User)
	var userIDs []int64

	for _, m := range res {
		userIDs = append(userIDs, m.UserId)
	}
	return
}