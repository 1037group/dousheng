package db

import (
	"context"

	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func MGetFavoriteListByUserId(ctx context.Context, tx *gorm.DB, user_id *int64) ([]*sql.Video, error) {
	klog.CtxInfof(ctx, "[MGetVideosByUserId] last_time: %+v\n", user_id)

	var res []*sql.Video

	query := sql.SQL_VIDEO_USER_ID + " = ?"
	if err := tx.WithContext(ctx).Order(sql.SQL_VIDEO_UTIME+" desc").Find(&res, query, user_id).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetVideosByUserId] res: %+v\n", res)
		return res, err
	}

	return res, nil
}

// CreateFavorite create a favorite for a video
func CreateFavorite(ctx context.Context, favor *sql.Favorite) error {
	klog.CtxInfof(ctx, "[CreateFavorite] user: %+v\n", favor)
	return DB.WithContext(ctx).Create(favor).Error
}