package pack

import (
	"github.com/1037group/dousheng/kitex_gen/douyin_feed"
	"github.com/1037group/dousheng/kitex_gen/douyin_user"
	"github.com/1037group/dousheng/pkg/configs/sql"
)

// Video pack video info
func Video(m *sql.Video, n *douyin_user.User, m_isfavor int32) *douyin_feed.Video {
	if m == nil {
		return nil
	}
	var isfavor bool
	if m_isfavor == 1 {
		isfavor = true
	} else {
		isfavor = false
	}
	return &douyin_feed.Video{
		Id:            m.VideoId,
		Author:        n,
		PlayUrl:       m.VideoPlayUrl,
		CoverUrl:      m.VideoCoverUrl,
		FavoriteCount: m.VideoFavoriteCount,
		CommentCount:  m.VideoCommentCount,
		IsFavorite:    isfavor,
		Title:         m.VideoTitle,
	}
}

// Videos pack list of videos info
func Videos(ms []*sql.Video, ns map[int64]douyin_user.User, isfavor []int32) []*douyin_feed.Video {
	videos := make([]*douyin_feed.Video, 0)
	for _, m := range ms {
		user := ns[m.UserId]
		if video := Video(m, &user, isfavor); video != nil {
			videos = append(videos, video)
		}
	}
	return videos
}
