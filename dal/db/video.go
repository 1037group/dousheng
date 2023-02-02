package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)
import "github.com/1037group/dousheng/pkg/configs/sql"

// MGetVideosByLastTime multiple get list of videos
func MGetVideosByLastTime(ctx context.Context, last_time *int64) ([]*sql.Video, error) {
	klog.CtxInfof(ctx, "[MGetVideosByLastTime] last_time: %+v\n", last_time)

	var res []*sql.Video

	if last_time == nil || *last_time == 0 {
		cur_time := time.Now().UnixMilli()
		klog.CtxInfof(ctx, "cur_time %+v", cur_time)
		klog.CtxInfof(ctx, "cur_time %+v", cur_time)

		last_time = &cur_time
	}

	query := sql.SQL_VIDEO_UTIME + " < ?"
	if err := DB.WithContext(ctx).Order(sql.SQL_VIDEO_UTIME+" desc").Find(&res, query, time.UnixMilli(*last_time)).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetVideosByLastTime] res: %+v\n", res)
		return res, err
	}
	return res, nil
}

func MGetVideosByUserId(ctx context.Context, user_id *int64) ([]*sql.Video, error) {
	klog.CtxInfof(ctx, "[MGetVideosByUserId] last_time: %+v\n", user_id)

	var res []*sql.Video

	query := sql.SQL_VIDEO_USER_ID + " = ?"
	if err := DB.WithContext(ctx).Order(sql.SQL_VIDEO_UTIME+" desc").Find(&res, query, user_id).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetVideosByUserId] res: %+v\n", res)
		return res, err
	}

	return res, nil
}

// CreateVideo create video
func CreateVideo(ctx context.Context, video *sql.Video) error {
	klog.CtxInfof(ctx, "[CreateVideo] video: %+v\n", video)
	return DB.WithContext(ctx).Create(video).Error
}
