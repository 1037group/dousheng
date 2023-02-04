package sql

import "time"

type Favorite struct {
	VideoId    int64 `gorm:"column:video_id" db:"video_id" json:"video_id" form:"video_id"`
	UserId     int64 `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	IsFavorite int64
	Ctime      time.Time `gorm:"column:ctime" db:"ctime" json:"ctime" form:"ctime"`
	Utime      time.Time `gorm:"column:utime" db:"utime" json:"utime" form:"utime"`
}
