package pack

import (
	"github.com/1037group/dousheng/cmd/api/biz/model/douyin_api"
	"github.com/1037group/dousheng/kitex_gen/douyin_message"
	"github.com/1037group/dousheng/pkg/configs/sql"
)

func Message(m *sql.Message) *douyin_message.Message {
	if m == nil {
		return nil
	}
	time := m.Ctime.String()
	return &douyin_message.Message{
		Id:         m.MessageId,
		Content:    m.CommentContent,
		CreateTime: &time,
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

func MessageChatResponseRpc2Api(m *douyin_message.MessageChatResponse) *douyin_api.MessageChatResponse {
	if m == nil {
		return nil
	}
	messageList := []*douyin_api.Message{}
	for _, message := range m.MessageList {
		one := douyin_api.Message{
			ID:         message.Id,
			Content:    message.Content,
			CreateTime: message.CreateTime,
		}
		messageList = append(messageList, &one)
	}
	return &douyin_api.MessageChatResponse{
		StatusCode:  m.StatusCode,
		StatusMsg:   m.StatusMsg,
		MessageList: messageList,
	}
}
