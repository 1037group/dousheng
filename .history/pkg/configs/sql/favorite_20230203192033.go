package sql

import "time"

type Favorite struct {
	VideoId    int64     `gorm:"column:video_id" db:"video_id" json:"video_id" form:"video_id"`
	UserId     int64     `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	IsFavorite int64     `gorm:"column:isfavorite" db:"isfavorite" json:"isfavorite" form:"isfavorite"`
	DelState   int64     `gorm:"column:delstate" db:"delstate" json:"delstate" form:"delstate"`
	Utime      time.Time `gorm:"column:utime" db:"utime" json:"utime" form:"utime"`
}
