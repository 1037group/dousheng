package main

import (
	"context"

	"github.com/1037group/dousheng/cmd/relation/logic"
	douyin_relation "github.com/1037group/dousheng/kitex_gen/douyin_relation"
	"github.com/cloudwego/kitex/pkg/klog"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *douyin_relation.RelationActionRequest) (resp *douyin_relation.RelationActionResponse, err error) {
	klog.CtxInfof(ctx, "[RelationAction] %+v", req)

	err = logic.RelationAction(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}
	return &douyin_relation.RelationActionResponse{
		StatusCode: 0,
		StatusMsg:  nil,
	}, nil
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *douyin_relation.RelationFollowListRequest) (resp *douyin_relation.RelationFollowListResponse, err error) {
	klog.CtxInfof(ctx, "[RelationFollowList] %+v", req)

	userList, err := logic.RelationFollowList(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}
	return &douyin_relation.RelationFollowListResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		UserList:   userList,
	}, nil
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *douyin_relation.RelationFriendListRequest) (resp *douyin_relation.RelationFriendListResponse, err error) {
	// TODO: Your code here...
	klog.CtxInfof(ctx, "[RelationFriendList] %+v", req)

	userList, err := logic.RelationFollowerList(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}
	return &douyin_relation.RelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		UserList:   userList,
	}, nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *douyin_relation.RelationFollowerListRequest) (resp *douyin_relation.RelationFollowerListResponse, err error) {
	klog.CtxInfof(ctx, "[RelationFollowerList] %+v", req)

	userList, err := logic.RelationFollowerList(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}
	return &douyin_relation.RelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		UserList:   userList,
	}, nil
}
