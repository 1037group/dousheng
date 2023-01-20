// Code generated by hertz generator. DO NOT EDIT.

package DouyinApi

import (
	douyin_api "github.com/1037group/dousheng/cmd/api/biz/handler/douyin_api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		_douyin.GET("/feed", append(_feedMw(), douyin_api.Feed)...)
		_douyin.GET("/user", append(_userMw(), douyin_api.User)...)
		{
			_comment := _douyin.Group("/comment", _commentMw()...)
			_comment.POST("/action", append(_comment_ctionMw(), douyin_api.CommentAction)...)
			_comment.GET("/list", append(_commentlistMw(), douyin_api.CommentList)...)
		}
		{
			_favorite := _douyin.Group("/favorite", _favoriteMw()...)
			_favorite.POST("/action", append(_favorite_ctionMw(), douyin_api.FavoriteAction)...)
			_favorite.GET("/list", append(_favoritelistMw(), douyin_api.FavoriteList)...)
		}
		{
			_message := _douyin.Group("/message", _messageMw()...)
			_message.POST("/action", append(_message_ctionMw(), douyin_api.MessageAction)...)
			_message.GET("/chat", append(_messagechatMw(), douyin_api.MessageChat)...)
		}
		{
			_publish := _douyin.Group("/publish", _publishMw()...)
			_publish.POST("/action", append(_publish_ctionMw(), douyin_api.PublishAction)...)
			_publish.GET("/list", append(_publishlistMw(), douyin_api.PublishList)...)
		}
		{
			_relation := _douyin.Group("/relation", _relationMw()...)
			_relation.POST("/action", append(_relation_ctionMw(), douyin_api.RelationAction)...)
			{
				_follow := _relation.Group("/follow", _followMw()...)
				_follow.GET("/list", append(_relationfollowlistMw(), douyin_api.RelationFollowList)...)
			}
			{
				_friend := _relation.Group("/friend", _friendMw()...)
				_friend.GET("/list", append(_relationfriendlistMw(), douyin_api.RelationFriendList)...)
			}
		}
		{
			_user0 := _douyin.Group("/user", _user0Mw()...)
			_user0.POST("/login", append(_userloginMw(), douyin_api.UserLogin)...)
			_user0.POST("/register", append(_userregisterMw(), douyin_api.UserRegister)...)
		}
	}
}
