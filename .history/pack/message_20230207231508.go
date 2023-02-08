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
	return &douyin_message.Message{
		Id:         m.MessageId,
		Content:    m.CommentContent,
		CreateTime: m.Ctime,
	}
}

func Messages(ms []*sql.Video) []*douyin_message.Message {
	videos := make([]*douyin_feed.Video, 0)
	for _, m := range ms {
		user := ns[m.UserId]
		if video := Video(m, &user); video != nil {
			videos = append(videos, video)
		}
	}
	return videos
}
