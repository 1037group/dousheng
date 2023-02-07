package logic

import (
	"context"

	"github.com/1037group/dousheng/dal/redis"
	"github.com/1037group/dousheng/kitex_gen/douyin_favorite"
	"github.com/1037group/dousheng/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
)

func FavoriteAction(ctx context.Context, req *douyin_favorite.FavoriteActionRequest) (err error) {
	klog.CtxInfof(ctx, "[logic.FavoriteAction] req: %+v", req)
	// Redis 加锁
	key := GetRelationActionLockKey(req.UserId, req.ToUserId)
	lock := redis.LockAcquire(ctx, key)
	if lock == nil {
		klog.CtxErrorf(ctx, errno.RedisLockFailed.ErrMsg)
		return errno.RedisLockFailed
	}
	defer lock.Release(ctx)
	return nil
}
