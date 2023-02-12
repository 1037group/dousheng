package db

import (
	"context"
	"time"

	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// 判断某用户是不是已经对该视频点赞了
func MIsFavoriteByUserId(ctx context.Context, tx *gorm.DB, user_id *int64, video_id *int64) ([]*sql.Favorite, error) {
	klog.CtxInfof(ctx, "[MIsFavoriteByUserId] user_id: %+v\n", user_id)
	var res []*sql.Favorite
	query := sql.SQL_FAVORITE_USER_ID + " = ?"
	queryAppend := sql.SQL_FAVORITE_VIDEO_ID + " = ?"
	if err := tx.WithContext(ctx).Order(sql.SQL_FAVORITE_UTIME+" desc").Where(query, user_id).Where(queryAppend, video_id).Find(&res).Error; err != nil {
		klog.CtxInfof(ctx, "[MIsFavoriteByUserId] res: %+v\n", res)
		return res, err
	}
	return res, nil
}

// 根据user_id找到该用户喜欢的video信息
func MGetFavoriteVideosByUserId(ctx context.Context, tx *gorm.DB, user_id *int64) ([]*sql.Video, error) {
	klog.CtxInfof(ctx, "[MGetFavoriteVideosByUserId] user_id: %+v\n", user_id)

	var midRes []*sql.Favorite
	res := make([]*sql.Video, 0)

	query := sql.SQL_FAVORITE_USER_ID + " = ?"
	queryAppend1 := sql.SQL_FAVORITE_ISFAVORITE + " = ?"
	if err := tx.WithContext(ctx).Order(sql.SQL_FAVORITE_UTIME+" desc").Where(query, user_id).Where(queryAppend1, 1).Find(&midRes).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetFavoriteVideosByUserId] midRes: %+v\n", midRes)
		return res, err
	}
	var videoIDs []int64
	//将获取的VideoId打包成数组
	for _, m := range midRes {
		videoIDs = append(videoIDs, m.VideoId)
	}
	queryAppend2 := sql.SQL_VIDEO_VIDEO_ID + " in ?"
	if err := tx.WithContext(ctx).Order(sql.SQL_FAVORITE_UTIME+" desc").Where(queryAppend2, videoIDs).Find(&res).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetFavoriteVideosByUserId] res: %+v\n", res)
		return res, err
	}

	return res, nil
}

// 给某个用户创建一条点赞信息
func CreateFavorite(ctx context.Context, tx *gorm.DB, favor *sql.Favorite) error {
	klog.CtxInfof(ctx, "[CreateFavorite] favor: %+v\n", favor)
	return tx.WithContext(ctx).Create(favor).Error
}

// 更新点赞信息
func UpdateFavorite(ctx context.Context, tx *gorm.DB, user_id int64, video_id int64, utime time.Time, action_type int32) error {
	klog.CtxInfof(ctx, "[CancelFavorite] user_id: %+v\n", user_id)

	query := sql.SQL_FAVORITE_USER_ID + " = ?"
	query1 := sql.SQL_FAVORITE_VIDEO_ID + " = ?"
	if action_type == 1 {
		return tx.Model(sql.Favorite{}).Where(query, user_id).Where(query1, video_id).UpdateColumns(map[string]interface{}{sql.SQL_FAVORITE_ISFAVORITE: 1, sql.SQL_FAVORITE_UTIME: utime}).Error
	}
	return tx.Model(sql.Favorite{}).Where(query, user_id).Where(query1, video_id).UpdateColumns(map[string]interface{}{sql.SQL_FAVORITE_ISFAVORITE: 0, sql.SQL_FAVORITE_UTIME: utime}).Error
}
