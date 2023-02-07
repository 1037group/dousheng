package logic

import (
	"context"

	"github.com/1037group/dousheng/kitex_gen/douyin_relation"
	"github.com/1037group/dousheng/kitex_gen/douyin_user"
)

func RelationFriendList(ctx context.Context, req *douyin_relation.RelationFollowListRequest) (users []*douyin_user.User, err error) {

}
