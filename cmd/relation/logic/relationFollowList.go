package logic

import (
	"context"
	"github.com/1037group/dousheng/dal/db"
	"github.com/1037group/dousheng/kitex_gen/douyin_relation"
	"github.com/1037group/dousheng/kitex_gen/douyin_user"
	"github.com/cloudwego/kitex/pkg/klog"
)

func RelationFollowList(ctx context.Context, req *douyin_relation.RelationFollowListRequest) (users []*douyin_user.User, err error) {
	klog.CtxInfof(ctx, "[logic.RelationFollowList] req: %+v", req)

	relations, err := db.MGetFollowList(ctx, db.DB, req.UserId)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}

	if len(relations) == 0 {
		return nil, nil
	}
	userIds := make([]int64, len(relations))
	for index, relation := range relations {
		userIds[index] = relation.ToUserId
	}

	sqlUsers, err := db.MGetUserByID(ctx, db.DB, userIds)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}

	users = make([]*douyin_user.User, len(sqlUsers))
	for index, _ := range users {
		isFollow, err := db.CheckFollow(ctx, db.DB, req.ReqUserId, sqlUsers[index].UserId)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return nil, err
		}
		users[index] = &douyin_user.User{
			Id:            sqlUsers[index].UserId,
			Name:          sqlUsers[index].UserName,
			FollowCount:   &sqlUsers[index].UserFollowCount,
			FollowerCount: &sqlUsers[index].UserFollowerCount,
			IsFollow:      isFollow,
		}
	}

	return users, nil
}
