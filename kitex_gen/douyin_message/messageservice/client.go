// Code generated by Kitex v0.4.4. DO NOT EDIT.

package messageservice

import (
	"context"
	douyin_message "github.com/1037group/dousheng/kitex_gen/douyin_message"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MessageChat(ctx context.Context, req *douyin_message.MessageChatRequest, callOptions ...callopt.Option) (r *douyin_message.MessageChatResponse, err error)
	MessageAction(ctx context.Context, req *douyin_message.MessageActionRequest, callOptions ...callopt.Option) (r *douyin_message.MessageActionResponse, err error)
	MessageSetUnRead(ctx context.Context, req *douyin_message.MessageSetUnReadRequest, callOptions ...callopt.Option) (r *douyin_message.MessageSetUnReadResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kMessageServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kMessageServiceClient struct {
	*kClient
}

func (p *kMessageServiceClient) MessageChat(ctx context.Context, req *douyin_message.MessageChatRequest, callOptions ...callopt.Option) (r *douyin_message.MessageChatResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageChat(ctx, req)
}

func (p *kMessageServiceClient) MessageAction(ctx context.Context, req *douyin_message.MessageActionRequest, callOptions ...callopt.Option) (r *douyin_message.MessageActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageAction(ctx, req)
}

func (p *kMessageServiceClient) MessageSetUnRead(ctx context.Context, req *douyin_message.MessageSetUnReadRequest, callOptions ...callopt.Option) (r *douyin_message.MessageSetUnReadResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageSetUnRead(ctx, req)
}
