package db

import (
	"context"

	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SendMessage(ctx context.Context, message *sql.Message) error {
	klog.CtxInfof(ctx, "[SendMessage] message: %+v\n", message)
	queryUid := sql.SQL_MESSAGE_USER_ID + " = ?"
	queryToUid := sql.SQL_MESSAGE_TO_USER_ID + " = ?"

	return DB.WithContext(ctx).Create(message).Error
}
