package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/1037group/dousheng/dal/db"
	"github.com/1037group/dousheng/dal/redis"
	"github.com/1037group/dousheng/kitex_gen/douyin_favorite"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/1037group/dousheng/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func GetFavoriteActionLockKey(userId, VideoId int64) string {
	return fmt.Sprintf("FavoriteActionLock: %+v-%+v", userId, VideoId)
}

// 点赞不存在记录,创建记录,使用事务
func CreateFavoriteAction(ctx context.Context, req *douyin_favorite.FavoriteActionRequest) (err error) {
	klog.CtxInfof(ctx, "[logic.CreateFavoriteAction] req: %+v", req)
	// Redis 加锁
	key := GetFavoriteActionLockKey(req.UserId, req.VideoId)
	lock := redis.LockAcquire(ctx, key)
	if lock == nil {
		klog.CtxErrorf(ctx, errno.RedisLockFailed.ErrMsg)
		return errno.RedisLockFailed
	}
	defer lock.Release(ctx)

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
		err = db.CreateFavorite(ctx, &favorite)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return err
		}

		err = db.AddFavoriteCount(ctx, tx, req.VideoId)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return err
		}
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

// 点赞存在过记录,更新记录,使用事务
func FavoriteAction(ctx context.Context, req *douyin_favorite.FavoriteActionRequest) (err error) {
	klog.CtxInfof(ctx, "[logic.FavoriteAction] req: %+v", req)
	// Redis 加锁
	key := GetFavoriteActionLockKey(req.UserId, req.VideoId)
	lock := redis.LockAcquire(ctx, key)
	if lock == nil {
		klog.CtxErrorf(ctx, errno.RedisLockFailed.ErrMsg)
		return errno.RedisLockFailed
	}
	defer lock.Release(ctx)
	// 需要事务
	t := time.Now()
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		err = db.UpdateFavorite(ctx, tx, req.UserId, req.VideoId, t, req.ActionType)
		if req.ActionType == 1 {
			err = db.AddFavoriteCount(ctx, tx, req.VideoId)
			if err != nil {
				klog.CtxErrorf(ctx, err.Error())
				return err
			}
		} else {
			err = db.MinusFavoriteCount(ctx, tx, req.VideoId)
			if err != nil {
				klog.CtxErrorf(ctx, err.Error())
				return err
			}
		}
		return err
	})
	if err != nil {
		return err
	}

	return nil
}
