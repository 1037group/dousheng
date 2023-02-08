// Code generated by hertz generator.

package douyin_api

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/1037group/dousheng/cmd/api/biz/logic"
	douyin_api "github.com/1037group/dousheng/cmd/api/biz/model/douyin_api"
	"github.com/1037group/dousheng/cmd/api/biz/mw"
	"github.com/1037group/dousheng/cmd/api/biz/rpc"
	"github.com/1037group/dousheng/kitex_gen/douyin_favorite"
	"github.com/1037group/dousheng/kitex_gen/douyin_feed"
	"github.com/1037group/dousheng/kitex_gen/douyin_message"
	"github.com/1037group/dousheng/kitex_gen/douyin_publish"
	"github.com/1037group/dousheng/kitex_gen/douyin_relation"
	"github.com/1037group/dousheng/kitex_gen/douyin_user"
	"github.com/1037group/dousheng/pack"
	"github.com/1037group/dousheng/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Feed .
// @router /douyin/feed [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "[Feed] api is called.")
	var err error
	var req douyin_api.FeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var pReqUserId *int64
	if req.Token != nil {
		// parse userId from token
		mw.JwtMiddleware.MiddlewareFunc()(ctx, c)
		type JwtUnauthorized struct {
			Code int32 `json:"code"`
		}
		var jwtUnauthorized JwtUnauthorized
		json.Unmarshal(c.Response.Body(), &jwtUnauthorized)

		if jwtUnauthorized.Code == errno.AuthorizationFailedErr.ErrCode {
			hlog.CtxErrorf(ctx, errno.AuthorizationFailedErr.ErrMsg)
			return
		}

		reqUser, _ := c.Get(mw.JwtMiddleware.IdentityKey)
		reqUserId := reqUser.(*douyin_api.User).ID
		hlog.CtxInfof(ctx, "userId: %+v", reqUserId)
		pReqUserId = &reqUserId
	}

	rpcResp, err := rpc.Feed(ctx, &douyin_feed.FeedRequest{
		LatestTime: req.LatestTime,
		ReqUserId:  pReqUserId,
	})
	hlog.CtxInfof(ctx, "[Feed] api call rpc end. rpcResp: %+v", rpcResp)
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	resp := pack.FeedResponseRpc2Api(rpcResp)

	c.JSON(consts.StatusOK, resp)
}

// UserLogin .
// @router /douyin/user/login [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "[UserLogin] api is called.")
	mw.JwtMiddleware.LoginHandler(ctx, c)
}

// UserRegister .
// @router /douyin/user/register [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "[UserRegister] api is called.")
	var err error
	var req douyin_api.UserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	rpcResp, err := rpc.UserRegister(ctx, &douyin_user.UserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	hlog.CtxInfof(ctx, "[UserRegister] api call rpc end. rpcResp: %+v", rpcResp)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	token, _, err := mw.JwtMiddleware.TokenGenerator(rpcResp.UserId)

	resp := new(douyin_api.UserRegisterResponse)
	resp.Token = token
	resp.UserID = rpcResp.UserId
	resp.StatusCode = rpcResp.StatusCode
	resp.StatusMsg = rpcResp.StatusMsg

	c.JSON(consts.StatusOK, resp)
}

// User .
// @router /douyin/user [GET]
func User(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "[User] api is called.")
	var err error
	var req douyin_api.UserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	user, _ := c.Get(mw.JwtMiddleware.IdentityKey)
	hlog.CtxInfof(ctx, "user_Id: %+v", user.(*douyin_api.User).ID)

	rpcResp, err := rpc.UserInfo(ctx, &douyin_user.UserRequest{
		UserId: req.UserID,
	})
	hlog.CtxInfof(ctx, "[User] api call rpc end. rpcResp: %+v", rpcResp)

	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := pack.UserResponseRpc2Api(rpcResp)

	c.JSON(consts.StatusOK, resp)
}

// PublishAction .
// @router /douyin/publish/action [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "[PublishAction] api is called.")
	var err error

	title := c.PostForm("title")
	if title == "" {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	file, err := c.FormFile("data")
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	// path
	path := "./tmp"

	// check
	if _, err := os.Stat(path); err != nil {
		hlog.CtxInfof(ctx, "path not exists ", path)
		err := os.MkdirAll(path, 0711)
		if err != nil {
			c.String(consts.StatusInternalServerError, err.Error())
			hlog.CtxErrorf(ctx, "Error creating directory tmp")
			return
		}
	}

	// Save file local
	videoPath := fmt.Sprintf("./tmp/%s", file.Filename)
	err = c.SaveUploadedFile(file, videoPath)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	reqUser, _ := c.Get(mw.JwtMiddleware.IdentityKey)
	reqUserId := reqUser.(*douyin_api.User).ID
	hlog.CtxInfof(ctx, "userId: %+v", reqUserId)

	// Upload mp4 file
	videoUrl, videoMd5, err := logic.UploadVideoToCOS(ctx, videoPath, reqUserId)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	hlog.CtxInfof(ctx, "[videoUrl]: %+v", videoUrl)

	// Generate cover
	snapshotPath := fmt.Sprintf("./tmp/%s.jpeg", videoMd5)
	err = logic.GetSnapshot(ctx, videoPath, snapshotPath, 1)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	// Upload jpeg file
	imgUrl, err := logic.UploadImgToCOS(ctx, snapshotPath, reqUserId)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	hlog.CtxInfof(ctx, "[imgUrl]: %+v", imgUrl)

	// rpc
	rpcResp, err := rpc.PublishAction(ctx, &douyin_publish.PublishActionRequest{
		ReqUserId:     reqUserId,
		Title:         title,
		VideoPlayUrl:  videoUrl,
		VideoCoverUrl: imgUrl,
	})
	hlog.CtxInfof(ctx, "[PublishAction] api call rpc end. rpcResp: %+v", rpcResp)
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	resp := douyin_api.PublishActionResponse{
		StatusCode: rpcResp.StatusCode,
		StatusMsg:  rpcResp.StatusMsg,
	}

	c.JSON(consts.StatusOK, resp)
}

// PublishList .
// @router /douyin/publish/list [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "[PublishList] api is called.")
	var err error
	var req douyin_api.PublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	rpcResp, err := rpc.PublishList(ctx, &douyin_publish.PublishListRequest{
		UserId: req.UserID,
	})
	hlog.CtxInfof(ctx, "[PublishList] api call rpc end. rpcResp: %+v", rpcResp)
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	resp := pack.PublishListResponseRpc2Api(rpcResp)
	c.JSON(consts.StatusOK, resp)
}

// FavoriteAction .
// @router /douyin/favorite/action [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "[FavoriteAction] api is called.")
	var err error
	var req douyin_api.FavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	// parse userId from token
	user, _ := c.Get(mw.JwtMiddleware.IdentityKey)
	hlog.CtxInfof(ctx, "user: %+v", user)
	userId := user.(*douyin_api.User).ID
	hlog.CtxInfof(ctx, "userId: %+v", userId)

	rpcResp, err := rpc.FavoriteAction(ctx, &douyin_favorite.FavoriteActionRequest{
		UserId:     userId,
		VideoId:    req.VideoID,
		ActionType: req.ActionType,
	})
	hlog.CtxInfof(ctx, "[FavoriteAction] api call rpc end. rpcResp: %+v", rpcResp)
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	resp := douyin_api.FavoriteActionResponse{
		StatusCode: rpcResp.StatusCode,
		StatusMsg:  rpcResp.StatusMsg,
	}

	c.JSON(consts.StatusOK, resp)
}

// FavoriteList .
// @router douyin/favorite/list [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "[FavoriteList] api is called.")
	var err error
	var req douyin_api.FavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	rpcResp, err := rpc.FavoriteList(ctx, &douyin_favorite.FavoriteListRequest{
		UserId: req.UserID,
	})
	hlog.CtxInfof(ctx, "[FavoriteList] api call rpc end. rpcResp: %+v", rpcResp)
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	resp := pack.FavoriteListResponseRpc2Api(rpcResp)

	c.JSON(consts.StatusOK, resp)
}

// CommentAction .
// @router douyin/comment/action [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_api.CommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin_api.CommentActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// CommentList .
// @router douyin/comment/list [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_api.CommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin_api.CommentListResponse)

	c.JSON(consts.StatusOK, resp)
}

// RelationAction .
// @router douyin/relation/action [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "[RelationAction] api is called.")

	var err error
	var req douyin_api.RelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// parse userId from token
	reqUser, _ := c.Get(mw.JwtMiddleware.IdentityKey)
	reqUserId := reqUser.(*douyin_api.User).ID
	hlog.CtxInfof(ctx, "userId: %+v", reqUserId)

	toUserId, err := strconv.ParseInt(req.ToUserID, 10, 64)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	actionType, err := strconv.Atoi(req.ActionType)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// rpc
	rpcResp, err := rpc.RelationAction(ctx, &douyin_relation.RelationActionRequest{
		ReqUserId:  reqUserId,
		ToUserId:   toUserId,
		ActionType: int32(actionType),
	})
	hlog.CtxInfof(ctx, "[RelationAction] api call rpc end. rpcResp: %+v", rpcResp)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	resp := &douyin_api.RelationActionResponse{
		StatusCode: rpcResp.StatusCode,
		StatusMsg:  rpcResp.StatusMsg,
	}

	c.JSON(consts.StatusOK, resp)
}

// RelationFollowList .
// @router douyin/relation/follow/list [GET]
func RelationFollowList(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "[RelationFollowList] api is called.")

	var err error
	var req douyin_api.RelationFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// parse userId from token
	reqUser, _ := c.Get(mw.JwtMiddleware.IdentityKey)
	reqUserId := reqUser.(*douyin_api.User).ID
	hlog.CtxInfof(ctx, "reqUserId: %+v", reqUserId)

	// rpc
	userId, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	rpcResp, err := rpc.RelationFollowList(ctx, &douyin_relation.RelationFollowListRequest{
		UserId:    userId,
		ReqUserId: reqUserId,
	})
	hlog.CtxInfof(ctx, "[RelationFollowList] api call rpc end. rpcResp: %+v", rpcResp)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	resp := pack.RelationFollowListResponseRpc2Api(rpcResp)

	c.JSON(consts.StatusOK, resp)
}

// RelationFriendList .
// @router douyin/relation/friend/list [GET]
func RelationFriendList(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "[RelationFriendList] api is called.")

	var err error
	var req douyin_api.RelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// parse userId from token
	reqUser, _ := c.Get(mw.JwtMiddleware.IdentityKey)
	reqUserId := reqUser.(*douyin_api.User).ID
	hlog.CtxInfof(ctx, "reqUserId: %+v", reqUserId)

	// rpc
	userId := req.UserID
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	rpcResp, err := rpc.RelationFriendList(ctx, &douyin_relation.RelationFriendListRequest{
		UserId:    userId,
		ReqUserId: reqUserId,
	})
	hlog.CtxInfof(ctx, "[RelationFriendList] api call rpc end. rpcResp: %+v", rpcResp)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	resp := pack.RelationFriendListResponseRpc2Api(rpcResp)

	c.JSON(consts.StatusOK, resp)
}

// MessageChat .
// @router douyin/message/chat [GET]
func MessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_api.MessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin_api.MessageChatResponse)

	c.JSON(consts.StatusOK, resp)
}

// MessageAction .
// @router douyin/message/action [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "[MessageAction] api is called.")

	var err error
	var req douyin_api.MessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	User, _ := c.Get(mw.JwtMiddleware.IdentityKey)
	UserId := User.(*douyin_api.User).ID
	hlog.CtxInfof(ctx, "userId: %+v", UserId)

	actionType := req.ActionType

	rpcResp, err := rpc.RelationAction(ctx, &douyin_message.MessageActionRequest{
		UserId:     UserId,
		ToUserId:   toUserId,
		ActionType: int32(actionType),
	})
	hlog.CtxInfof(ctx, "[RelationAction] api call rpc end. rpcResp: %+v", rpcResp)

	resp := new(douyin_api.MessageActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// RelationFollowerList .
// @router douyin/relation/follower/list/ [GET]
func RelationFollowerList(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "[RelationFollowerList] api is called.")

	var err error
	var req douyin_api.RelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// parse userId from token
	reqUser, _ := c.Get(mw.JwtMiddleware.IdentityKey)
	reqUserId := reqUser.(*douyin_api.User).ID
	hlog.CtxInfof(ctx, "reqUserId: %+v", reqUserId)

	// rpc
	userId, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	rpcResp, err := rpc.RelationFollowerList(ctx, &douyin_relation.RelationFollowerListRequest{
		UserId:    userId,
		ReqUserId: reqUserId,
	})
	hlog.CtxInfof(ctx, "[RelationFollowerList] api call rpc end. rpcResp: %+v", rpcResp)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	resp := pack.RelationFollowerListResponseRpc2Api(rpcResp)

	c.JSON(consts.StatusOK, resp)
}
