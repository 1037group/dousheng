package logic

import (
	"context"
	"fmt"

	"github.com/1037group/dousheng/dal/db"
	"github.com/1037group/dousheng/dal/redis"
	"github.com/1037group/dousheng/kitex_gen/douyin_favorite"
	"github.com/1037group/dousheng/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func GetRelationActionLockKey(userId, VideoId int64) string {
	return fmt.Sprintf("RelationActionLock: %+v-%+v", userId, VideoId)
}

func FavoriteAction(ctx context.Context, req *douyin_favorite.FavoriteActionRequest) (err error) {
	klog.CtxInfof(ctx, "[logic.FavoriteAction] req: %+v", req)
	// Redis 加锁
	key := GetRelationActionLockKey(req.UserId, req.VideoId)
	lock := redis.LockAcquire(ctx, key)
	if lock == nil {
		klog.CtxErrorf(ctx, errno.RedisLockFailed.ErrMsg)
		return errno.RedisLockFailed
	}
	defer lock.Release(ctx)
	// 需要事务
	err = db.DB.Transaction(func(tx *gorm.DB) error {
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
