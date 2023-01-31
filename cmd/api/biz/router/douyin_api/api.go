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
		{
			_comment := _douyin.Group("/comment", _commentMw()...)
			{
				_action := _comment.Group("/action", _actionMw()...)
				_action.POST("/", append(_comment_ctionMw(), douyin_api.CommentAction)...)
			}
			{
				_list := _comment.Group("/list", _listMw()...)
				_list.GET("/", append(_commentlistMw(), douyin_api.CommentList)...)
			}
		}
		{
			_favorite := _douyin.Group("/favorite", _favoriteMw()...)
			{
				_action0 := _favorite.Group("/action", _action0Mw()...)
				_action0.POST("/", append(_favorite_ctionMw(), douyin_api.FavoriteAction)...)
			}
			{
				_list0 := _favorite.Group("/list", _list0Mw()...)
				_list0.GET("/", append(_favoritelistMw(), douyin_api.FavoriteList)...)
			}
		}
		{
			_feed := _douyin.Group("/feed", _feedMw()...)
			_feed.GET("/", append(_feed0Mw(), douyin_api.Feed)...)
		}
		{
			_message := _douyin.Group("/message", _messageMw()...)
			{
				_action1 := _message.Group("/action", _action1Mw()...)
				_action1.POST("/", append(_message_ctionMw(), douyin_api.MessageAction)...)
			}
			{
				_chat := _message.Group("/chat", _chatMw()...)
				_chat.GET("/", append(_messagechatMw(), douyin_api.MessageChat)...)
			}
		}
		{
			_publish := _douyin.Group("/publish", _publishMw()...)
			{
				_action2 := _publish.Group("/action", _action2Mw()...)
				_action2.POST("/", append(_publish_ctionMw(), douyin_api.PublishAction)...)
			}
			{
				_list1 := _publish.Group("/list", _list1Mw()...)
				_list1.GET("/", append(_publishlistMw(), douyin_api.PublishList)...)
			}
		}
		{
			_relation := _douyin.Group("/relation", _relationMw()...)
			{
				_action3 := _relation.Group("/action", _action3Mw()...)
				_action3.POST("/", append(_relation_ctionMw(), douyin_api.RelationAction)...)
			}
			{
				_follow := _relation.Group("/follow", _followMw()...)
				{
					_list2 := _follow.Group("/list", _list2Mw()...)
					_list2.GET("/", append(_relationfollowlistMw(), douyin_api.RelationFollowList)...)
				}
			}
			{
				_friend := _relation.Group("/friend", _friendMw()...)
				{
					_list3 := _friend.Group("/list", _list3Mw()...)
					_list3.GET("/", append(_relationfriendlistMw(), douyin_api.RelationFriendList)...)
				}
			}
		}
		{
			_user := _douyin.Group("/user", _userMw()...)
			_user.GET("/", append(_user0Mw(), douyin_api.User)...)
			{
				_login := _user.Group("/login", _loginMw()...)
				_login.POST("/", append(_userloginMw(), douyin_api.UserLogin)...)
			}
			{
				_register := _user.Group("/register", _registerMw()...)
				_register.POST("/", append(_userregisterMw(), douyin_api.UserRegister)...)
			}
		}
	}
}
