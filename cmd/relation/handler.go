package main

import (
	"context"
	douyin_relation "github.com/1037group/dousheng/kitex_gen/douyin_relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *douyin_relation.RelationActionRequest) (resp *douyin_relation.RelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *douyin_relation.RelationFollowListRequest) (resp *douyin_relation.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *douyin_relation.RelationFriendListRequest) (resp *douyin_relation.RelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}
