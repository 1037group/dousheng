// Code generated by Kitex v0.4.4. DO NOT EDIT.

package commentservice

import (
	"context"
	douyin_comment "github.com/1037group/dousheng/kitex_gen/douyin_comment"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return commentServiceServiceInfo
}

var commentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CommentService"
	handlerType := (*douyin_comment.CommentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CommentAction": kitex.NewMethodInfo(commentActionHandler, newCommentServiceCommentActionArgs, newCommentServiceCommentActionResult, false),
		"CommentList":   kitex.NewMethodInfo(commentListHandler, newCommentServiceCommentListArgs, newCommentServiceCommentListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "douyin_comment",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func commentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyin_comment.CommentServiceCommentActionArgs)
	realResult := result.(*douyin_comment.CommentServiceCommentActionResult)
	success, err := handler.(douyin_comment.CommentService).CommentAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCommentActionArgs() interface{} {
	return douyin_comment.NewCommentServiceCommentActionArgs()
}

func newCommentServiceCommentActionResult() interface{} {
	return douyin_comment.NewCommentServiceCommentActionResult()
}

func commentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyin_comment.CommentServiceCommentListArgs)
	realResult := result.(*douyin_comment.CommentServiceCommentListResult)
	success, err := handler.(douyin_comment.CommentService).CommentList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCommentListArgs() interface{} {
	return douyin_comment.NewCommentServiceCommentListArgs()
}

func newCommentServiceCommentListResult() interface{} {
	return douyin_comment.NewCommentServiceCommentListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CommentAction(ctx context.Context, req *douyin_comment.CommentActionRequest) (r *douyin_comment.CommentActionResponse, err error) {
	var _args douyin_comment.CommentServiceCommentActionArgs
	_args.Req = req
	var _result douyin_comment.CommentServiceCommentActionResult
	if err = p.c.Call(ctx, "CommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentList(ctx context.Context, req *douyin_comment.CommentListRequest) (r *douyin_comment.CommentListResponse, err error) {
	var _args douyin_comment.CommentServiceCommentListArgs
	_args.Req = req
	var _result douyin_comment.CommentServiceCommentListResult
	if err = p.c.Call(ctx, "CommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
