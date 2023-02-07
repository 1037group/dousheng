package db

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
)

func SendMessage(ctx context.Context, message *sql.Message) error {
	klog.CtxInfof(ctx, "[SendMessage] message: %+v\n", message)
	return DB.WithContext(ctx).Create(message).Error
}
