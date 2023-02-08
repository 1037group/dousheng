package main

import (
	"context"
	"time"

	"github.com/1037group/dousheng/dal/db"
	douyin_message "github.com/1037group/dousheng/kitex_gen/douyin_message"
	"github.com/1037group/dousheng/pack"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *douyin_message.MessageChatRequest) (resp *douyin_message.MessageChatResponse, err error) {
	klog.CtxInfof(ctx, "[MessageChat] %+v", req)

	res, err := db.MGetMessageList(ctx, db.DB, &req.UserId, &req.ToUserId)
	msg := "Get messageList success"
	//转换数组到douyin_message.MessageChatResponse
	messages := pack.Messages(res)
	return &douyin_message.MessageChatResponse{
		StatusCode:  0,
		StatusMsg:   &msg,
		MessageList: res,
	}, nil
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *douyin_message.MessageActionRequest) (resp *douyin_message.MessageActionResponse, err error) {
	klog.CtxInfof(ctx, "[MessageAction] %+v", req)
	t := time.Now()
	message := sql.Message{
		UserId:         req.UserId,
		ToUserId:       req.ToUserId,
		CommentContent: req.Content,
		Ctime:          t,
		Utime:          t,
	}
	if req.ActionType == 1 {
		err := db.SendMessage(ctx, &message)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return nil, err
		}
	} else {
		msg := "Incorrect action type"
		return &douyin_message.MessageActionResponse{
			StatusCode: 0,
			StatusMsg:  &msg,
		}, nil
	}
	msg := "Send message success"
	return &douyin_message.MessageActionResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
	}, nil
}
