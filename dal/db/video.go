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

func AddFavoriteCount(ctx context.Context, tx *gorm.DB, video_id int64) error {
	klog.CtxInfof(ctx, "[db.AddFavoriteCount] video_id : %+v\n", video_id)

	video := &sql.Video{VideoId: video_id}
	return tx.Model(&video).UpdateColumn(sql.SQL_VIDEO_VIDEO_FAVORITE_COUNT, gorm.Expr(sql.SQL_VIDEO_VIDEO_FAVORITE_COUNT+" + ?", 1)).Error
}

func MinusFavoriteCount(ctx context.Context, tx *gorm.DB, video_id int64) error {
	klog.CtxInfof(ctx, "[db.AddFavoriteCount] video_id : %+v\n", video_id)

	video := &sql.Video{VideoId: video_id}
	return tx.Model(&video).UpdateColumn(sql.SQL_VIDEO_VIDEO_FAVORITE_COUNT, gorm.Expr(sql.SQL_VIDEO_VIDEO_FAVORITE_COUNT+" - ?", 1)).Error
}

func GetFavoriteCount(ctx context.Context, tx *gorm.DB, video_id int64) (int64, error) {
	klog.CtxInfof(ctx, "[db.UpdateFavoriteCount] video_id : %+v\n", video_id)

	video := &sql.Video{VideoId: video_id}
	res := tx.First(&video)
	return video.VideoFavoriteCount, res.Error
}

func UpdateVideo(ctx context.Context, tx *gorm.DB, video_id int64, param map[string]interface{}) error {
	klog.CtxInfof(ctx, "[db.UpdateVideo] video_id: %+v, param: %+v\n", video_id, param)

	video := &sql.Video{VideoId: video_id}
	return tx.Model(&video).Updates(param).Error
}

func AddCommentCount(ctx context.Context, tx *gorm.DB, video_id int64) error {
	klog.CtxInfof(ctx, "[db.AddCommentCount] video_id : %+v\n", video_id)

	video := &sql.Video{VideoId: video_id}
	return tx.Model(&video).UpdateColumn(sql.SQL_VIDEO_VIDEO_COMMENT_COUNT, gorm.Expr(sql.SQL_VIDEO_VIDEO_COMMENT_COUNT+" + ?", 1)).Error
}

func MinusCommentCount(ctx context.Context, tx *gorm.DB, video_id int64) error {
	klog.CtxInfof(ctx, "[db.MinusCommentCount] video_id : %+v\n", video_id)

	video := &sql.Video{VideoId: video_id}
	return tx.Model(&video).UpdateColumn(sql.SQL_VIDEO_VIDEO_COMMENT_COUNT, gorm.Expr(sql.SQL_VIDEO_VIDEO_COMMENT_COUNT+" - ?", 1)).Error
}

func UpdateCommentCount(ctx context.Context, tx *gorm.DB, video_id int64, param map[string]interface{}) error {
	klog.CtxInfof(ctx, "[db.UpdateCommentCount] video_id : %+v\n", video_id)

	video := &sql.Video{VideoId: video_id}
	return tx.Model(&video).Updates(param).Error
}

func GetCommentCount(ctx context.Context, tx *gorm.DB, video_id int64) (int64, error) {
	klog.CtxInfof(ctx, "[db.GetCommentCount] video_id : %+v\n", video_id)

	video := &sql.Video{VideoId: video_id}
	res := tx.First(&video)
	return video.VideoCommentCount, res.Error
}
