package sql

import "time"

type Favorite struct {
	// FavoriteId int64     `gorm:"column:favorite_id" db:"favorite_id" json:"favorite_id" form:"favorite_id"`
	VideoId    int64     `gorm:"column:video_id" db:"video_id" json:"video_id" form:"video_id"`
	UserId     int64     `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	IsFavorite bool      `gorm:"column:isfavorite" db:"isfavorite" json:"isfavorite" form:"isfavorite"`
	DelState   int32     `gorm:"column:delstate" db:"delstate" json:"delstate" form:"delstate"`
	Utime      time.Time `gorm:"column:utime" db:"utime" json:"utime" form:"utime"`
}

func (Favorite) TableName() string {
	return "favorite"
}

// const SQL_FAVORITE_ID = "favorite_id"
const SQL_FAVORITE_VIDEO_ID = "video_id"
const SQL_FAVORITE_USER_ID = "user_id"
const SQL_FAVORITE_ISFAVORITE = "isfavorite"
const SQL_FAVORITE_DELSTATE = "delstate"
const SQL_FAVORITE_UTIME = "utime"
