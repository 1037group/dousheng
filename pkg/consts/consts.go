package consts

const (
	SecretKey       = "secret key"
	IdentityKey     = "id"
	ApiServiceName  = "douyinapi"
	FeedServiceName = "feed_service"
	UserServiceName = "user_service"
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP             = "tcp"
	UserServiceAddr = ":9000"
	FeedServiceAddr = ":10000"
	ExportEndpoint  = ":4317"
	ETCDAddress     = "127.0.0.1:2379"
	DefaultLimit    = 10
)
