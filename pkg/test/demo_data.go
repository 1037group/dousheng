package test

import (
	"github.com/1037group/dousheng/cmd/api/biz/model/douyin_api"
)

var DemoVideos = []*douyin_api.Video{
	{
		ID:            1,
		Author:        &DemoUser,
		PlayURL:       "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var followCount int64 = 0
var followerCount int64 = 0

var DemoUser = douyin_api.User{
	ID:            1,
	Name:          "TestUser",
	FollowCount:   &followCount,
	FollowerCount: &followerCount,
	IsFollow:      false,
}
