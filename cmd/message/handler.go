package main

import (
	"context"

	"github.com/1037group/dousheng/cmd/message/logic"
	douyin_message "github.com/1037group/dousheng/kitex_gen/douyin_message"
	"github.com/1037group/dousheng/pack"
	"github.com/cloudwego/kitex/pkg/klog"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageChat implements the MessageServiceImpl interface. Message功能应该仅在打开聊天框的时候被调用一次
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *douyin_message.MessageChatRequest) (resp *douyin_message.MessageChatResponse, err error) {
	klog.CtxInfof(ctx, "[MessageChat] %+v", req)

	//使用事务更新消息读取情况
	res, err := logic.GetMessageList(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}
	msg := "Get messageList success"
	//转换数组到douyin_message.MessageChatResponse
	messages := pack.Messages(res)
	return &douyin_message.MessageChatResponse{
		StatusCode:  0,
		StatusMsg:   &msg,
		MessageList: messages,
	}, nil
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *douyin_message.MessageActionRequest) (resp *douyin_message.MessageActionResponse, err error) {
	klog.CtxInfof(ctx, "[MessageAction] %+v", req)

	if req.ActionType == 1 {
		//使用事务处理消息写入
		err := logic.MessageAction(ctx, req)
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

// MessageSetUnRead implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageSetUnRead(ctx context.Context, req *douyin_message.MessageSetUnReadRequest) (resp *douyin_message.MessageSetUnReadResponse, err error) {
	klog.CtxInfof(ctx, "[MessageSetUnRead] %+v", req)
	err = logic.MessageSetUnRead(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}

	msg := "Send message UnRead success"
	return &douyin_message.MessageSetUnReadResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
	}, nil
}
