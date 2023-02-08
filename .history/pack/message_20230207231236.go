package pack

import (
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
