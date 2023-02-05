package main

import (
	"context"
	douyin_comment "github.com/1037group/dousheng/kitex_gen/douyin_comment"
	"github.com/cloudwego/kitex/pkg/klog"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *douyin_comment.CommentActionRequest) (resp *douyin_comment.CommentActionResponse, err error) {
	// TODO: Your code here...
	klog.CtxInfof(ctx, "[CommentAction] %+v", req)

	return
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *douyin_comment.CommentListRequest) (resp *douyin_comment.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}
