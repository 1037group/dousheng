package dal

import (
	"github.com/1037group/dousheng/dal/db"
	"github.com/1037group/dousheng/dal/redis"
)

// Init init dal
func Init() {
	db.Init() // mysql init
	redis.Init()
}
