package db

import (
	"context"
	"fmt"
	"time"

	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func SendMessage(ctx context.Context, tx *gorm.DB, message *sql.Message) error {
	klog.CtxInfof(ctx, "[SendMessage] message: %+v\n", message)
	return tx.WithContext(ctx).Create(message).Error
}

func MGetMessageList(ctx context.Context, tx *gorm.DB, store_by_user_id *int64, user_id *int64, to_user_id *int64) ([]*sql.Message, error) {
	klog.CtxInfof(ctx, "[MGetMessageList] user_id: %+v", user_id)
	res := make([]*sql.Message, 0)

	query := fmt.Sprintf("%s = ? and %s = ? and %s = ? and %s = ?", sql.SQL_MESSAGE_STORE_BY_USER_ID, sql.SQL_MESSAGE_USER_ID, sql.SQL_MESSAGE_TO_USER_ID, sql.SQL_MESSAGE_IS_READ)
	if err := tx.WithContext(ctx).Order(sql.SQL_MESSAGE_UTIME+" desc").Where(query, store_by_user_id, user_id, to_user_id, 0).Find(&res).Error; err != nil {
		return res, err
	}
	klog.CtxInfof(ctx, "[MGetMessageList] res: %+v\n", res)
	return res, nil
}

func MUpdateMessageListMESSAGE_IS_READ(ctx context.Context, tx *gorm.DB, store_by_user_id *int64, user_id *int64, to_user_id *int64, utime time.Time) error {
	klog.CtxInfof(ctx, "[MUpdateMessageListMESSAGE_IS_READ] user_id: %+v", user_id)

	query := fmt.Sprintf("%s = ? and %s = ? and %s = ?", sql.SQL_MESSAGE_STORE_BY_USER_ID, sql.SQL_MESSAGE_USER_ID, sql.SQL_MESSAGE_TO_USER_ID)

	return tx.Model(sql.Message{}).Where(query, store_by_user_id, user_id, to_user_id).UpdateColumns(map[string]interface{}{sql.SQL_MESSAGE_IS_READ: 1, sql.SQL_MESSAGE_UTIME: utime}).Error
}

func MUpdateMessageListMESSAGE_IS_UNREAD(ctx context.Context, tx *gorm.DB, store_by_user_id *int64, utime time.Time) error {
	klog.CtxInfof(ctx, "[MUpdateMessageListMESSAGE_IS_UNREAD] store_by_user_id: %+v", store_by_user_id)

	query := fmt.Sprintf("%s = ?", sql.SQL_MESSAGE_STORE_BY_USER_ID)

	return tx.Model(sql.Message{}).Where(query, store_by_user_id).UpdateColumns(map[string]interface{}{sql.SQL_MESSAGE_IS_READ: 0, sql.SQL_MESSAGE_UTIME: utime}).Error
}
