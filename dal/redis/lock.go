package redis

import (
	"context"
	"fmt"
	"github.com/bsm/redislock"
	"github.com/cloudwego/kitex/pkg/klog"

	"time"
)

func LockAcquire(ctx context.Context, key string) *redislock.Lock {
	lock, err := locker.Obtain(ctx, key, 900*time.Millisecond, nil)
	if err == redislock.ErrNotObtained {
		fmt.Println("Could not obtain lock!")
	} else if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil
	}
	return lock
}
