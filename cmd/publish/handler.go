package main

import (
	"context"
	douyin_publish "github.com/1037group/dousheng/kitex_gen/douyin_publish"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *douyin_publish.PublishActionRequest) (resp *douyin_publish.PublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *douyin_publish.PublishListRequest) (resp *douyin_publish.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}
