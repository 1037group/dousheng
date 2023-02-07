package db

import (
	"context"
	"time"

	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// MGetVideosByLastTime multiple get list of videos
func MGetVideosByLastTime(ctx context.Context, tx *gorm.DB, last_time *int64) ([]*sql.Video, error) {
	klog.CtxInfof(ctx, "[MGetVideosByLastTime] last_time: %+v\n", last_time)

	var res []*sql.Video

	if last_time == nil || *last_time == 0 {
		cur_time := time.Now().UnixMilli()
		klog.CtxInfof(ctx, "cur_time %+v", cur_time)
		klog.CtxInfof(ctx, "cur_time %+v", cur_time)

		last_time = &cur_time
	}

	query := sql.SQL_VIDEO_UTIME + " < ?"
	if err := tx.WithContext(ctx).Order(sql.SQL_VIDEO_UTIME+" desc").Find(&res, query, time.UnixMilli(*last_time)).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetVideosByLastTime] res: %+v\n", res)
		return res, err
	}
	return res, nil
}

func MGetVideosByUserId(ctx context.Context, tx *gorm.DB, user_id *int64) ([]*sql.Video, error) {
	klog.CtxInfof(ctx, "[MGetVideosByUserId] last_time: %+v\n", user_id)

	var res []*sql.Video

	query := sql.SQL_VIDEO_USER_ID + " = ?"
	if err := tx.WithContext(ctx).Order(sql.SQL_VIDEO_UTIME+" desc").Find(&res, query, user_id).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetVideosByUserId] res: %+v\n", res)
		return res, err
	}

	return res, nil
}

// CreateVideo create video
func CreateVideo(ctx context.Context, tx *gorm.DB, video *sql.Video) error {
	klog.CtxInfof(ctx, "[CreateVideo] video: %+v\n", video)
	return tx.WithContext(ctx).Create(video).Error
}

func AddFavoriteCount(ctx context.Context, tx *gorm.DB, video_id *int64) {
	klog.CtxInfof(ctx, "[db.AddFavoriteCount] video_id : %+v\n", video_id)

	user := &sql.Video{VideoId: video_id}
	return tx.Model(&user).UpdateColumn(sql.SQL_USER_USER_FOLLOW_COUNT, gorm.Expr(sql.SQL_USER_USER_FOLLOW_COUNT+" + ?", 1)).Error
}

func MinusFavoriteCount(ctx context.Context, tx *gorm.DB, video_id *int64) {

}
