package pack

import (
	"github.com/1037group/dousheng/cmd/api/biz/model/douyin_api"
	"github.com/1037group/dousheng/kitex_gen/douyin_user"
)

func UserResponseRpc2Api(m *douyin_user.UserResponse) *douyin_api.UserResponse {
	if m == nil || m.User == nil {
		return nil
	}

	user := douyin_api.User{
		ID:            m.User.Id,
		Name:          m.User.Name,
		FollowCount:   m.User.FollowCount,
		FollowerCount: m.User.FollowerCount,
		IsFollow:      m.User.IsFollow,
	}

	return &douyin_api.UserResponse{
		StatusCode: m.StatusCode,
		StatusMsg:  m.StatusMsg,
		User:       &user,
	}
}
