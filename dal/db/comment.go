package db

import (
	"context"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// MGetCommentByVideoId multiple get list of comments, order by the time of comment
func MGetCommentByVideoId(ctx context.Context, tx *gorm.DB, vedio_id *int64) ([]*sql.Comment, error) {
	klog.CtxInfof(ctx, "[MGetCommentVideoId] video_id: %v\n", vedio_id)
	var res []*sql.Comment
	query := sql.SQL_VIDEO_VIDEO_ID + " = ?"
	if err := tx.WithContext(ctx).Order(sql.SQL_COMMENT_UTIME+" desc").Find(&res, query, vedio_id).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetCommentVideoId] res: %v\n", res)
		return res, err
	}
	return res, nil
}

// CreateComment create comment
func CreateComment(ctx context.Context, tx *gorm.DB, comment *sql.Comment) error {
	klog.CtxInfof(ctx, "[CreateComment] comment: %v\n", comment)
	return tx.WithContext(ctx).Create(comment).Error
}

// DeleteComment delete comment
func DeleteComment(ctx context.Context, tx *gorm.DB, comment_id int64) error {
	klog.CtxInfof(ctx, "[DeleteComment] comment_id: %v\n", comment_id)
	return tx.WithContext(ctx).Where("comment_id = ? ", comment_id).Delete(&sql.Comment{}).Error
}
