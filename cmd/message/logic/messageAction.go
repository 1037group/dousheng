package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/1037group/dousheng/dal/db"
	"github.com/1037group/dousheng/dal/redis"
	"github.com/1037group/dousheng/kitex_gen/douyin_message"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/1037group/dousheng/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func GetMessageActionLockKey(userId, ToUserId int64) string {
	return fmt.Sprintf("FavoriteActionLock: %+v-%+v", userId, ToUserId)
}

// 创建信息，使用事务
func MessageAction(ctx context.Context, req *douyin_message.MessageActionRequest) (err error) {
	klog.CtxInfof(ctx, "[logic.MessageAction] req: %+v", req)
	// Redis 加锁
	key := GetMessageActionLockKey(req.UserId, req.ToUserId)
	lock := redis.LockAcquire(ctx, key)
	if lock == nil {
		klog.CtxErrorf(ctx, errno.RedisLockFailed.ErrMsg)
		return errno.RedisLockFailed
	}
	defer lock.Release(ctx)
	// 需要事务
	t := time.Now()
	message := sql.Message{
		UserId:         req.UserId,
		ToUserId:       req.ToUserId,
		CommentContent: req.Content,
		IsRead:         0,
		Ctime:          t,
		Utime:          t,
	}

	err = db.DB.Transaction(func(tx *gorm.DB) error {
		err := db.SendMessage(ctx, &message)
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
