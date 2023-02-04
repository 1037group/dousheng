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
	favorite := sql.Favorite{
		UserId:     req.UserId,
		VideoId:    req.VideoId,
		IsFavorite: 0, //1代表点赞，0代表无点赞
		DelState:   0,
		Utime:      t,
	}
	t := time.Now()
	if req.ActionType == 1 {
		favorite.IsFavorite = 1
	} else if req.ActionType == 2 {
		favorite.IsFavorite = 0
	}

	err = db.CreateFavorite(ctx, &favorite)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, errno.ParamErr
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

	return
}
