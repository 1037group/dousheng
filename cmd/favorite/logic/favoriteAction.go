package logic

import (
	"context"
	"time"

	"github.com/1037group/dousheng/dal/db"
	"github.com/1037group/dousheng/kitex_gen/douyin_favorite"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// 点赞不存在记录,创建记录,使用事务
func CreateFavoriteAction(ctx context.Context, req *douyin_favorite.FavoriteActionRequest) (err error) {
	klog.CtxInfof(ctx, "[logic.CreateFavoriteAction] req: %+v", req)

	t := time.Now()
	favorite := sql.Favorite{
		UserId:     req.UserId,
		VideoId:    req.VideoId,
		IsFavorite: 1, //1代表点赞，0代表无点赞
		DelState:   0,
		Ctime:      t,
		Utime:      t,
	}
	// 需要事务
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		err = db.CreateFavorite(ctx, tx, &favorite)
		if err != nil {
			return err
		}
		// count的计数放在redis由cronjob异步更新
		//err = db.AddFavoriteCount(ctx, tx, req.VideoId)
		//if err != nil {
		//	return err
		//}
		return err
	})
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return err
	}

	return nil
}

// 点赞存在过记录,更新记录,使用事务
func FavoriteAction(ctx context.Context, req *douyin_favorite.FavoriteActionRequest) (err error) {
	klog.CtxInfof(ctx, "[logic.FavoriteAction] req: %+v", req)

	// 需要事务
	t := time.Now()
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		err = db.UpdateFavorite(ctx, tx, req.UserId, req.VideoId, t, req.ActionType)
		// redis处理点赞计数
		//if req.ActionType == 1 {
		//	err = db.AddFavoriteCount(ctx, tx, req.VideoId)
		//	if err != nil {
		//		return err
		//	}
		//} else {
		//	err = db.MinusFavoriteCount(ctx, tx, req.VideoId)
		//	if err != nil {
		//		return err
		//	}
		//}
		return err
	})
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return err
	}

	return nil
}
