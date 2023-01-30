package sql

import "time"

type User struct {
	UserId            int64     `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	UserName          string    `gorm:"column:user_name" db:"user_name" json:"user_name" form:"user_name"`
	UserFollowCount   int64     `gorm:"column:user_follow_count" db:"user_follow_count" json:"user_follow_count" form:"user_follow_count"`
	UserFollowerCount int64     `gorm:"column:user_follower_count" db:"user_follower_count" json:"user_follower_count" form:"user_follower_count"`
	Ctime             time.Time `gorm:"column:ctime" db:"ctime" json:"ctime" form:"ctime"`
	Utime             time.Time `gorm:"column:utime" db:"utime" json:"utime" form:"utime"`
	PasswordHash      string    `gorm:"column:password_hash" db:"password_hash" json:"password_hash" form:"password_hash"`
}

func (User) TableName() string {
	return "user"
}

const SQL_USER_USER_ID = "user_id"
const SQL_USER_USER_NAME = "user_name"
const SQL_USER_USER_FOLLOW_COUNT = "user_follow_count"
const SQL_USER_USER_FOLLOWER_COUNT = "user_follower_count"
const SQL_USER_CTIME = "ctime"
const SQL_USER_UTIME = "utime"
const SQL_USER_PASSWORDHASH = "password_hash"
