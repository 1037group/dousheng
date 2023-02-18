package logic

import (
	"context"
	"sort"
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

	var err error
	var resA2B []*sql.Message
	var resB2A []*sql.Message

	// 需要事务
	t := time.Now()
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		resA2B, err = db.MGetMessageList(ctx, tx, &req.UserId, &req.UserId, &req.ToUserId)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return err
		}

		resB2A, err = db.MGetMessageList(ctx, tx, &req.UserId, &req.ToUserId, &req.UserId)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return err
		}

		err = db.MUpdateMessageListMESSAGE_IS_READ(ctx, tx, &req.UserId, &req.UserId, &req.ToUserId, t)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return err
		}

		err = db.MUpdateMessageListMESSAGE_IS_READ(ctx, tx, &req.UserId, &req.ToUserId, &req.UserId, t)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}

	res := sql.MessageSort{}
	res = append(res, resA2B...)
	res = append(res, resB2A...)

	// 排序
	sort.Sort(res)
	return res, nil
}
