package db

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
)

func SendMessage(ctx context.Context, message *sql.Message) {
	klog.CtxInfof(ctx, "[SendMessage] user: %+v\n", favor)
	return DB.WithContext(ctx).Create(message).Error
}
