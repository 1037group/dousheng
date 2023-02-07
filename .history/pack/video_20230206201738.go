package pack

import (
	"github.com/1037group/dousheng/kitex_gen/douyin_feed"
	"github.com/1037group/dousheng/kitex_gen/douyin_user"
	"github.com/1037group/dousheng/pkg/configs/sql"
)

// Video pack video info
func Video(m *sql.Video, n *douyin_user.User, isfavorite *int64) *douyin_feed.Video {
	if m == nil {
		return nil
	}

	return &douyin_feed.Video{
		Id:            m.VideoId,
		Author:        n,
		PlayUrl:       m.VideoPlayUrl,
		CoverUrl:      m.VideoCoverUrl,
		FavoriteCount: m.VideoFavoriteCount,
		CommentCount:  m.VideoCommentCount,
		IsFavorite:    isfavorite,
		Title:         m.VideoTitle,
	}
}

// Videos pack list of videos info
func Videos(ms []*sql.Video, ns map[int64]douyin_user.User, isfavorite []int64) []*douyin_feed.Video {
	videos := make([]*douyin_feed.Video, 0)
	for i, m := range ms {
		user := ns[m.UserId]
		if video := Video(m, &user, &isfavorite[i]); video != nil {
			videos = append(videos, video)
		}
	}
	return videos
}
