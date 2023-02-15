package logic

import (
	"context"
	"fmt"

	"github.com/1037group/dousheng/dal/db"
	"github.com/1037group/dousheng/dal/redis"
	"github.com/1037group/dousheng/kitex_gen/douyin_relation"
	"github.com/1037group/dousheng/pkg/errno"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func GetRelationActionLockKey(userId, toUserId int64) string {
	return fmt.Sprintf("RelationActionLock: %+v-%+v", userId, toUserId)
}

func RelationAction(ctx context.Context, req *douyin_relation.RelationActionRequest) (err error) {
	klog.CtxInfof(ctx, "[logic.RelationAction] req: %+v", req)

	// Redis 加锁
	key := GetRelationActionLockKey(req.ReqUserId, req.ToUserId)
	lock := redis.LockAcquire(ctx, key)
	if lock == nil {
		klog.CtxErrorf(ctx, errno.RedisLockFailed.ErrMsg)
		return errno.RedisLockFailed
	}
	defer lock.Release(ctx)

	// 需要事务
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		err = db.UpdateRelation(ctx, tx, req.ReqUserId, req.ToUserId, req.ActionType)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return err
		}
		// 由redis来处理
		// if req.ActionType == 1 {
		// 	err = db.AddFollowerCount(ctx, tx, req.ToUserId)
		// 	if err != nil {
		// 		klog.CtxErrorf(ctx, err.Error())
		// 		return err
		// 	}
		// 	err = db.AddFollowCount(ctx, tx, req.ReqUserId)
		// 	if err != nil {
		// 		klog.CtxErrorf(ctx, err.Error())
		// 		return err
		// 	}
		// } else {
		// 	err = db.MinusFollowerCount(ctx, tx, req.ToUserId)
		// 	if err != nil {
		// 		klog.CtxErrorf(ctx, err.Error())
		// 		return err
		// 	}
		// 	err = db.MinusFollowCount(ctx, tx, req.ReqUserId)
		// 	if err != nil {
		// 		klog.CtxErrorf(ctx, err.Error())
		// 		return err
		// 	}
		// }
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
