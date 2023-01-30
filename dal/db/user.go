package db

import (
	"context"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
)

// MGetUserById multiple get list of user info
func MGetUserByID(ctx context.Context, userIDs []int64) ([]*sql.User, error) {
	klog.CtxInfof(ctx, "[MGetUserByID] userIDs: %+v\n", userIDs)

	res := make([]*sql.User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	query := sql.SQL_USER_USER_ID + " in ?"
	if err := DB.WithContext(ctx).Where(query, userIDs).Find(&res).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetUserByID] res: %+v\n", res)
		return res, err
	}
	return res, nil
}

// QueryUser get user by userName
func GetUserByUserName(ctx context.Context, userName string) ([]*sql.User, error) {
	klog.CtxInfof(ctx, "[GetUserByUserName] userName: %+v\n", userName)

	res := make([]*sql.User, 0)

	query := sql.SQL_USER_USER_NAME + " = ?"
	if err := DB.WithContext(ctx).Where(query, userName).Find(&res).Error; err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, user *sql.User) error {
	klog.CtxInfof(ctx, "[CreateUser] userName: %+v\n", user)
	return DB.WithContext(ctx).Create(user).Error
}
