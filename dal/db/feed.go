package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)
import "github.com/1037group/dousheng/pkg/configs/sql"

// MGetVideos multiple get list of videos
func MGetVideos(ctx context.Context, last_time *int64) ([]*sql.Video, error) {
	klog.CtxInfof(ctx, "[MGetVideos] last_time: %+v\n", last_time)

	var res []*sql.Video

	if last_time == nil || *last_time == 0 {
		cur_time := time.Now().UnixMilli()
		klog.CtxInfof(ctx, "cur_time %+v", cur_time)
		klog.CtxInfof(ctx, "cur_time %+v", cur_time)

		last_time = &cur_time
	}

	query := sql.SQL_VIDEO_UTIME + " < ?"
	if err := DB.WithContext(ctx).Order(sql.SQL_VIDEO_UTIME+" desc").Find(&res, query, time.UnixMilli(*last_time)).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetVideos] res: %+v\n", res)
		return res, err
	}
	return res, nil
}
