package sql

import "time"

type Message struct {
	MessageId      int64     `gorm:"primaryKey;column:comment_id" db:"comment_id" json:"comment_id" form:"comment_id"`
	VideoId        int64     `gorm:"column:video_id" db:"video_id" json:"video_id" form:"video_id"`
	ToUserId       int64     `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	CommentContent string    `gorm:"column:comment_content" db:"comment_content" json:"comment_content" form:"comment_content"`
	Ctime          time.Time `gorm:"column:ctime" db:"ctime" json:"ctime" form:"ctime"`
	Utime          time.Time `gorm:"column:utime" db:"utime" json:"utime" form:"utime"`
}

func (Message) TableName() string {
	return "message"
}
