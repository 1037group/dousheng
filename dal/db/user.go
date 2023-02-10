package db

import (
	"context"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// GetUserById  get user info
func GetUserByID(ctx context.Context, tx *gorm.DB, userID int64) (*sql.User, error) {
	klog.CtxInfof(ctx, "[db.GetUserByID] userID: %+v\n", userID)

	res := &sql.User{}

	query := sql.SQL_USER_USER_ID + " = ?"
	if err := tx.WithContext(ctx).Where(query, userID).Find(&res).Error; err != nil {
		klog.CtxInfof(ctx, "[db.MGetUserByID] res: %+v\n", res)
		return res, err
	}
	return res, nil
}

// MGetUserById multiple get list of user info
func MGetUserByID(ctx context.Context, tx *gorm.DB, userIDs []int64) ([]*sql.User, error) {
	klog.CtxInfof(ctx, "[db.MGetUserByID] userIDs: %+v\n", userIDs)

	res := make([]*sql.User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	query := sql.SQL_USER_USER_ID + " in ?"
	if err := tx.WithContext(ctx).Where(query, userIDs).Find(&res).Error; err != nil {
		klog.CtxInfof(ctx, "[db.MGetUserByID] res: %+v\n", res)
		return res, err
	}
	return res, nil
}

// QueryUser get user by userName
func GetUserByUserName(ctx context.Context, tx *gorm.DB, userName string) ([]*sql.User, error) {
	klog.CtxInfof(ctx, "[db.GetUserByUserName] userName: %+v\n", userName)

	res := make([]*sql.User, 0)

	query := sql.SQL_USER_USER_NAME + " = ?"
	if err := tx.WithContext(ctx).Where(query, userName).Find(&res).Error; err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, tx *gorm.DB, user *sql.User) error {
	klog.CtxInfof(ctx, "[db.CreateUser] user: %+v\n", user)

	return tx.WithContext(ctx).Create(user).Error
}

func AddFollowCount(ctx context.Context, tx *gorm.DB, userId int64) error {
	klog.CtxInfof(ctx, "[db.AddFollowCount] userId: %+v\n", userId)

	user := &sql.User{UserId: userId}
	return tx.Model(&user).UpdateColumn(sql.SQL_USER_USER_FOLLOW_COUNT, gorm.Expr(sql.SQL_USER_USER_FOLLOW_COUNT+" + ?", 1)).Error
}

func AddFollowerCount(ctx context.Context, tx *gorm.DB, userId int64) error {
	klog.CtxInfof(ctx, "[db.AddFollowerCount] userId: %+v\n", userId)

	user := &sql.User{UserId: userId}
	return tx.Model(&user).UpdateColumn(sql.SQL_USER_USER_FOLLOWER_COUNT, gorm.Expr(sql.SQL_USER_USER_FOLLOWER_COUNT+" + ?", 1)).Error
}

func MinusFollowCount(ctx context.Context, tx *gorm.DB, userId int64) error {
	klog.CtxInfof(ctx, "[db.MinusFollowCount] userId: %+v\n", userId)

	user := &sql.User{UserId: userId}
	return tx.Model(&user).UpdateColumn(sql.SQL_USER_USER_FOLLOW_COUNT, gorm.Expr(sql.SQL_USER_USER_FOLLOW_COUNT+" - ?", 1)).Error
}

func MinusFollowerCount(ctx context.Context, tx *gorm.DB, userId int64) error {
	klog.CtxInfof(ctx, "[db.MinusFollowerCount] userId: %+v\n", userId)

	user := &sql.User{UserId: userId}
	return tx.Model(&user).UpdateColumn(sql.SQL_USER_USER_FOLLOWER_COUNT, gorm.Expr(sql.SQL_USER_USER_FOLLOWER_COUNT+" - ?", 1)).Error
}
