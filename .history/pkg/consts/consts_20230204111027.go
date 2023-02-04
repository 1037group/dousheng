package consts

const (
	SecretKey           = "secret key"
	IdentityKey         = "id"
	ApiServiceName      = "douyinapi"
	CommentServiceName  = "comment_service"
	FavoriteServiceName = "favorite_service"
	FeedServiceName     = "feed_service"
	MessageServiceName  = "message_service"
	PublishServiceName  = "publish_service"
	RelationServiceName = "relation_service"
	UserServiceName     = "user_service"

	MySQLDefaultDSN = "gorm:gorm@tcp(172.26.48.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP             = "tcp"

	RedisIp         = "172.26.48.1"
	RedisPort       = "6379"
	RedisExpireTime = 600

	CommentServiceAddr  = ":10370"
	FavoriteServiceAddr = ":10371"
	FeedServiceAddr     = ":10372"
	MessageServiceAddr  = ":10373"
	PublishServiceAddr  = ":10374"
	RelationServiceAddr = ":10375"
	UserServiceAddr     = ":10376"

	ExportEndpoint = ":4317"
	ETCDAddress    = "127.0.0.1:2379"
	DefaultLimit   = 10
)
