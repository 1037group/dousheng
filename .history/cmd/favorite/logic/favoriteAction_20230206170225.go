package logic

import (
	"context"

	"github.com/1037group/dousheng/kitex_gen/douyin_favorite"
	"github.com/cloudwego/kitex/pkg/klog"
)

func FavoriteAction(ctx context.Context, req *douyin_favorite.FavoriteActionRequest) (err error) {
	klog.CtxInfof(ctx, "[logic.FavoriteAction] req: %+v", req)

	return nil
}
