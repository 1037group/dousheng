package main

import (
	"context"
	"time"

	douyin_message "github.com/1037group/dousheng/kitex_gen/douyin_message"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *douyin_message.MessageChatRequest) (resp *douyin_message.MessageChatResponse, err error) {
	// TODO: Your code here...
	return
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *douyin_message.MessageActionRequest) (resp *douyin_message.MessageActionResponse, err error) {
	klog.CtxInfof(ctx, "[MessageAction] %+v", req)
	t := time.Now()
	message := sql.Message{
		UserId:         req.UserId,
		ToUserId:       req.ToUserId,
		CommentContent: req.Content,
		Utime:          t,
	}
	return
}
