package pack

import (
	"github.com/1037group/dousheng/kitex_gen/douyin_feed"
	"github.com/1037group/dousheng/kitex_gen/douyin_message"
	"github.com/1037group/dousheng/pkg/configs/sql"
)

func Message(m *sql.Message) *douyin_message.Message {
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
		IsFavorite:    false,
		Title:         m.VideoTitle,
	}
}
