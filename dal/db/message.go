package db

import (
	"context"
	"reflect"
	"time"

	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// 工具函数，用来反转一个interface{}
func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func SendMessage(ctx context.Context, message *sql.Message) error {
	klog.CtxInfof(ctx, "[SendMessage] message: %+v\n", message)
	return DB.WithContext(ctx).Create(message).Error
}

// 获取的应该是5条最近信息+未读信息
func MGetMessageList(ctx context.Context, tx *gorm.DB, user_id *int64, to_user_id *int64) ([]*sql.Message, error) {
	klog.CtxInfof(ctx, "[MGetMessageList] user_id: %+v", user_id)
	res1 := make([]*sql.Message, 0)
	res2 := make([]*sql.Message, 0)
	message_query1 := &sql.Message{
		UserId:   *user_id,
		ToUserId: *to_user_id,
		IsRead:   1,
	}

	if err := tx.WithContext(ctx).Order(sql.SQL_MESSAGE_CTIME + " desc").Where(message_query1).Limit(10).Find(&res1).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetMessageList] res: %+v\n", res1)
		return res1, err
	}
	tx.Model(&res1).Order(sql.SQL_MESSAGE_CTIME + " asc")
	reverseAny(res1)
	query_user_id := sql.SQL_MESSAGE_USER_ID + " = ?"
	query_to_user_id := sql.SQL_MESSAGE_TO_USER_ID + " = ?"
	query_is_read := sql.SQL_MESSAGE_IS_READ + " = ?"

	if err := tx.WithContext(ctx).Order(sql.SQL_MESSAGE_CTIME+" asc").Where(query_user_id, user_id).Where(query_to_user_id, to_user_id).Where(query_is_read, 0).Find(&res2).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetMessageList] res: %+v\n", res2)
		return res2, err
	}
	res := append(res1, res2...)
	return res, nil
}

// 获取最后一条信息
func GetLastMessage(ctx context.Context, tx *gorm.DB, user_id *int64, to_user_id *int64) ([]*sql.Message, error) {
	klog.CtxInfof(ctx, "[GetLastMessage] user_id: %+v", user_id)
	res := make([]*sql.Message, 0)
	message_query1 := &sql.Message{
		UserId:   *user_id,
		ToUserId: *to_user_id,
	}
	message_query2 := &sql.Message{
		UserId:   *to_user_id,
		ToUserId: *user_id,
	}
	if err := tx.WithContext(ctx).Order(sql.SQL_MESSAGE_CTIME + " desc").Where(message_query1).Or(message_query2).Limit(1).Find(&res).Error; err != nil {
		klog.CtxInfof(ctx, "[db.GetLastMessage] res: %+v\n", res)
		return res, err
	}
	return res, nil
}

func MUpdateMessageList(ctx context.Context, tx *gorm.DB, user_id *int64, to_user_id *int64, utime time.Time) error {
	klog.CtxInfof(ctx, "[MUpdateMessageList] user_id: %+v", user_id)
	message_query := &sql.Message{
		UserId:   *user_id,
		ToUserId: *to_user_id,
	}
	return tx.Model(sql.Message{}).Where(message_query).UpdateColumns(map[string]interface{}{sql.SQL_MESSAGE_IS_READ: 1, sql.SQL_MESSAGE_UTIME: utime}).Error
}
