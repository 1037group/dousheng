package pack

import (
	"github.com/1037group/dousheng/cmd/api/biz/model/douyin_api"
	"github.com/1037group/dousheng/kitex_gen/douyin_comment"
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

func CommentListResponseRpc2Api(m *douyin_comment.CommentListResponse) *douyin_api.CommentListResponse {
	if m == nil {
		return nil
	}

	var commentList []*douyin_api.Comment

	for _, comment := range m.CommentList {
		var user douyin_api.User
		if comment.User != nil {
			user = douyin_api.User{
				ID:            comment.User.Id,
				Name:          comment.User.Name,
				FollowCount:   comment.User.FollowCount,
				FollowerCount: comment.User.FollowerCount,
				IsFollow:      comment.User.IsFollow,
			}
		}
		one := douyin_api.Comment{
			ID:         comment.Id,
			User:       &user,
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
		}
		commentList = append(commentList, &one)
	}

	return &douyin_api.CommentListResponse{
		StatusCode:  m.StatusCode,
		StatusMsg:   m.StatusMsg,
		CommentList: commentList,
	}
}

func CommentActionResponseRpc2Api(m *douyin_comment.CommentActionResponse) *douyin_api.CommentActionResponse {
	if m == nil {
		return nil
	}

	var user douyin_api.User
	if m.Comment != nil && m.Comment.User != nil {
		user = douyin_api.User{
			ID:            m.Comment.User.Id,
			Name:          m.Comment.User.Name,
			FollowCount:   m.Comment.User.FollowCount,
			FollowerCount: m.Comment.User.FollowerCount,
			IsFollow:      m.Comment.User.IsFollow,
		}
	}
	comment := douyin_api.Comment{
		ID:         m.Comment.Id,
		User:       &user,
		Content:    m.Comment.Content,
		CreateDate: m.Comment.CreateDate,
	}

	return &douyin_api.CommentActionResponse{
		StatusCode: m.StatusCode,
		StatusMsg:  m.StatusMsg,
		Comment:    &comment,
	}
}
func FavoriteListResponseRpc2Api(m *douyin_favorite.FavoriteListResponse) *douyin_api.FavoriteListResponse {
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

	return &douyin_api.FavoriteListResponse{
		StatusCode: m.StatusCode,
		StatusMsg:  m.StatusMsg,
		VideoList:  videoList,
	}
}
