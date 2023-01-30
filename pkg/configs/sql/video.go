package sql

import "time"

type Video struct {
	VideoId            int64     `gorm:"column:video_id" db:"video_id" json:"video_id" form:"video_id"`
	UserId             int64     `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	VideoPlayUrl       string    `gorm:"column:video_play_url" db:"video_play_url" json:"video_play_url" form:"video_play_url"`
	VideoCoverUrl      string    `gorm:"column:video_cover_url" db:"video_cover_url" json:"video_cover_url" form:"video_cover_url"`
	VideoFavoriteCount int64     `gorm:"column:video_favorite_count" db:"video_favorite_count" json:"video_favorite_count" form:"video_favorite_count"`
	VideoCommentCount  int64     `gorm:"column:video_comment_count" db:"video_comment_count" json:"video_comment_count" form:"video_comment_count"`
	VideoTitle         string    `gorm:"column:video_title" db:"video_title" json:"video_title" form:"video_title"`
	Ctime              time.Time `gorm:"column:ctime" db:"ctime" json:"ctime" form:"ctime"`
	Utime              time.Time `gorm:"column:utime" db:"utime" json:"utime" form:"utime"`
}

func (Video) TableName() string {
	return "video"
}

const SQL_VIDEO_VIDEO_ID = "video_id"
const SQL_VIDEO_USER_ID = "user_id"
const SQL_VIDEO_VIDEO_PLAY_URL = "video_play_url"
const SQL_VIDEO_VIDEO_COVER_URL = "video_cover_url"
const SQL_VIDEO_VIDEO_FAVORITE_COUNT = "video_favorite_count"
const SQL_VIDEO_VIDEO_COMMENT_COUNT = "video_comment_count"
const SQL_VIDEO_VIDEO_TITLE = "video_title"
const SQL_VIDEO_CTIME = "ctime"
const SQL_VIDEO_UTIME = "utime"
