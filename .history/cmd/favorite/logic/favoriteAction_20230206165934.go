package logic

import (
	"context"

	"github.com/1037group/dousheng/kitex_gen/douyin_relation"
	"github.com/cloudwego/kitex/pkg/klog"
)

func FavoriteAction(ctx context.Context, req *douyin_relation.RelationActionRequest) (err error) {
	klog.CtxInfof(ctx, "[logic.FavoriteAction] req: %+v", req)
}
