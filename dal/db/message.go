package db

import (
	"context"
	"time"

	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func SendMessage(ctx context.Context, tx *gorm.DB, message *sql.Message) error {
	klog.CtxInfof(ctx, "[SendMessage] message: %+v\n", message)
	return tx.WithContext(ctx).Create(message).Error
}

func MGetMessageList(ctx context.Context, tx *gorm.DB, user_id *int64, to_user_id *int64) ([]*sql.Message, error) {
	klog.CtxInfof(ctx, "[MGetMessageList] user_id: %+v", user_id)
	res := make([]*sql.Message, 0)
	query_user_id := sql.SQL_MESSAGE_USER_ID + " = ?"
	query_to_user_id := sql.SQL_MESSAGE_TO_USER_ID + " = ?"
	queryAppend_is_read := sql.SQL_MESSAGE_IS_READ + " = ?"
	if err := tx.WithContext(ctx).Order(sql.SQL_MESSAGE_UTIME+" desc").Where(query_user_id, user_id).Where(query_to_user_id, to_user_id).Where(queryAppend_is_read, 0).Find(&res).Error; err != nil {
		return res, err
	}
	klog.CtxInfof(ctx, "[MGetMessageList] res: %+v\n", res)
	return res, nil
}

func MUpdateMessageList(ctx context.Context, tx *gorm.DB, user_id *int64, to_user_id *int64, utime time.Time) error {
	klog.CtxInfof(ctx, "[MUpdateMessageList] user_id: %+v", user_id)
	query_user_id := sql.SQL_MESSAGE_USER_ID + " = ?"
	query_to_user_id := sql.SQL_MESSAGE_TO_USER_ID + " = ?"
	return tx.Model(sql.Message{}).Where(query_user_id, user_id).Where(query_to_user_id, to_user_id).UpdateColumns(map[string]interface{}{sql.SQL_MESSAGE_IS_READ: 1, sql.SQL_MESSAGE_UTIME: utime}).Error
}
