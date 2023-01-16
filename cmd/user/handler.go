package main

import (
	"context"
	douyin_user "github.com/1037group/dousheng/kitex_gen/douyin_user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *douyin_user.UserLoginRequest) (resp *douyin_user.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *douyin_user.UserRegisterRequest) (resp *douyin_user.UserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// User implements the UserServiceImpl interface.
func (s *UserServiceImpl) User(ctx context.Context, req *douyin_user.UserRequest) (resp *douyin_user.UserResponse, err error) {
	// TODO: Your code here...
	return
}
