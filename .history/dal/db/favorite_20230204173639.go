package db

import (
	"context"

	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func MGetVideoIdtByUserId(ctx context.Context, tx *gorm.DB, user_id *int64) ([]*sql.Favorite, error) {
	klog.CtxInfof(ctx, "[MGetFavoriteListByUserId] last_time: %+v\n", user_id)

	var res []*sql.Favorite

	query := sql.SQL_VIDEO_USER_ID + " = ?"
	if err := tx.WithContext(ctx).Order(sql.SQL_FAVORITE_UTIME+" desc").Find(&res, query, user_id).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetFavoriteListByUserId] res: %+v\n", res)
		return res, err
	}

	return res, nil
}

// CreateFavorite create a favorite for a video
func CreateFavorite(ctx context.Context, favor *sql.Favorite) error {
	klog.CtxInfof(ctx, "[CreateFavorite] user: %+v\n", favor)
	return DB.WithContext(ctx).Create(favor).Error
}

// CancelFavorite cancel favorite for a video
func CancelFavorite(ctx context.Context, favor *sql.Favorite) error {
	klog.CtxInfof(ctx, "[CancelFavorite] user: %+v\n", favor)
	return DB.WithContext(ctx).Save(favor).Error
}
