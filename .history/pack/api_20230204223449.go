package pack

import (
	"github.com/1037group/dousheng/cmd/api/biz/model/douyin_api"
	"github.com/1037group/dousheng/kitex_gen/douyin_favorite"
	"github.com/1037group/dousheng/kitex_gen/douyin_feed"
	"github.com/1037group/dousheng/kitex_gen/douyin_publish"
)

func FeedResponseRpc2Api(m *douyin_feed.FeedResponse) *douyin_api.FeedResponse {
	if m == nil {
		return nil
	}

	var videoList []*douyin_api.Video

	for _, video := range m.VideoList {
		var user douyin_api.User
		if video.Author != nil {
			user = douyin_api.User{
				ID:            video.Author.Id,
				Name:          video.Author.Name,
				FollowCount:   video.Author.FollowCount,
				FollowerCount: video.Author.FollowerCount,
				IsFollow:      video.Author.IsFollow,
			}
		}

		one := douyin_api.Video{
			ID:            video.Id,
			Author:        &user,
			PlayURL:       video.PlayUrl,
			CoverURL:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		}
		videoList = append(videoList, &one)
	}

	return &douyin_api.FeedResponse{
		StatusCode: m.StatusCode,
		StatusMsg:  m.StatusMsg,
		VideoList:  videoList,
		NextTime:   m.NextTime,
	}
}

func PublishListResponseRpc2Api(m *douyin_publish.PublishListResponse) *douyin_api.PublishListResponse {
	if m == nil {
		return nil
	}

	var videoList []*douyin_api.Video

	for _, video := range m.VideoList {
		var user douyin_api.User
		if video.Author != nil {
			user = douyin_api.User{
				ID:            video.Author.Id,
				Name:          video.Author.Name,
				FollowCount:   video.Author.FollowCount,
				FollowerCount: video.Author.FollowerCount,
				IsFollow:      video.Author.IsFollow,
			}
		}

		one := douyin_api.Video{
			ID:            video.Id,
			Author:        &user,
			PlayURL:       video.PlayUrl,
			CoverURL:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		}
		videoList = append(videoList, &one)
	}

	return &douyin_api.PublishListResponse{
		StatusCode: m.StatusCode,
		StatusMsg:  m.StatusMsg,
		VideoList:  videoList,
	}
}

func FavoriteListResponseRpc2Api(m *douyin_favorite.FavoriteListResponse) *douyin_api.FavoriteListResponse {

}
