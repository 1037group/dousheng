package db

import (
	"context"

	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
)

//

// CreateFavorite create a favorite for a video
func CreateFavorite(ctx context.Context, favor *sql.Favorite) error {
	klog.CtxInfof(ctx, "[CreateFavorite] user: %+v\n", favor)
	return DB.WithContext(ctx).Create(favor).Error
}
