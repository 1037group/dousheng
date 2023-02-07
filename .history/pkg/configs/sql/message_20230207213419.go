package sql

import "time"

type Message struct {
	MessageId      int64     `gorm:"primaryKey;column:message_id" db:"message_id" json:"message_id" form:"message_id"`
	UserId         int64     `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	ToUserId       int64     `gorm:"column:to_user_id" db:"to_user_id" json:"to_user_id" form:"to_user_id"`
	CommentContent string    `gorm:"column:message_content" db:"message_content" json:"message_content" form:"message_content"`
	Ctime          time.Time `gorm:"column:ctime" db:"ctime" json:"ctime" form:"ctime"`
	Utime          time.Time `gorm:"column:utime" db:"utime" json:"utime" form:"utime"`
}

func (Message) TableName() string {
	return "message"
}
