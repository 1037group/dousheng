package db

import (
	"context"
	"fmt"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

func UpdateRelation(ctx context.Context, tx *gorm.DB, userId int64, toUserId int64, actionType int32) error {
	klog.CtxInfof(ctx, "[db.UpdateRelation] userId: %+v, toUserId: %+v, actionType: %+v", userId, toUserId, actionType)
	t := time.Now()
	relation := sql.Relation{
		UserId:   userId,
		ToUserId: toUserId,
		Status:   1,
		Ctime:    t,
		Utime:    t,
	}
	tableName := sql.Relation{}.TableName()
	results := tx.Table(tableName).Where(fmt.Sprintf("%s = ? and %s = ?", sql.SQL_RELATION_USER_ID, sql.SQL_RELATION_TO_USER_ID), userId, toUserId).First(&relation)
	if results.Error != nil {
		if results.Error == gorm.ErrRecordNotFound {
			results := tx.Table(tableName).Create(&relation)
			if results.Error != nil {
				klog.CtxErrorf(ctx, results.Error.Error())
				return results.Error
			}
		}
	} else {
		relation.Status = uint(actionType)
		results := tx.Table(tableName).Where(fmt.Sprintf("%s = ? and %s = ?", sql.SQL_RELATION_USER_ID, sql.SQL_RELATION_TO_USER_ID), userId, toUserId).Updates(map[string]interface{}{sql.SQL_RELATION_STATUS: relation.Status, sql.SQL_RELATION_UTIME: t})
		if results.Error != nil {
			klog.CtxErrorf(ctx, results.Error.Error())
			return results.Error
		}
	}
	return nil
}

// CheckFollow check if userId follows toUserId
func CheckFollow(ctx context.Context, tx *gorm.DB, userId int64, toUserId int64) (bool, error) {
	klog.CtxInfof(ctx, "[db.CheckFollow] userId: %+v, toUserId: %+v", userId, toUserId)
	relation := sql.Relation{}
	results := tx.Table(relation.TableName()).Where(fmt.Sprintf("%s = ? and %s = ? and %s = ?", sql.SQL_RELATION_USER_ID, sql.SQL_RELATION_TO_USER_ID, sql.SQL_RELATION_STATUS), userId, toUserId, sql.SQL_RELATION_STATUS_FOLLOW).First(&relation)
	if results.Error != nil {
		if results.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		klog.CtxErrorf(ctx, results.Error.Error())
		return false, results.Error
	}
	return true, nil
}

func MGetFollowList(ctx context.Context, tx *gorm.DB, userId int64) ([]*sql.Relation, error) {
	klog.CtxInfof(ctx, "[db.MGetFollowList] userId: %+v", userId)
	res := make([]*sql.Relation, 0)

	results := tx.WithContext(ctx).Where(fmt.Sprintf("%s = ? and %s = ?", sql.SQL_RELATION_USER_ID, sql.SQL_RELATION_STATUS), userId, sql.SQL_RELATION_STATUS_FOLLOW).Find(&res)
	if results.Error != nil {
		if results.Error == gorm.ErrRecordNotFound {
			return res, nil
		}
		klog.CtxErrorf(ctx, results.Error.Error())
		return nil, results.Error
	}
	return res, nil
}

func MGetFollowerList(ctx context.Context, tx *gorm.DB, userId int64) ([]*sql.Relation, error) {
	klog.CtxInfof(ctx, "[db.MGetFollowerList] userId: %+v", userId)
	res := make([]*sql.Relation, 0)

	results := tx.WithContext(ctx).Where(fmt.Sprintf("%s = ? and %s = ?", sql.SQL_RELATION_TO_USER_ID, sql.SQL_RELATION_STATUS), userId, sql.SQL_RELATION_STATUS_FOLLOW).Find(&res)
	if results.Error != nil {
		if results.Error == gorm.ErrRecordNotFound {
			return res, nil
		}
		klog.CtxErrorf(ctx, results.Error.Error())
		return nil, results.Error
	}
	return res, nil
}
