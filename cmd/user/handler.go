package main

import (
	"context"
	"github.com/1037group/dousheng/cmd/user/logic"
	"github.com/1037group/dousheng/dal/db"
	douyin_user "github.com/1037group/dousheng/kitex_gen/douyin_user"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/1037group/dousheng/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *douyin_user.UserLoginRequest) (*douyin_user.UserLoginResponse, error) {
	klog.CtxInfof(ctx, "[UserLoginRequest] %+v", req)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		return nil, errno.AuthorizationFailedErr
	}

	users, err := db.GetUserByUserName(ctx, req.Username)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, errno.AuthorizationFailedErr
	}
	if len(users) == 0 {
		klog.CtxErrorf(ctx, errno.UserNotExistErr.ErrMsg)
		return nil, errno.UserNotExistErr
	}

	user := users[0]

	klog.CtxInfof(ctx, "pwdHash:%+v", user.PasswordHash)

	if logic.CheckPasswordHash(user.PasswordHash, req.Password) == false {
		klog.CtxInfof(ctx, "req.Password: %+v", req.Password)
		return nil, errno.AuthorizationFailedErr
	}

	resp := douyin_user.UserLoginResponse{UserId: user.UserId}
	return &resp, nil
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *douyin_user.UserRegisterRequest) (*douyin_user.UserRegisterResponse, error) {
	klog.CtxInfof(ctx, "[UserRegister] %+v", req)

	users, err := db.GetUserByUserName(ctx, req.Username)

	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, errno.ParamErr
	}

	if len(users) != 0 {
		klog.CtxErrorf(ctx, errno.UserAlreadyExistErr.ErrMsg)
		return nil, errno.UserAlreadyExistErr
	}

	// TODO redis lock username

	pwdHash, err := logic.HashPassword(req.Password)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, errno.ParamErr
	}

	t := time.Now()
	user := sql.User{
		UserName:          req.Username,
		UserFollowCount:   0,
		UserFollowerCount: 0,
		Ctime:             t,
		Utime:             t,
		PasswordHash:      pwdHash,
	}
	err = db.CreateUser(ctx, &user)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, errno.ParamErr
	}

	users, err = db.GetUserByUserName(ctx, req.Username)

	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, errno.ParamErr
	}

	if len(users) == 0 {
		klog.CtxErrorf(ctx, errno.ParamErr.ErrMsg)
		return nil, errno.ParamErr
	}

	klog.CtxInfof(ctx, "userId: %+v", users[0].UserId)
	resp := douyin_user.UserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		UserId:     users[0].UserId,
	}
	return &resp, nil
}

// User implements the UserServiceImpl interface.
func (s *UserServiceImpl) User(ctx context.Context, req *douyin_user.UserRequest) (resp *douyin_user.UserResponse, err error) {
	klog.CtxInfof(ctx, "[User] %+v", req)

	userIDs := []int64{req.UserId}
	users, err := db.MGetUserByID(ctx, userIDs)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}
	if len(users) == 0 {
		klog.CtxErrorf(ctx, errno.UserNotExistErr.ErrMsg)
		return nil, errno.UserNotExistErr
	}

	// TODO IsFollow
	user := douyin_user.User{
		Id:            users[0].UserId,
		Name:          users[0].UserName,
		FollowCount:   &users[0].UserFollowCount,
		FollowerCount: &users[0].UserFollowerCount,
		IsFollow:      false,
	}
	resp = &douyin_user.UserResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		User:       &user,
	}

	return
}
