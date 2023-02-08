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

func Messages(ms []*sql.Message) []*douyin_message.Message {
	messages := make([]*douyin_message.Message, 0)
	for _, m := range ms {
		if message := Message(m); message != nil {
			messages = append(messages, message)
		}
	}
	return messages
}
