package sql

import "time"

type Message struct {
	MessageId      int64     `gorm:"primaryKey;column:message_id" db:"message_id" json:"message_id" form:"message_id"`
	StoreByUserId  int64     `gorm:"column:store_by_user_id" db:"store_by_user_id" json:"store_by_user_id" form:"store_by_user_id"`
	UserId         int64     `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	ToUserId       int64     `gorm:"column:to_user_id" db:"to_user_id" json:"to_user_id" form:"to_user_id"`
	CommentContent string    `gorm:"column:message_content" db:"message_content" json:"message_content" form:"message_content"`
	IsRead         int32     `gorm:"column:is_read" db:"is_read" json:"is_read" form:"is_read"`
	Ctime          time.Time `gorm:"column:ctime" db:"ctime" json:"ctime" form:"ctime"`
	Utime          time.Time `gorm:"column:utime" db:"utime" json:"utime" form:"utime"`
}

func (Message) TableName() string {
	return "message"
}

type MessageSort []*Message

//PersonSort 实现sort SDK 中的Interface接口

func (s MessageSort) Len() int {
	//返回传入数据的总数
	return len(s)
}
func (s MessageSort) Swap(i, j int) {
	//两个对象满足Less()则位置对换
	//表示执行交换数组中下标为i的数据和下标为j的数据
	s[i], s[j] = s[j], s[i]
}
func (s MessageSort) Less(i, j int) bool {
	//按字段比较大小,此处是降序排序
	//返回数组中下标为i的数据是否小于下标为j的数据
	return s[i].Ctime.Before(s[j].Ctime)
}

const SQL_MESSAGE_COMMENT_ID = "message_id"
const SQL_MESSAGE_STORE_BY_USER_ID = "store_by_user_id"
const SQL_MESSAGE_USER_ID = "user_id"
const SQL_MESSAGE_TO_USER_ID = "to_user_id"
const SQL_MESSAGE_COMMENT_CONTENT = "comment_content"
const SQL_MESSAGE_IS_READ = "is_read"
const SQL_MESSAGE_CTIME = "ctime"
const SQL_MESSAGE_UTIME = "utime"
