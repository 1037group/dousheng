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

// 根据user_id找到该用户喜欢的video_id
func MGetVideoIdtByUserId(ctx context.Context, tx *gorm.DB, user_id *int64) ([]*sql.Video, error) {
	klog.CtxInfof(ctx, "[MGetVideoIdtByUserId] user_id: %+v\n", user_id)

	var midRes []*sql.Favorite
	res := make([]*sql.Video, 0)

	query := sql.SQL_FAVORITE_USER_ID + " = ?"
	queryAppend1 := sql.SQL_FAVORITE_ISFAVORITE + " = ?"
	if err := tx.WithContext(ctx).Order(sql.SQL_FAVORITE_UTIME+" desc").Where(query, user_id).Where(queryAppend1, 1).Find(&midRes).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetVideoIdtByUserId] midRes: %+v\n", midRes)
		return res, err
	}
	//根据得到的videoid，继续处理得到user信息，在这个函数里做完逻辑，不要用两个函数
	var videoIDs []int64
	//将获取的VideoId打包成数组
	for _, m := range midRes {
		videoIDs = append(videoIDs, m.VideoId)
	}
	queryAppend2 := sql.SQL_VIDEO_VIDEO_ID + " = ?"
	if err := tx.WithContext(ctx).Order(sql.SQL_FAVORITE_UTIME+" desc").Where(queryAppend2, videoIDs).Find(&res).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetVideoIdtByUserId] res: %+v\n", res)
		return res, err
	}

	return res, nil
}

// 给某个用户创建一条点赞信息
func CreateFavorite(ctx context.Context, favor *sql.Favorite) error {
	klog.CtxInfof(ctx, "[CreateFavorite] user: %+v\n", favor)
	return DB.WithContext(ctx).Create(favor).Error
}

// 给某个用户进行一个取消点赞操作
func CancelFavorite(ctx context.Context, tx *gorm.DB, user_id int64, video_id int64, utime time.Time) error {
	klog.CtxInfof(ctx, "[CancelFavorite] user_id: %+v\n", user_id)

	query := sql.SQL_FAVORITE_USER_ID + " = ?"
	query1 := sql.SQL_FAVORITE_VIDEO_ID + " = ?"
	return tx.Model(sql.Favorite{}).Where(query, user_id).Where(query1, video_id).UpdateColumns(map[string]interface{}{sql.SQL_FAVORITE_ISFAVORITE: 0, sql.SQL_FAVORITE_UTIME: utime}).Error
}

// 给某个用户进行一个点赞操作
func DoFavorite(ctx context.Context, tx *gorm.DB, user_id int64, video_id int64, utime time.Time) error {
	klog.CtxInfof(ctx, "[DoFavorite] user_id: %+v\n", user_id)

	query := sql.SQL_FAVORITE_USER_ID + " = ?"
	query1 := sql.SQL_FAVORITE_VIDEO_ID + " = ?"
	return tx.Model(sql.Favorite{}).Where(query, user_id).Where(query1, video_id).UpdateColumns(map[string]interface{}{sql.SQL_FAVORITE_ISFAVORITE: 1, sql.SQL_FAVORITE_UTIME: utime}).Error
}
