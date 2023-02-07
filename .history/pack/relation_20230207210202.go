package pack

import (
	"github.com/1037group/dousheng/cmd/api/biz/model/douyin_api"
	"github.com/1037group/dousheng/kitex_gen/douyin_relation"
)

func RelationFollowListResponseRpc2Api(m *douyin_relation.RelationFollowListResponse) *douyin_api.RelationFollowListResponse {
	if m == nil {
		return nil
	}

	userList := []*douyin_api.User{}

	for _, user := range m.UserList {
		one := douyin_api.User{
			ID:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		}
		userList = append(userList, &one)
	}

	return &douyin_api.RelationFollowListResponse{
		StatusCode: 0,
		StatusMsg:  m.StatusMsg,
		UserList:   userList,
	}
}

func RelationFollowerListResponseRpc2Api(m *douyin_relation.RelationFollowerListResponse) *douyin_api.RelationFollowerListResponse {
	if m == nil {
		return nil
	}

	userList := []*douyin_api.User{}

	for _, user := range m.UserList {
		one := douyin_api.User{
			ID:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		}
		userList = append(userList, &one)
	}

	return &douyin_api.RelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  m.StatusMsg,
		UserList:   userList,
	}
}

func RelationFollowerListResponseRpc2Api(m *douyin_relation.RelationFollowerListResponse) *douyin_api.RelationFollowerListResponse {
	if m == nil {
		return nil
	}

	userList := []*douyin_api.User{}

	for _, user := range m.UserList {
		one := douyin_api.User{
			ID:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		}
		userList = append(userList, &one)
	}

	return &douyin_api.RelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  m.StatusMsg,
		UserList:   userList,
	}
}
