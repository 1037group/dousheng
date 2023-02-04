package db

import (
	"context"

	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// 判断某用户是不是已经对该视频点赞了
func MIsFavoriteByUserId(ctx context.Context, tx *gorm.DB, user_id *int64, video_id *int64) ([]*sql.Favorite, error) {
	klog.CtxInfof(ctx, "[MIsFavoriteByUserId] user_id: %+v\n", user_id)
	var res []*sql.Favorite
	query := sql.SQL_FAVORITE_USER_ID + " = ?"
	query1 := sql.SQL_FAVORITE_VIDEO_ID + " = ?"
	if err := tx.WithContext(ctx).Order(sql.SQL_FAVORITE_UTIME+" desc").Where(query, user_id).Where(query1, video_id).Find(&res).Error; err != nil {
		klog.CtxInfof(ctx, "[MIsFavoriteByUserId] res: %+v\n", res)
		return res, err
	}
	return res, nil
}

// 根据user_id找到该用户喜欢的video_id
func MGetVideoIdtByUserId(ctx context.Context, tx *gorm.DB, user_id *int64) ([]*sql.Favorite, error) {
	klog.CtxInfof(ctx, "[MGetVideoIdtByUserId] user_id: %+v\n", user_id)

	var res []*sql.Favorite

	query := sql.SQL_FAVORITE_USER_ID + " = ?"
	query1 := sql.SQL_FAVORITE_ISFAVORITE + " = ?"
	if err := tx.WithContext(ctx).Order(sql.SQL_FAVORITE_UTIME+" desc").Where(query, user_id).Where(query1, 1).Find(&res).Error; err != nil {
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
func CancelFavorite(ctx context.Context, tx *gorm.DB, user_id int64, video_id int64) error {
	klog.CtxInfof(ctx, "[CancelFavorite] user_id: %+v\n", user_id)
	favorite := &sql.Favorite{
		UserId:  user_id,
		VideoId: video_id,
	}
	return tx.Model(&favorite).UpdateColumn(sql.SQL_FAVORITE_ISFAVORITE, 0).Error
}
