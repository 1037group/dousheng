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

func MGetMessageList(ctx context.Context, tx *gorm.DB, user_id int64, to_user_id int64) error {
	klog.CtxInfof(ctx, "[MGetMessageList] userId: %+v", user_id)
	res := make([]*sql.Message, 0)
	query := sql.SQL_MESSAGE_USER_ID + " = ?"
	queryAppend1 := sql.SQL_MESSAGE_TO_USER_ID + " = ?"
	results := tx.WithContext(ctx).Order(sql.SQL_FAVORITE_UTIME+" desc").Where(query, user_id).Where(queryAppend1,to_user_id).Find().Error; err != nil {
		klog.CtxInfof(ctx, "[MGetMessageList] midRes: %+v\n", midRes)
		return err
	}
}
