package db

import (
	"context"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
)

// MGetCommentByVideoId multiple get list of comments
func MGetCommentByVideoId(ctx context.Context, vedio_id *int64) ([]*sql.Comment, error) {
	klog.CtxInfof(ctx, "[MGetCommentVideoId] last_time: %v\n", vedio_id)
	var res []*sql.Comment

	return res, nil
}
