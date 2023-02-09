package main

import (
	"context"
	"github.com/1037group/dousheng/dal/db"
	douyin_comment "github.com/1037group/dousheng/kitex_gen/douyin_comment"
	"github.com/1037group/dousheng/kitex_gen/douyin_user"
	"github.com/1037group/dousheng/pack"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/1037group/dousheng/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *douyin_comment.CommentActionRequest) (resp *douyin_comment.CommentActionResponse, err error) {

	klog.CtxInfof(ctx, "[CommentAction] %+v", req)

	switch req.ActionType {
	case 1:
		t := time.Now()
		comment := sql.Comment{
			CommentId:      0,
			VideoId:        req.VideoId,
			UserId:         req.UserId,
			CommentContent: *req.CommentText,
			Ctime:          t,
			Utime:          t,
		}
		err = db.CreateComment(ctx, db.DB, &comment)
		if err != nil {
			klog.CtxInfof(ctx, err.Error())
			return nil, errno.ParamErr
		}

		user, err := db.GetUserByID(ctx, db.DB, req.UserId)
		if err != nil {
			klog.CtxInfof(ctx, err.Error())
			return nil, err
		}

		msg := "post comment success"
		userinfo := douyin_user.User{
			Id:            user.UserId,
			Name:          user.UserName,
			FollowCount:   &user.UserFollowCount,
			FollowerCount: &user.UserFollowerCount,
			IsFollow:      true, //自己post的评论，默认true
		}
		commentinfo := pack.Comment(&comment, &userinfo)
		return &douyin_comment.CommentActionResponse{
			StatusCode: 0,
			StatusMsg:  &msg,
			Comment:    commentinfo,
		}, nil
	case 2:
		err := db.DeleteComment(ctx, db.DB, *req.CommentId)
		if err != nil {
			return nil, err
		}
		msg := "delete comment success"
		return &douyin_comment.CommentActionResponse{
			StatusCode: 0,
			StatusMsg:  &msg,
			Comment:    nil,
		}, nil
	default:
		klog.CtxInfof(ctx, "action type error\n")
		return nil, nil
	}

}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *douyin_comment.CommentListRequest) (resp *douyin_comment.CommentListResponse, err error) {

	klog.CtxInfof(ctx, "[CommentList]: %v\n", req)

	res, err := db.MGetCommentByVideoId(ctx, db.DB, &req.VideoId)
	if err != nil {
		klog.CtxInfof(ctx, err.Error())
		return nil, err
	}
	klog.CtxInfof(ctx, "res: %v\n", res)
	userMap := make(map[int64]douyin_user.User)
	var userIDs []int64
	for _, m := range res {
		userIDs = append(userIDs, m.UserId)
	}
	users, err := db.MGetUserByID(ctx, db.DB, userIDs)
	if err != nil {
		klog.CtxInfof(ctx, err.Error())
		return nil, err
	}
	for _, m := range users {
		var isFollow bool

		// 设置了token时检验是否follow
		if req.ReqUserId != nil {
			isFollow, err = db.CheckFollow(ctx, db.DB, *req.ReqUserId, m.UserId)
			if err != nil {
				klog.CtxErrorf(ctx, err.Error())
				return nil, err
			}
		}
		userMap[m.UserId] = douyin_user.User{
			Id:            m.UserId,
			Name:          m.UserName,
			FollowCount:   &m.UserFollowCount,
			FollowerCount: &m.UserFollowerCount,
			IsFollow:      isFollow,
		}
	}

	commentList := pack.Comments(res, userMap)
	return &douyin_comment.CommentListResponse{
		StatusCode:  0,
		StatusMsg:   "string",
		CommentList: commentList,
	}, nil
}
