// Code generated by hertz generator.

package douyin_api

import (
	"context"
	"fmt"
	"github.com/1037group/dousheng/cmd/api/biz/logic"
	douyin_api "github.com/1037group/dousheng/cmd/api/biz/model/douyin_api"
	"github.com/1037group/dousheng/cmd/api/biz/mw"
	"github.com/1037group/dousheng/cmd/api/biz/rpc"
	"github.com/1037group/dousheng/kitex_gen/douyin_feed"
	"github.com/1037group/dousheng/kitex_gen/douyin_publish"
	"github.com/1037group/dousheng/kitex_gen/douyin_user"
	"github.com/1037group/dousheng/pack"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"os"
)

// Feed .
// @router /douyin/feed [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "[Feed] api is called.")
	var err error
	var req douyin_api.FeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	rpcResp, err := rpc.Feed(ctx, &douyin_feed.FeedRequest{
		LatestTime: req.LatestTime,
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

	user, _ := c.Get(mw.JwtMiddleware.IdentityKey)
	userId := user.(*douyin_api.User).ID
	hlog.CtxInfof(ctx, "userId: %+v", userId)

	// Upload mp4 file
	videoUrl, videoMd5, err := logic.UploadVideoToCOS(ctx, videoPath, userId)
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
	imgUrl, err := logic.UploadImgToCOS(ctx, snapshotPath, userId)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	hlog.CtxInfof(ctx, "[imgUrl]: %+v", imgUrl)

	// rpc
	rpcResp, err := rpc.PublishAction(ctx, &douyin_publish.PublishActionRequest{
		UserId:        userId,
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
	var err error
	var req douyin_api.FavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin_api.FavoriteActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// FavoriteList .
// @router douyin/favorite/list [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_api.FavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin_api.FavoriteListResponse)

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
	var err error
	var req douyin_api.RelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin_api.RelationActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// RelationFollowList .
// @router douyin/relation/follow/list [GET]
func RelationFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_api.RelationFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin_api.RelationFollowListResponse)

	c.JSON(consts.StatusOK, resp)
}

// RelationFriendList .
// @router douyin/relation/friend/list [GET]
func RelationFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_api.RelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin_api.RelationFriendListResponse)

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
	var err error
	var req douyin_api.MessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin_api.MessageActionResponse)

	c.JSON(consts.StatusOK, resp)
}
