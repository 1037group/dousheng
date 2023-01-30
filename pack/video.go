package pack

import (
	"github.com/1037group/dousheng/kitex_gen/douyin_feed"
	"github.com/1037group/dousheng/kitex_gen/douyin_user"
	"github.com/1037group/dousheng/pkg/configs/sql"
)

// Video pack video info
func Video(m *sql.Video, n *sql.User) *douyin_feed.Video {
	if m == nil {
		return nil
	}

	return &douyin_feed.Video{
		Id:            m.VideoId,
		Author:        User(n),
		PlayUrl:       m.VideoPlayUrl,
		CoverUrl:      m.VideoCoverUrl,
		FavoriteCount: m.VideoFavoriteCount,
		CommentCount:  m.VideoCommentCount,
		IsFavorite:    false,
		Title:         m.VideoTitle,
	}
}

// User pack user info
func User(m *sql.User) *douyin_user.User {
	if m == nil {
		return nil
	}

	return &douyin_user.User{
		Id:            m.UserId,
		Name:          m.UserName,
		FollowCount:   &m.UserFollowCount,
		FollowerCount: &m.UserFollowerCount,
		IsFollow:      false,
	}
}

// Videos pack list of videos info
func Videos(ms []*sql.Video, ns map[int64]sql.User) []*douyin_feed.Video {
	videos := make([]*douyin_feed.Video, 0)
	for _, m := range ms {
		user := ns[m.UserId]
		if video := Video(m, &user); video != nil {
			videos = append(videos, video)
		}
	}
	return videos
}
