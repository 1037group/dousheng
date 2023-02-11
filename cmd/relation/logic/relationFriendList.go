package logic

import (
	"context"

	"github.com/1037group/dousheng/dal/db"
	"github.com/1037group/dousheng/kitex_gen/douyin_relation"
	"github.com/1037group/dousheng/kitex_gen/douyin_user"
	"github.com/cloudwego/kitex/pkg/klog"
)

func RelationFriendList(ctx context.Context, req *douyin_relation.RelationFriendListRequest) (users []*douyin_user.FriendUser, err error) {
	klog.CtxInfof(ctx, "[logic.RelationFriendList] req: %+v", req)
	relations, err := db.MGetFollowerList(ctx, db.DB, req.UserId)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}

	if len(relations) == 0 {
		return nil, nil
	}
	userIds := make([]int64, len(relations))
	for index, relation := range relations {
		userIds[index] = relation.UserId
	}

	sqlUsers, err := db.MGetUserByID(ctx, db.DB, userIds)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}

	users = make([]*douyin_user.FriendUser, len(sqlUsers))
	for index := range users {
		isFollow, err := db.CheckFollow(ctx, db.DB, req.ReqUserId, sqlUsers[index].UserId)
		klog.CtxInfof(ctx, "[isFollow] isFollow: %+v", isFollow)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return nil, err
		}
		//获取最后一条信息
		res, err := db.GetLastMessage(ctx, db.DB, &req.ReqUserId, &sqlUsers[index].UserId)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return nil, err
		}

		var msgType int64
		var CommentContent string
		//需要区分两种情况，存在信息和不存在信息
		if len(res) > 0 {
			if res[0].UserId == req.ReqUserId {
				msgType = 1
			} else {
				msgType = 0
			}
			CommentContent = res[0].CommentContent
		} else {
			msgType = 0
			CommentContent = ""
		}

		users[index] = &douyin_user.FriendUser{
			Id:            sqlUsers[index].UserId,
			Name:          sqlUsers[index].UserName,
			FollowCount:   &sqlUsers[index].UserFollowCount,
			FollowerCount: &sqlUsers[index].UserFollowerCount,
			IsFollow:      isFollow,
			Avatar:        "",
			Message:       &CommentContent,
			MsgType:       msgType,
		}
	}

	return users, nil
}
