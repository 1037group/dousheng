package sql

import "time"

type Relation struct {
	RelationId int64     `gorm:"primaryKey;column:relation_id" db:"relation_id" json:"relation_id" form:"relation_id"`
	UserId     int64     `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	ToUserId   int64     `gorm:"column:to_user_id" db:"to_user_id" json:"to_user_id" form:"to_user_id"`
	Status     uint      `gorm:"column:status" db:"status" json:"status" form:"status"`
	Ctime      time.Time `gorm:"column:ctime" db:"ctime" json:"ctime" form:"ctime"`
	Utime      time.Time `gorm:"column:utime" db:"utime" json:"utime" form:"utime"`
}

func (Relation) TableName() string {
	return "relation"
}

const SQL_RELATION_RELATION_ID = "relation_id"
const SQL_RELATION_USER_ID = "user_id"
const SQL_RELATION_TO_USER_ID = "to_user_id"
const SQL_RELATION_STATUS = "status"
const SQL_RELATION_CTIME = "ctime"
const SQL_RELATION_UTIME = "utime"
const SQL_RELATION_STATUS_FOLLOW = 1
const SQL_RELATION_STATUS_NOTFOLLOW = 2
