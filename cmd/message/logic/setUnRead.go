package logic

import (
	"context"
	"github.com/1037group/dousheng/dal/db"
	"github.com/1037group/dousheng/kitex_gen/douyin_message"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

func MessageSetUnRead(ctx context.Context, req *douyin_message.MessageSetUnReadRequest) (err error) {
	klog.CtxInfof(ctx, "[logic.MessageSetUnRead] req: %+v", req)

	t := time.Now()

	err = db.DB.Transaction(func(tx *gorm.DB) error {
		err = db.MUpdateMessageListMESSAGE_IS_UNREAD(ctx, tx, &req.ReqUserId, t)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return err
		}
		return err
	})
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return err
	}

	return nil
}
