package db

import (
	"context"

	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func SendMessage(ctx context.Context, message *sql.Message) error {
	klog.CtxInfof(ctx, "[SendMessage] message: %+v\n", message)
	return DB.WithContext(ctx).Create(message).Error
}

func MGetMessageList(ctx context.Context, tx *gorm.DB, userId int64, touserId int64) error {
	klog.CtxInfof(ctx, "[db.MGetMessageList] userId: %+v", userId)
}
