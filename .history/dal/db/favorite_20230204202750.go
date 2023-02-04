package db

import (
	"context"

	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)


//根据user_id找到该用户喜欢的video_id
func MGetVideoIdtByUserId(ctx context.Context, tx *gorm.DB, user_id *int64) ([]*sql.Favorite, error) {
	klog.CtxInfof(ctx, "[MGetVideoIdtByUserId] user_id: %+v\n", user_id)

	var res []*sql.Favorite

	query := sql.SQL_FAVORITE_USER_ID + " = ?"
	if err := tx.WithContext(ctx).Order(sql.SQL_FAVORITE_UTIME+" desc").Find(&res, query, user_id).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetFavoriteListByUserId] res: %+v\n", res)
		return res, err
	}

	return res, nil
}

//根据video_id找到对应的video数据
func MGetVideotByVideoId((ctx context.Context, tx *gorm.DB, video_id *int64) ([]*sql.Video, error)) {
	klog.CtxInfof(ctx, "[MGetVideotByVideoId] video_id: %+v\n", video_id)

	var res []*sql.Video

	query := sql.SQL_VIDEO_VIDEO_ID  + " = ?"
	if err := tx.WithContext(ctx).Order(sql.SQL_FAVORITE_UTIME+" desc").Find(&res, query, user_id).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetFavoriteListByUserId] res: %+v\n", res)
		return res, err
	}

	return res, nil
}

func MGetIsFavoriteByUserId(ctx context.Context, tx *gorm.DB, user_id *int64, video_id *int64) (bool, error) {
	klog.CtxInfof(ctx, "[MGetIsFavoriteByUserId] user_id: %+v video_id: %+v\n", user_id, video_id)

	var res []*sql.Favorite
	// query := sql.SQL_FAVORITE_USER_ID + " = ? AND "+ sql.SQL_FAVORITE_VIDEO_ID+=" = ? "
	tx.WithContext(ctx).Where("user_id = ? AND video_id = ? ", user_id, video_id).Find(&res)
	query := sql.SQL_FAVORITE_USER_ID + " in ?"

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