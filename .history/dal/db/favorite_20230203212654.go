package db

import (
	"context"

	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
)

// CreateFavorite create a favorite for a video
func CreateFavorite(ctx context.Context, user *sql.Favorite) error {
	klog.CtxInfof(ctx, "[CreateUser] userName: %+v\n", user)
	return DB.WithContext(ctx).Create(user).Error
}
