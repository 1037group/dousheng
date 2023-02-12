package logic

import (
	"context"
	"time"

	"github.com/1037group/dousheng/dal/db"
	"github.com/1037group/dousheng/kitex_gen/douyin_message"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// 获取message列表，并更新其中的已读信息使用事务
func GetMessageList(ctx context.Context, req *douyin_message.MessageChatRequest) ([]*sql.Message, error) {
	klog.CtxInfof(ctx, "[logic.GetMessageList] req: %+v", req)

	res, err := db.MGetMessageList(ctx, db.DB, &req.UserId, &req.ToUserId)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return res, err
	}

	// 需要事务
	t := time.Now()
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		err = db.MUpdateMessageList(ctx, tx, &req.UserId, &req.ToUserId, t)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return res, err
	}

	return res, nil
}
