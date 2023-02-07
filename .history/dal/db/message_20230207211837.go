package db

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
)

func AddMessage(ctx context.Context, message *sql.Message) {
	klog.CtxInfof(ctx, "[CreateFavorite] user: %+v\n", favor)
	return DB.WithContext(ctx).Create(message).Error
}
